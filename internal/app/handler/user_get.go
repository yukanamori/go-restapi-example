package handler

import (
	"myapp/pkg/erreurs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Get はユーザを取得します。
func (h *UserHandler) Get(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		zap.L().Info("failed to parse id", zap.Error(err), zap.String("id", idStr))
		return erreurs.ErrBadRequestInvalidParameter
	}

	user, err := h.useucase.GetByID(uint(id))
	if err != nil {
		if err == erreurs.ErrUserNotFound {
			return echo.ErrNotFound
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, user)
}
