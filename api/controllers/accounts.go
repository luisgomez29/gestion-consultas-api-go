package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	api "github.com/luisgomez29/gestion-consultas-api/api/errors"
	repo "github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// AccountsController encapsula la lógica de negocio de autenticación
type AccountsController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type accountsController struct {
	auth         auth.Auth
	accountsRepo repo.AccountRepository
}

// NewAccountsController crea un nuevo controlador de autenticación
func NewAccountsController(at auth.Auth, a repo.AccountRepository) AccountsController {
	return accountsController{auth: at, accountsRepo: a}
}

func (ct accountsController) SignUp(c echo.Context) error {
	input := new(responses.SignUpResponse)
	if err := c.Bind(input); err != nil {
		return api.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	if input.Password != input.PasswordConfirmation {
		return api.PasswordMismatch
	}

	// Generating password Hash
	hash, err := ct.auth.HashPassword(input.Password)
	if err != nil {
		return err
	}

	input.Password = hash
	user, err := ct.accountsRepo.CreateUser(input)
	if err != nil {
		return err
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, auth.JWTResponse{
		Token: token,
		User:  responses.UserResponse(user),
	})
}

func (ct accountsController) Login(c echo.Context) error {
	input := new(responses.LoginResponse)
	if err := c.Bind(input); err != nil {
		return api.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	user, err := ct.accountsRepo.FindUser(input)
	if err != nil {
		return err
	}

	// Check if password is valid
	match, err := ct.auth.VerifyPassword(input.Password, user.Password)
	if !match || err != nil {
		return api.Unauthorized("la contraseña ingresada es incorrecta")
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, auth.JWTResponse{
		Token:        token,
		RefreshToken: "",
		User:         responses.UserResponse(user),
	})
}
