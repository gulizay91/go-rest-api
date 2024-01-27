package utils

import (
	"log"

	"github.com/gulizay91/go-rest-api/pkg/models"
	"github.com/go-playground/validator/v10"
)

func Validate(model interface{}) *[]models.ValidationErrors {
	err := validator.New().Struct(model)
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
