package utils

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/go-playground/validator/v10"
	"github.com/gulizay91/go-rest-api/pkg/models"
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
	log.Warn("Validation failed:")
	for _, e := range err.(validator.ValidationErrors) {
		log.Warn("Field: %s, Error: %s\n", e.Field(), e.Tag())
		validationErrors = append(validationErrors, models.ValidationErrors{Field: e.StructNamespace(), Error: e.Tag()})
	}

	return &validationErrors
}

type EnumValid interface {
	Valid() bool
}

func RegisterEnumValidator(v *validator.Validate) error {
	err := v.RegisterValidation("enum", ValidateEnum)
	if err != nil {
		return err
	}
	return nil
}

func ValidateEnum(fl validator.FieldLevel) bool {
	if enum, ok := fl.Field().Interface().(EnumValid); ok {
		return enum.Valid()
	}
	return false
}
