package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// AuthController encapsula la lógica de negocio de autenticación
type AuthController interface {
	SignUp(c echo.Context) error
}

type authController struct {
	repo repositories.AuthRepository
}

// NewAuthController crea un nuevo controlador de autenticación
func NewAuthController(repo repositories.AuthRepository) AuthController {
	return authController{repo: repo}
}

func (ctrl authController) SignUp(c echo.Context) error {
	signup := new(responses.SignUpResponse)
	if err := c.Bind(signup); err != nil {
		return err
	}

	if err := signup.Validate(); err != nil {
		return responses.ValidationErrorResponse(c, err)
	}

	res, err := ctrl.repo.SignUp(signup)
	if err != nil {
		fmt.Printf("ERRORR : %v", err)
		return responses.ValidationErrorResponse(c, err)
	}
	return c.JSON(http.StatusCreated, res)
}
