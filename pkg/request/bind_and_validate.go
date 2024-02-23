package request

import "github.com/labstack/echo/v4"

// BindAndValidate はリクエストをバインドしてバリデーションします。
func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}

	return c.Validate(i)
}
