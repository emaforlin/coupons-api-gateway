package internal

import (
	"context"
	"net/http"

	accountsPb "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	accountsModels "github.com/emaforlin/accounts-service/x/models"
	"github.com/labstack/echo/v4"
)

type accountHttpHandler struct {
	router AccountRouter
}

// SignupPerson implements AccountHandler.
func (a *accountHttpHandler) SignupPerson(c echo.Context) error {
	reqBody := new(accountsModels.AddPersonAccountData)

	if err := c.Bind(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind body")
	}

	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	err := a.router.SignupPerson(ctx, &accountsPb.AddPersonAccountRequest{
		Username:    reqBody.Username,
		Email:       reqBody.Email,
		PhoneNumber: reqBody.PhoneNumber,
		Password:    reqBody.Password,
		FirstName:   reqBody.FirstName,
		LastName:    reqBody.LastName,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Account created successfully")
}

func NewAccountHandler(r AccountRouter) AccountHandler {
	return &accountHttpHandler{
		router: r,
	}

}
