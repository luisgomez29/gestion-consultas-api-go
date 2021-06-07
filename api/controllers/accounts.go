package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	apierrors "github.com/luisgomez29/gestion-consultas-api/api/errors"
	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// AccountsController encapsula la lógica de negocio de autenticación
type AccountsController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type accountsController struct {
	auth         auth.Auth
	accountsRepo repositories.AccountRepository
}

// NewAccountsController crea un nuevo controlador de autenticación
func NewAccountsController(at auth.Auth, a repositories.AccountRepository) AccountsController {
	return accountsController{auth: at, accountsRepo: a}
}

func (ct accountsController) SignUp(c echo.Context) error {
	input := new(responses.SignUpResponse)
	if err := c.Bind(input); err != nil {
		return apierrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	if input.Password != input.PasswordConfirmation {
		return apierrors.PasswordMismatch
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

	return ct.accountResponse(c, user)
}

func (ct accountsController) Login(c echo.Context) error {
	input := new(responses.LoginResponse)
	if err := c.Bind(input); err != nil {
		return apierrors.BadRequest("")
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
		return apierrors.Unauthorized("la contraseña ingresada es incorrecta")
	}

	return ct.accountResponse(c, user)
}

// accountResponse retorna los tokens y el usuario
func (ct accountsController) accountResponse(c echo.Context, user *models.User) error {
	tokens, err := ct.auth.TokenObtainPair(user.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, auth.JWTResponse{
		AccessToken:  tokens["access"],
		RefreshToken: tokens["refresh"],
		User:         responses.UserResponse(user),
	})
}
