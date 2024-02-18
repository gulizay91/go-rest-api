package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gulizay91/go-rest-api/pkg/models"
	"github.com/gulizay91/go-rest-api/pkg/utils"
	"net/http"
	"path/filepath"
)

type IAwsService interface {
	UploadToS3Path(uploadModel *models.UploadS3FileModel) (*models.ServiceResponseModel, error)
	ListObjectsFromS3Path(deleteModel *models.ListS3FileModel) (*models.ServiceResponseModel, error)
	DeleteObjectsFromS3Path(deleteModel *models.DeleteS3FileModel) (*models.ServiceResponseModel, error)
}

type AwsService struct {
	S3Client *s3.Client
}

func NewAwsService() AwsService {
	// aws config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panicf("error: %v", err)
	}
	S3Client := s3.NewFromConfig(cfg)
	return AwsService{S3Client}
}

func (s AwsService) UploadToS3Path(uploadModel *models.UploadS3FileModel) (*models.ServiceResponseModel, error) {
	var res models.ServiceResponseModel = *models.NewErrorServiceResponseModel(nil)
	var objectIds []types.ObjectIdentifier
	if uploadModel.BeforeDeleteAllObjects == true {
		// list all files in path
		result, err := s.getListObjects(uploadModel.BucketName, uploadModel.FilePath)
		if err != nil {
			log.Errorf("Couldn't list objects from s3 bucket %v prefix %v, error: %v", uploadModel.BucketName, uploadModel.FilePath, err)
			return &res, err
		}

		for _, object := range *result {
			objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(*object.Key)})
		}
	}

	uploader := manager.NewUploader(s.S3Client, func(u *manager.Uploader) {
		// Define a strategy that will buffer 20 MiB in memory
		u.BufferProvider = manager.NewBufferedReadSeekerWriteToPool(20 * 1024 * 1024)
	})

	var uploadResultUrls models.UploadedS3FileModel
	// Loop through files:
	for _, newFile := range uploadModel.Files {
		uploadFile, err := newFile.Open()
		//log.Println(newFile.Filename, newFile.Size, newFile.Header["Content-Type"][0])
		hash, err := utils.HashRandomString(newFile.Filename)
		extension := filepath.Ext(newFile.Filename)
		newFile.Filename = *hash + extension
		key := aws.String(uploadModel.FilePath + "/" + newFile.Filename)
		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(uploadModel.BucketName),
			Key:    key,
			Body:   uploadFile,
		})

		// Check for errors
		if err != nil {
			log.Errorf("Failed upload to s3 for %v file, error: %v", newFile.Filename, err)
		}
		log.Infof("Success upload to s3 for %v - %v, aws-url: %v", newFile.Filename, newFile.Header["Content-Type"][0], result.Location)
		uploadResultUrls.FileS3Urls = append(uploadResultUrls.FileS3Urls, result.Location)
		if uploadModel.CdnUrl != nil {
			uploadResultUrls.FileCdnUrls = append(uploadResultUrls.FileCdnUrls, *uploadModel.CdnUrl+*key)
		}
	}

	// Check for errors
	if len(uploadResultUrls.FileS3Urls) == 0 {
		res.Message = "Some files can not uploaded to s3"
		return &res, errors.New(res.Message)
	}
	// Check should delete old objects
	if len(objectIds) > 0 {
		_, err := s.deleteObjects(uploadModel.BucketName, objectIds)
		if err != nil {
			log.Errorf("Couldn't delete objects from s3 bucket %v prefix %v objectIds %v, error: %v", uploadModel.BucketName, uploadModel.FilePath, objectIds, err)
		}
	}

	res = *models.NewSuccessServiceResponseModel(uploadResultUrls)
	res.StatusCode = http.StatusOK

	return &res, nil
}

func (s AwsService) ListObjectsFromS3Path(deleteModel *models.ListS3FileModel) (*models.ServiceResponseModel, error) {
	var res models.ServiceResponseModel = *models.NewErrorServiceResponseModel(nil)

	// list all files in path
	result, err := s.getListObjects(deleteModel.BucketName, deleteModel.FilePath)
	if err != nil {
		log.Errorf("Couldn't list objects from s3 bucket %v prefix %v, error: %v", deleteModel.BucketName, deleteModel.FilePath, err)
		return &res, err
	}

	res = *models.NewSuccessServiceResponseModel(result)
	res.StatusCode = http.StatusOK

	return &res, nil
}

func (s AwsService) DeleteObjectsFromS3Path(deleteModel *models.DeleteS3FileModel) (*models.ServiceResponseModel, error) {
	var res models.ServiceResponseModel = *models.NewErrorServiceResponseModel(nil)

	// list all files in path
	result, err := s.getListObjects(deleteModel.BucketName, deleteModel.FilePath)
	if err != nil {
		log.Errorf("Couldn't list objects from s3 bucket %v prefix %v, error: %v", deleteModel.BucketName, deleteModel.FilePath, err)
		return &res, err
	}

	// delete all files in path
	var objectIds []types.ObjectIdentifier
	for _, object := range *result {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(*object.Key)})
	}
	output, err := s.deleteObjects(deleteModel.BucketName, objectIds)
	if err != nil {
		log.Errorf("Couldn't delete objects from s3 bucket %v prefix %v , error: %v", deleteModel.BucketName, deleteModel.FilePath, err)
		return &res, err
	}

	res = *models.NewSuccessServiceResponseModel(nil)
	res.StatusCode = http.StatusOK
	res.Message = fmt.Sprintf("Deleted %v objects.\n", len(output.Deleted))

	return &res, nil
}

func (s AwsService) getListObjects(bucketName, filePath string) (*[]types.Object, error) {
	// list all files in path
	result, err := s.S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(filePath),
	})
	if err != nil {
		return nil, err
	}
	return &result.Contents, nil
}

func (s AwsService) deleteObjects(bucketName string, objectIds []types.ObjectIdentifier) (*s3.DeleteObjectsOutput, error) {
	// list all files in path
	output, err := s.S3Client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{Objects: objectIds},
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}
