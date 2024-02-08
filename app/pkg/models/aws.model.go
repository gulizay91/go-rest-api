package models

import "mime/multipart"

type UploadS3FileModel struct {
	Files      []*multipart.FileHeader
	FilePath   string
	BucketName string
}

type UploadedS3FileModel struct {
	FileUrls []string `json:"fileUrls"`
}
