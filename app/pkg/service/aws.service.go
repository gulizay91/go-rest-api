package service

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gulizay91/go-rest-api/pkg/models"
	"net/http"
)

type IAwsService interface {
	UploadToS3(uploadModel *models.UploadS3FileModel) (*models.ServiceResponseModel, error)
	Get() (*models.ServiceResponseModel, error)
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

func (s AwsService) UploadToS3(uploadModel *models.UploadS3FileModel) (*models.ServiceResponseModel, error) {
	var res models.ServiceResponseModel = *models.NewErrorServiceResponseModel(nil)

	uploader := manager.NewUploader(s.S3Client, func(u *manager.Uploader) {
		// Define a strategy that will buffer 20 MiB in memory
		u.BufferProvider = manager.NewBufferedReadSeekerWriteToPool(20 * 1024 * 1024)
	})

	var uploadResultUrls models.UploadedS3FileModel
	// Loop through files:
	for _, newFile := range uploadModel.Files {
		uploadFile, err := newFile.Open()
		//log.Println(newFile.Filename, newFile.Size, newFile.Header["Content-Type"][0])
		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(uploadModel.BucketName),
			Key:    aws.String(uploadModel.FilePath + "/" + newFile.Filename),
			Body:   uploadFile,
			ACL:    "public-read",
		})

		// Check for errors
		if err != nil {
			log.Errorf("Failed upload to s3 for %v file, error: %v", newFile.Filename, err)
		}
		log.Infof("Success upload to s3 for %v - %v, aws-url: %v", newFile.Filename, newFile.Header["Content-Type"][0], result.Location)
		uploadResultUrls.FileUrls = append(uploadResultUrls.FileUrls, result.Location)
	}

	// Check for errors
	if len(uploadResultUrls.FileUrls) == 0 {
		res.Message = "Some files can not uploaded to s3"
		return &res, errors.New(res.Message)
	}

	res = *models.NewSuccessServiceResponseModel(uploadResultUrls)
	res.StatusCode = http.StatusOK

	return &res, nil
}
