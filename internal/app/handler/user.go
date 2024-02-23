package handler

import "myapp/internal/domain/usecase"

// UserHandler はユーザハンドラーです。
type UserHandler struct {
	useucase usecase.UserUsecase
}

// NewUserHandler はユーザハンドラーを生成します。
func NewUserHandler(useucase usecase.UserUsecase) *UserHandler {
	return &UserHandler{useucase}
}
