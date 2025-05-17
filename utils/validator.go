package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationError(err error, obj any) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			jsonField := getJSONFieldName(obj, fieldErr.StructField())
			msg := buildErrorMessage(fieldErr)
			msg = strings.Replace(msg, fieldErr.Field(), jsonField, 1)
			errors[jsonField] = msg
		}
	} else {
		errors["error"] = err.Error()
	}

	return errors
}

func buildErrorMessage(fe validator.FieldError) string {
	fieldName := fe.Field()
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldName)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldName)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", fieldName, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", fieldName, fe.Param())
	default:
		return fmt.Sprintf("%s is invalid", fieldName)
	}
}

func getJSONFieldName(structType any, fieldName string) string {
	t := reflect.TypeOf(structType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if f, ok := t.FieldByName(fieldName); ok {
		jsonTag := f.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			return jsonTag
		}
	}
	return fieldName
}
