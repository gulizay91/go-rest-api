package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	Server        ServerConfig     `mapstructure:"server"`
	Service       ServiceConfig    `mapstructure:"service"`
	MongoDB       MongoConfig      `mapstructure:"mongoDB"`
	AwsService    AwsServiceConfig `mapstructure:"awsService"`
	K8sCluster    string           `mapstructure:"k8sCluster"`
	TestKey       string           `mapstructure:"testKey"`
	TestAnchorKey string           `mapstructure:"testAnchorKey"`
	VaultOptions  VaultOptions     `mapstructure:"vaultOptions"`
}

func (config *Config) Validate() error {
	err := validation.ValidateStruct(
		config,
		validation.Field(&config.Server),
		validation.Field(&config.Service),
		validation.Field(&config.MongoDB),
	)
	return err
}

type ServerConfig struct {
	Port string
	Addr string
}

func (config ServerConfig) Validate() error {
	err := validation.ValidateStruct(
		&config,
		validation.Field(&config.Addr, validation.Required),
		validation.Field(&config.Port, is.Port),
	)
	return err
}

type ServiceConfig struct {
	LogLevel    string `mapstructure:"logLevel"`
	Name        string
	Environment string
}

func (config ServiceConfig) Validate() error {
	err := validation.ValidateStruct(
		&config,
		validation.Field(&config.LogLevel, validation.Required),
		validation.Field(&config.Environment, validation.Required),
	)
	return err
}

type MongoConfig struct {
	Uri      string
	Database string
}

func (config MongoConfig) Validate() error {
	err := validation.ValidateStruct(
		&config,
		validation.Field(&config.Uri, validation.Required),
		validation.Field(&config.Database, validation.Required),
	)
	return err
}

type AwsServiceConfig struct {
	Region    string          `mapstructure:"region"`
	AccessKey string          `mapstructure:"accessKey"`
	SecretKey string          `mapstructure:"secretKey"`
	S3Service S3ServiceConfig `mapstructure:"s3Service"`
}

func (config AwsServiceConfig) Validate() error {
	err := validation.ValidateStruct(
		&config,
		validation.Field(&config.Region, validation.Required),
		validation.Field(&config.AccessKey, validation.Required),
		validation.Field(&config.SecretKey, validation.Required),
		validation.Field(&config.S3Service, validation.Required),
	)
	return err
}

type S3ServiceConfig struct {
	Bucket string
	CdnUrl *string
}

func (config S3ServiceConfig) Validate() error {
	err := validation.ValidateStruct(
		&config,
		validation.Field(&config.Bucket, validation.Required),
	)
	return err
}

type VaultOptions struct {
	MountPoint    string
	TestAnchorKey string
}
