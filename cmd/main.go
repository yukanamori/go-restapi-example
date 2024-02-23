package main

import (
	"fmt"
	"myapp/config"
	"myapp/internal/infrastructure/router"
	"myapp/pkg/logger"
	"myapp/pkg/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	// 設定ファイルの読み込み
	c := config.GetConfig()

	// ロガーの設定
	l := logger.NewDevelopment()
	if c.AppEnv.IsProduction() {
		l = logger.NewProduction(c.Log.Debug)
	}
	defer l.Sync()

	zap.ReplaceGlobals(l)

	// Echoの設定
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(requestLoggerConfig(l)))

	// バリデーターの設定
	e.Validator = validator.New()

	// ルーティングの設定
	r := router.New(e)
	r.Register()

	// サーバーの起動
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", c.Server.Port)))
}

func requestLoggerConfig(logger *zap.Logger) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogRequestID: true,
		LogRemoteIP:  true,
		LogHost:      true,
		LogMethod:    true,
		LogURI:       true,
		LogUserAgent: true,
		LogStatus:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("id", v.RequestID),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("host", v.Host),
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.String("user_agent", v.UserAgent),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}
}
