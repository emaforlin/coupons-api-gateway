package main

import (
	"context"
	"net/http"
	"time"

	accountsPb "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	accountsModels "github.com/emaforlin/accounts-service/x/models"

	"github.com/labstack/echo/v4"
)

func (gw *apigatewayServer) signupHandler(c echo.Context) error {
	reqBody := new(accountsModels.AddPersonAccountData)
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, ErrBindingBody)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := accountsPb.NewAccountsClient(gw.accountSvcConn)
	res, err := client.AddPersonAccount(ctx, reqBody.ToRPCStruct())
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrCreatingAccount)
	}
	c.Set("user_id", res.Userid)
	return c.JSON(http.StatusCreated, "account successfully created")
}
