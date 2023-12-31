package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	Service ServiceConfig `mapstructure:"service"`
	MongoDB MongoConfig   `mapstructure:"mongoDB"`
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
	LogLevel    string
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
