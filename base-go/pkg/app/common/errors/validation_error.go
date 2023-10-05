package errors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorData struct {
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

var ErrValidation = errors.New("validation_error")

func GetValidationErrors(errs validator.ValidationErrors) map[string]ValidationErrorData {
	data := make(map[string]ValidationErrorData)
	for _, e := range errs {
		data[e.Field()] = ValidationErrorData{
			Tag:   e.Tag(),
			Param: e.Param(),
		}
	}

	return data
}
