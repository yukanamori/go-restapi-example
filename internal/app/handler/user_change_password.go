package handler

import (
	"myapp/pkg/erreurs"
	"myapp/pkg/request"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserChangePasswordRequest はユーザのパスワード変更リクエストです。
type UserChangePasswordRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

// ChangePassword はユーザのパスワードを変更します。
func (h *UserHandler) ChangePassword(c echo.Context) error {
	res := &UserChangePasswordRequest{}
	if err := request.BindAndValidate(c, res); err != nil {
		zap.L().Info("failed to bind and validate request", zap.Error(err))
		return erreurs.ErrBadRequestInvalidParameter
	}

	if err := h.useucase.UpdatePassword(res.Username, res.Password); err != nil {
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}
