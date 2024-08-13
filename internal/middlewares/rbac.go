package middlewares

import (
	"fmt"
	"slices"

	"github.com/emaforlin/api-gateway/internal/entities"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AllowedRoles(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			fmt.Println(user.Raw)
			claims := user.Claims.(*entities.CustomClaims)
			if slices.Index(roles, claims.Role) == -1 {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
