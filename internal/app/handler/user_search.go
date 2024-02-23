package handler

import (
	"myapp/internal/domain/entity"
	"myapp/pkg/erreurs"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserSearchRequest はユーザ検索リクエストです。
type UserSearchRequest struct {
	Username  string `query:"username"`
	Email     string `query:"email"`
	FirstName string `query:"first_name"`
	LastName  string `query:"last_name"`
}

// Search はユーザを検索します。
func (h *UserHandler) Search(c echo.Context) error {
	req := &UserSearchRequest{}
	if err := c.Bind(req); err != nil {
		zap.L().Info("failed to bind request", zap.Error(err))
		return erreurs.ErrBadRequestInvalidParameter
	}

	condition := &entity.User{
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	users, err := h.useucase.Search(condition)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, users)
}
