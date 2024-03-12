package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// RFC3339Validation はRFC3339形式の日時かどうかを判定します。
func RFC3339Validation(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}
