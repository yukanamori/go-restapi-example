package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

const RFC3339Milli = "2006-01-02T15:04:05.999Z07:00"

// RFC3339Validation はRFC3339形式の日時かどうかを判定します。
func RFC3339Validation(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}

// RFC3339MilliValidation はRFC3339形式の日時かどうかを判定します。
func RFC3339MilliValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse(RFC3339Milli, fl.Field().String())
	return err == nil
}
