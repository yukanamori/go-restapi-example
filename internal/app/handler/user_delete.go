package handler

import (
	"myapp/pkg/erreurs"
	"myapp/pkg/request"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserDeleteRequest はユーザ削除リクエストです。
type UserDeleteRequest struct {
	Version uint `query:"version" validate:"required"`
}

// Delete はユーザを削除します。
func (h *UserHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		zap.L().Info("failed to parse id", zap.Error(err), zap.String("id", idStr))
		return erreurs.ErrBadRequestInvalidParameter
	}

	req := &UserDeleteRequest{}
	if err := request.BindAndValidate(c, req); err != nil {
		zap.L().Info("failed to bind and validate request", zap.Error(err))
		return erreurs.ErrBadRequestInvalidParameter
	}

	if err := h.useucase.Delete(uint(id), req.Version); err != nil {
		if err == erreurs.ErrUserNotFound {
			return echo.ErrNotFound
		} else if err == erreurs.ErrVersionMismatch {
			return echo.ErrConflict
		}

		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}
