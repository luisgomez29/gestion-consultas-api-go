package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// AuthController encapsula la lógica de negocio de autenticación
type AuthController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type authController struct {
	repo repositories.AuthRepository
}

// NewAuthController crea un nuevo controlador de autenticación
func NewAuthController(repo repositories.AuthRepository) AuthController {
	return authController{repo: repo}
}

func (ctrl authController) SignUp(c echo.Context) error {
	input := new(responses.SignUpResponse)
	if err := c.Bind(input); err != nil {
		return responses.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	user, err := ctrl.repo.SignUp(input)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func (ctrl authController) Login(c echo.Context) error {
	input := new(responses.LoginResponse)
	if err := c.Bind(input); err != nil {
		return responses.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	user, err := ctrl.repo.Login(input)
	if err != nil {
		return err
	}

	if err := auth.VerifyPassword(user.Password, input.Password); err != nil {
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
