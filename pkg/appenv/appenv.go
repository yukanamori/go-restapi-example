package appenv

// AppEnv はアプリケーションの環境を表します。
type AppEnv string

const (
	Development AppEnv = "development"
	Production  AppEnv = "production"
	Test        AppEnv = "test"
)

// IsProduction は環境が本番環境かどうかを判定します。
func (e AppEnv) IsProduction() bool {
	return e == Production
}
