package models

import "mime/multipart"

type UploadS3FileModel struct {
	Files                  []*multipart.FileHeader
	FilePath               string
	BucketName             string
	BeforeDeleteAllObjects bool
	CdnUrl                 *string
}

type UploadedS3FileModel struct {
	FileS3Urls  []string `json:"fileS3Urls"`
	FileCdnUrls []string `json:"fileCdnUrls"`
}

type DeleteS3FileModel struct {
	FilePath   string
	BucketName string
}

type ListS3FileModel struct {
	FilePath   string
	BucketName string
}
