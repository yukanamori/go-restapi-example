package router

import (
	"myapp/internal/app/di"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Router はルーターです。
type Router struct {
	echo *echo.Echo
}

// New はルーターを作成します。
func New(e *echo.Echo) *Router {
	return &Router{e}
}

// Register はルーティングを登録します。
func (r *Router) Register() {
	// ヘルスチェック
	r.echo.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// ユーザ
	userHandler := di.InitializeUserHandler()
	r.echo.POST("/users", userHandler.Create)
	r.echo.GET("/users/:id", userHandler.Get)
	r.echo.GET("/users/search", userHandler.Search)
	r.echo.GET("/users", userHandler.GetAll)
	r.echo.PUT("/users/:id", userHandler.Update)
	r.echo.PUT("/users/change-password", userHandler.ChangePassword)
	r.echo.DELETE("/users/:id", userHandler.Delete)
}
