package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAll はユーザを全件取得します。
func (h *UserHandler) GetAll(c echo.Context) error {
	users, err := h.useucase.GetAll()
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, users)
}
