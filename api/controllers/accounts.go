package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	"github.com/luisgomez29/gestion-consultas-api/api/config"
	apierrors "github.com/luisgomez29/gestion-consultas-api/api/errors"
	"github.com/luisgomez29/gestion-consultas-api/api/mailers"
	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
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
	auth         auth.Auth
	accountsRepo repositories.AccountRepository
}

// NewAccountsController create a new accounts controller.
func NewAccountsController(at auth.Auth, a repositories.AccountRepository) AccountsController {
	return accountsController{auth: at, accountsRepo: a}
}

func (ct accountsController) SignUp(c echo.Context) error {
	input := new(responses.SignUpRequest)
	if err := c.Bind(input); err != nil {
		return apierrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	if input.Password != input.PasswordConfirm {
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
	input := new(responses.LoginRequest)
	if err := c.Bind(input); err != nil {
		return apierrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	user, err := ct.accountsRepo.FindUser(input.Username)
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

func (ct accountsController) VerifyToken(c echo.Context) error {
	input := new(responses.TokenRequest)
	if err := c.Bind(input); err != nil {
		return apierrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	claims, err := auth.VerifyToken(input.Token)
	if err != nil {
		return err
	}

	res := map[string]interface{}{
		"success": true,
		"payload": claims,
	}
	return c.JSON(http.StatusOK, res)
}

// PasswordReset verify if the user exists an email is sent with the link to reset the password,
// which has a time of 15 minutes to expire. If the user does not have an email address, nothing is sent.
func (ct accountsController) PasswordReset(c echo.Context) error {
	input := new(responses.PasswordResetRequest)
	if err := c.Bind(input); err != nil {
		return apierrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	user, err := ct.accountsRepo.FindUser(input.Username)
	if err != nil {
		var apiErr *apierrors.ErrNoRows
		if errors.As(err, &apiErr) {
			return apierrors.NewErrNoRows("el usuario no está asignado a ninguna cuenta")
		}
		return err
	}

	// Response
	res := map[string]interface{}{
		"send_email": false,
		"username":   user.Username,
		"email":      user.Email,
	}

	// Verify if the user does not have email
	if user.Email == nil {
		return c.JSON(http.StatusOK, res)
	}

	// Generate token
	claims := auth.NewClaims(user)
	claims.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()
	claims.TokenType = auth.JWTPasswordResetToken

	token, err := auth.GenerateToken(claims)
	if err != nil {
		return err
	}

	// Email message
	em := &mailers.EmailMessage{
		To:      mail.Address{Name: user.FirstName, Address: *user.Email},
		Subject: "Solicitud de recuperación de contraseña",
		Template: mailers.Template{
			Name: "accounts/password_reset_key_message.html",
			Context: map[string]interface{}{
				"currentSite":   "Gestión consultas",
				"userFirstName": user.FirstName,
				"passwordResetUrl": fmt.Sprintf(
					"%s/password/reset/key/%s", config.Load("DEFAULT_DOMAIN"), token,
				),
			},
		},
	}

	// Send email
	ok, err := mailers.Send(em)
	if !ok || err != nil {
		return err
	}

	res["send_email"] = ok
	return c.JSON(http.StatusOK, res)
}

// PasswordResetConfirm allows the user to reset the password given a token
func (ct accountsController) PasswordResetConfirm(c echo.Context) error {
	input := new(responses.PasswordResetConfirmRequest)
	if err := c.Bind(input); err != nil {
		return apierrors.BadRequest("")
	}

	if err := input.Validate(); err != nil {
		return err
	}

	claims, err := auth.VerifyToken(input.Token, auth.JWTPasswordResetToken)
	if err != nil {
		return err
	}

	if input.Password != input.PasswordConfirm {
		return apierrors.PasswordMismatch
	}

	password, err := ct.auth.HashPassword(input.Password)
	if err != nil {
		return err
	}

	if err = ct.accountsRepo.SetPasswordUser(claims["username"].(string), password); err != nil {
		return err
	}

	res := map[string]interface{}{
		"success": true,
		"message": "restablecimiento de contraseña completado",
	}
	return c.JSON(http.StatusOK, res)
}

// accountResponse returns the access and refresh JWT tokens and the user.
//  For the user the attributes are shown depending on the role.
func (ct accountsController) accountResponse(c echo.Context, user *models.User) error {
	tokens, err := ct.auth.TokenObtainPair(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, auth.JWTResponse{
		AccessToken:  tokens["access"],
		RefreshToken: tokens["refresh"],
		User:         responses.UserResponse(user),
	})
}
