package di

import (
	"myapp/internal/app/handler"
	"myapp/internal/domain/repository"
	"myapp/internal/domain/usecase"
	"myapp/internal/infrastructure/db"
)

// InitializeUserHandler は依存を解決したユーザハンドラーを作成します。
func InitializeUserHandler() *handler.UserHandler {
	d := db.GetDB()
	r := repository.NewUserRepository(d)
	u := usecase.NewUserUsecase(r)
	h := handler.NewUserHandler(u)
	return h
}
