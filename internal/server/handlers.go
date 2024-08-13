package server

import (
	"context"
	"net/http"
	"time"

	accountsPb "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	accountsModels "github.com/emaforlin/accounts-service/x/models"

	"github.com/labstack/echo/v4"
)

func (gw *APIGatewayServer) SignupHandler(c echo.Context) error {
	reqBody := new(accountsModels.AddPersonAccountData)
	if err := c.Bind(reqBody); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	client := accountsPb.NewAccountsClient(gw.AccountSvcConn)
	_, err := client.AddPersonAccount(ctx, reqBody.ToRPCStruct())
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest

	}
	return c.JSON(http.StatusCreated, "account successfully created")
}

func (gw *APIGatewayServer) SignupPartnerHandler(c echo.Context) error {
	reqBody := new(accountsModels.AddFoodPlaceAccountData)
	if err := c.Bind(reqBody); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	client := accountsPb.NewAccountsClient(gw.AccountSvcConn)
	_, err := client.AddFoodPlaceAccount(ctx, reqBody.ToRPCStruct())
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, "account successfully created")
}
