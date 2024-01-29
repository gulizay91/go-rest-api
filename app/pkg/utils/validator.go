package utils

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gulizay91/go-rest-api/pkg/models"
)

func Validate(model interface{}) *[]models.ValidationErrors {
	v := validator.New()
	RegisterEnumValidator(v)
	err := v.Struct(model)
	if err == nil {
		log.Println("Validation succeeded")
		return nil
	}

	var validationErrors []models.ValidationErrors
	log.Println("Validation failed:")
	for _, e := range err.(validator.ValidationErrors) {
		log.Printf("Field: %s, Error: %s\n", e.Field(), e.Tag())
		validationErrors = append(validationErrors, models.ValidationErrors{Field: e.StructNamespace(), Error: e.Tag()})
	}

	return &validationErrors
}

type EnumValid interface {
	Valid() bool
}

func RegisterEnumValidator(v *validator.Validate) {
	v.RegisterValidation("enum", ValidateEnum)
}

func ValidateEnum(fl validator.FieldLevel) bool {
	if enum, ok := fl.Field().Interface().(EnumValid); ok {
		return enum.Valid()
	}
	return false
}
