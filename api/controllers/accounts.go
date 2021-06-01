package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	repo "github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// AccountsController encapsula la lógica de negocio de autenticación
type AccountsController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type accountsController struct {
	accountsRepo repo.AccountsRepository
	auth         auth.Auth
}

// NewAccountsController crea un nuevo controlador de autenticación
func NewAccountsController(a repo.AccountsRepository, auth auth.Auth) AccountsController {
	return accountsController{accountsRepo: a, auth: auth}
}

func (ct accountsController) SignUp(c echo.Context) error {
	input := new(responses.SignUpResponse)
	if err := c.Bind(input); err != nil {
		return responses.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	user, err := ct.accountsRepo.SignUp(input)
	if err != nil {
		return err
	}

	token, err := auth.GenerateToken(user.Username)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, auth.JWTResponse{
		Token: token,
		User:  user,
	})
}

func (ct accountsController) Login(c echo.Context) error {
	input := new(responses.LoginResponse)
	if err := c.Bind(input); err != nil {
		return responses.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	user, err := ct.accountsRepo.Login(input)
	if err != nil {
		return err
	}

	if err := ct.auth.VerifyPassword(user.Password, input.Password); err != nil {
		return responses.Unauthorized("la contraseña ingresada es incorrecta")
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, auth.JWTResponse{
		Token:        token,
		RefreshToken: "",
		User:         responses.UserResponse(user),
	})
}
