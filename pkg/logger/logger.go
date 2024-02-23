package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// NewProduction は新しい本番用ロガーを返します。デバッグモードの場合はデバッグレベルに設定します。
func NewProduction(debug bool) *zap.Logger {
	config := zap.NewProductionConfig()

	if debug {
		// デバッグモードの場合はデバッグレベルに設定
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}

	return logger
}

// NewDevelopment は新しい開発用ロガーを返します。
func NewDevelopment() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}

	return logger
}
