package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			fieldName := fieldErr.Field()
			tag := fieldErr.Tag()

			switch tag {
			case "required":
				errors[fieldName] = fieldName + " is required."
			case "email":
				errors[fieldName] = fieldName + " must be a valid email address."
			case "min":
				errors[fieldName] = fieldName + " must be at least " + fieldErr.Param() + " characters long."
			case "max":
				errors[fieldName] = fieldName + " must be at most " + fieldErr.Param() + " characters long."
			default:
				errors[fieldName] = fieldName + " is invalid."
			}
		}
	} else {
		errors["error"] = err.Error()
	}

	return errors
}
