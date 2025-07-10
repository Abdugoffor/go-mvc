package validator

import (
	validatorV10 "github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validatorV10.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
