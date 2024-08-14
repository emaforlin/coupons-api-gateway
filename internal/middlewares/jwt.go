package middlewares

import (
	"github.com/emaforlin/api-gateway/internal/config"
	"github.com/labstack/echo/v4"
)

func JwtSkipperFunc(c echo.Context) bool {
	path := c.Request().URL.Path
	if path == config.BaseURL+"/login" {
		return true
	}
	if path == config.BaseURL+config.AccountsBaseUrl+"/signup" {
		return true
	}
	if path == config.BaseURL+"/signup/partner" {
		return true
	}
	return false
}
