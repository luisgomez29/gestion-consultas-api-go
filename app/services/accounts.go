package services

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/app/auth"
	"github.com/luisgomez29/gestion-consultas-api/app/repositories"
)

type AccountsService interface {
	SignUp(c echo.Context) (auth.JWTResponse, error)
	Login(c echo.Context) (auth.JWTResponse, error)
	VerifyToken(c echo.Context) (echo.Map, error)
	PasswordReset(c echo.Context) (echo.Map, error)
	PasswordResetConfirm(c echo.Context) (echo.Map, error)
}

type accountsService struct {
	auth         auth.Auth
	accountsRepo repositories.AccountRepository
}
