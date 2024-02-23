package handler

import (
	"myapp/pkg/erreurs"
	"myapp/pkg/request"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserCreateRequest はユーザ作成リクエストです。
type UserCreateRequest struct {
	Username     string `json:"username" validate:"required,alphanum,min=4,max=32"`
	Password     string `json:"password" validate:"required,password"`
	Email        string `json:"email" validate:"required,email,max=255"`
	FirstName    string `json:"firstName" validate:"required,max=50"`
	LastName     string `json:"lastName" validate:"required,max=50"`
	ProfileImage string `json:"profileImage" validate:"base64"`
}

// Create はユーザを作成します。
func (h *UserHandler) Create(c echo.Context) error {
	req := &UserCreateRequest{}
	if err := request.BindAndValidate(c, req); err != nil {
		zap.L().Info("failed to bind and validate request", zap.Error(err))
		return erreurs.ErrBadRequestInvalidParameter
	}

	if err := h.useucase.Create(req.Username, req.Password, req.Email, req.FirstName, req.LastName, req.ProfileImage); err != nil {
		if err == erreurs.ErrUsernameAlreadyExists {
			return erreurs.ErrBadRequestUsernameAlreadyExists
		} else if err == erreurs.ErrEmailAlreadyExists {
			return erreurs.ErrBadRequestEmailAlreadyExists
		}

		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusCreated)
}
