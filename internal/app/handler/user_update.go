package handler

import (
	"myapp/pkg/erreurs"
	"myapp/pkg/request"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserUpdateRequest はユーザ更新リクエストです。
type UserUpdateRequest struct {
	Username     string `json:"username" validate:"alphanum,min=4,max=32"`
	Email        string `json:"email" validate:"email,max=255"`
	FirstName    string `json:"first_name" validate:"max=50"`
	LastName     string `json:"last_name" validate:"max=50"`
	ProfileImage string `json:"profile_image" validate:"base64"`
	Version      int    `json:"version" validate:"required"`
}

// Update はユーザを更新します。
func (h *UserHandler) Update(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		zap.L().Info("failed to parse id", zap.Error(err), zap.String("id", idStr))
		return erreurs.ErrBadRequestInvalidParameter
	}

	req := &UserUpdateRequest{}
	if err := request.BindAndValidate(c, req); err != nil {
		zap.L().Info("failed to bind and validate request", zap.Error(err))
		return erreurs.ErrBadRequestInvalidParameter
	}

	if err := h.useucase.Update(uint(id), uint(req.Version), req.Username, req.Email, req.FirstName, req.LastName, req.ProfileImage); err != nil {
		if err == erreurs.ErrUserNotFound {
			return echo.ErrNotFound
		} else if err == erreurs.ErrVersionMismatch {
			return echo.ErrConflict
		} else if err == erreurs.ErrUsernameAlreadyExists {
			return erreurs.ErrBadRequestUsernameAlreadyExists
		} else if err == erreurs.ErrEmailAlreadyExists {
			return erreurs.ErrBadRequestEmailAlreadyExists
		}

		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}
