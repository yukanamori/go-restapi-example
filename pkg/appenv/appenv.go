package appenv

// AppEnv はアプリケーションの環境を表します。
type AppEnv string

const (
	Production  AppEnv = "production"
	Test        AppEnv = "test"
	Development AppEnv = "development"
)

// IsProduction は環境が本番環境かどうかを判定します。
func (e AppEnv) IsProduction() bool {
	return e == Production
}

// IsTest は環境がテスト環境かどうかを判定します。
func (e AppEnv) IsTest() bool {
	return e == Test
}

// IsDevelopment は環境が開発環境かどうかを判定します。
func (e AppEnv) IsDevelopment() bool {
	return e == Development
}
