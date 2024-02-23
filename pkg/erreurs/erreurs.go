package erreurs

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Application Error
var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrVersionMismatch       = errors.New("version mismatch")
)

// BadRequestError は400系のエラーを表します。
type BadRequestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error はエラーメッセージを返します。
func (e *BadRequestError) Error() string {
	return e.Message
}

// NewBadRequestError はBadRequestErrorを生成します。
func NewBadRequestError(code int, message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, &BadRequestError{code, message})
}

// BadRequest Error
var (
	ErrBadRequestInvalidParameter      = NewBadRequestError(1001, "invalid parameter")
	ErrBadRequestUsernameAlreadyExists = NewBadRequestError(1002, "username already exists")
	ErrBadRequestEmailAlreadyExists    = NewBadRequestError(1003, "email already exists")
)
