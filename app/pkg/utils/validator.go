package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gulizay91/go-rest-api/pkg/models"
	"reflect"
)

func Validate(model interface{}) *[]models.ValidationErrors {
	v := validator.New()
	err := RegisterEnumValidator(v)
	if err != nil {
		log.Panicf("Enum Validation Register error: %v", err.Error())
		return nil
	}
	err = v.Struct(model)
	if err == nil {
		log.Debug("Validation succeeded")
		return nil
	}

	var validationErrors []models.ValidationErrors
	log.Warn("Validation failed!")
	for _, e := range err.(validator.ValidationErrors) {
		log.Warnf("Field: %s, Error: %s", e.Field(), e.Tag())
		validationErrors = append(validationErrors, models.ValidationErrors{Field: e.StructNamespace(), Error: e.Tag()})
	}

	return &validationErrors
}

type EnumValid interface {
	IsValid() bool
}

func RegisterEnumValidator(v *validator.Validate) error {
	err := v.RegisterValidation("enum", ValidateEnum, true)
	if err != nil {
		return err
	}
	return nil
}

func ValidateEnum(fl validator.FieldLevel) bool {
	log.Debugf("validation for %s", fl.FieldName())
	log.Debugf("field kind %v", fl.Field().Kind())
	if fl.Field().Kind() == reflect.Ptr {
		log.Debug("field is nil, default valid")
		return true
	}
	if enum, ok := fl.Field().Interface().(EnumValid); ok {
		return enum.IsValid()
	}
	return false
}
