package utils

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
	Error       bool
}

type XValidator struct {
	validate *validator.Validate
}

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	err := v.validate.Struct(data)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			elem := ErrorResponse{
				FailedField: e.Field(),
				Tag:         e.Tag(),
				Value:       e.Value(),
				Error:       true,
			}
			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
