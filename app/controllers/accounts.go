package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/app/auth"
	apiErrors "github.com/luisgomez29/gestion-consultas-api/app/resources/api/errors"
	"github.com/luisgomez29/gestion-consultas-api/app/resources/api/requests"
	"github.com/luisgomez29/gestion-consultas-api/app/services"
)

// AccountsController represents endpoints for authentication.
type AccountsController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	VerifyToken(c echo.Context) error
	PasswordReset(c echo.Context) error
	PasswordResetConfirm(c echo.Context) error
}

type accountsController struct {
	auth            auth.Auth
	accountsService services.AccountsService
}

// NewAccountsController create a new accounts controller.
func NewAccountsController(at auth.Auth, as services.AccountsService) AccountsController {
	return accountsController{auth: at, accountsService: as}
}

func (ct accountsController) SignUp(c echo.Context) error {
	input := new(requests.SignUpRequest)
	if err := c.Bind(input); err != nil {
		return apiErrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	if input.Password != input.PasswordConfirm {
		return apiErrors.PasswordMismatch
	}

	res, err := ct.accountsService.SignUp(input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (ct accountsController) Login(c echo.Context) error {
	input := new(requests.LoginRequest)
	if err := c.Bind(input); err != nil {
		return apiErrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	res, err := ct.accountsService.Login(input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (ct accountsController) VerifyToken(c echo.Context) error {
	input := new(requests.TokenRequest)
	if err := c.Bind(input); err != nil {
		return apiErrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	res, err := ct.accountsService.VerifyToken(input.Token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

// PasswordReset verify if the user exists an email is sent with the link to reset the password,
// which has a time of 15 minutes to expire. If the user does not have an email address, nothing is sent.
func (ct accountsController) PasswordReset(c echo.Context) error {
	input := new(requests.PasswordResetRequest)
	if err := c.Bind(input); err != nil {
		return apiErrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	res, err := ct.accountsService.PasswordReset(input.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

// PasswordResetConfirm allows the user to reset the password given a token
func (ct accountsController) PasswordResetConfirm(c echo.Context) error {
	input := new(requests.PasswordResetConfirmRequest)
	if err := c.Bind(input); err != nil {
		return apiErrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	claims, err := auth.VerifyToken(input.Token, auth.JWTPasswordResetToken)
	if err != nil {
		return err
	}

	if input.Password != input.PasswordConfirm {
		return apiErrors.PasswordMismatch
	}

	res, err := ct.accountsService.PasswordResetConfirm(claims["username"].(string), input.PasswordConfirm)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
