package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator はバリデーターです。
type Validator struct {
	validator *validator.Validate
}

// New はバリデーターを作成します。
func New() *Validator {
	validator := validator.New()
	validator.RegisterValidation("rfc3339", RFC3339Validation)
	validator.RegisterValidation("rfc3339milli", RFC3339MilliValidation)
	validator.RegisterValidation("password", PasswordValidation)
	return &Validator{validator}
}

// Validate はバリデーションを行います。
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
