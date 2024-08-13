package server

import (
	"context"
	"net/http"
	"os"
	"time"

	accountsPb "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	accountsModels "github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/api-gateway/internal/entities"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

func (gw *APIGatewayServer) LoginHandler(c echo.Context) error {
	client := accountsPb.NewAccountsClient(gw.AccountSvcConn)
	reqBody := new(accountsModels.CheckLoginData)
	if err := c.Bind(reqBody); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := client.CheckLoginData(ctx, reqBody.ToRPCStruct())
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	if !res.GetOk() {
		return echo.ErrUnauthorized
	}
	claims := &entities.CustomClaims{
		Role: reqBody.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return echo.ErrUnauthorized
	}
	c.Set("user", token)
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

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
