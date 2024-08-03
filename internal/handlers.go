package internal

import "github.com/labstack/echo/v4"

type AccountHandler interface {
	SignupPerson(c echo.Context) error
}
