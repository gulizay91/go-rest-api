package utils

import (
	"reflect"
)

func ConvertStruct(source interface{}, dest interface{}) {
	sourceValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(dest).Elem()

	for i := 0; i < sourceValue.NumField(); i++ {
		destField := destValue.FieldByName(sourceValue.Type().Field(i).Name)
		if destField.IsValid() {
			destField.Set(sourceValue.Field(i))
		}
	}
}
