package services

import (
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/app/auth"
	"github.com/luisgomez29/gestion-consultas-api/app/models"
	"github.com/luisgomez29/gestion-consultas-api/app/repositories"
	apiErrors "github.com/luisgomez29/gestion-consultas-api/app/resources/api/errors"
	"github.com/luisgomez29/gestion-consultas-api/app/resources/api/requests"
	"github.com/luisgomez29/gestion-consultas-api/app/resources/api/responses"
	"github.com/luisgomez29/gestion-consultas-api/app/resources/mailer"
	"github.com/luisgomez29/gestion-consultas-api/pkg/config"
)

type AccountsService interface {
	SignUp(input *requests.SignUpRequest) (auth.JWTResponse, error)
	Login(input *requests.LoginRequest) (auth.JWTResponse, error)
	VerifyToken(token string) (echo.Map, error)
	PasswordReset(username string) (echo.Map, error)
	PasswordResetConfirm(username, password string) (echo.Map, error)
}

type accountsService struct {
	auth         auth.Auth
	accountsRepo repositories.AccountRepository
}

// NewAccountsService create a new accounts service.
func NewAccountsService(at auth.Auth, u repositories.AccountRepository) AccountsService {
	return accountsService{auth: at, accountsRepo: u}
}

func (s accountsService) SignUp(input *requests.SignUpRequest) (auth.JWTResponse, error) {
	// Generating password Hash
	hash, err := s.auth.HashPassword(input.Password)
	if err != nil {
		return auth.JWTResponse{}, err
	}

	input.Password = hash
	user, err := s.accountsRepo.CreateUser(input)
	if err != nil {
		return auth.JWTResponse{}, err
	}

	return s.TokensAndUser(user)
}

func (s accountsService) Login(input *requests.LoginRequest) (auth.JWTResponse, error) {
	user, err := s.accountsRepo.FindUser(input.Username)
	if err != nil {
		return auth.JWTResponse{}, err
	}

	match, err := s.auth.VerifyPassword(input.Password, user.Password)
	if !match || err != nil {
		return auth.JWTResponse{}, apiErrors.Unauthorized("la contraseña ingresada es incorrecta")
	}

	if err = s.accountsRepo.UpdateLastLogin(user.Username); err != nil {
		return auth.JWTResponse{}, err
	}

	return s.TokensAndUser(user)
}

func (s accountsService) VerifyToken(token string) (echo.Map, error) {
	claims, err := auth.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	return echo.Map{
		"success": true,
		"payload": claims,
	}, nil
}

func (s accountsService) PasswordReset(username string) (echo.Map, error) {
	user, err := s.accountsRepo.FindUser(username)
	if err != nil {
		var apiErr *apiErrors.ErrNoRows
		if errors.As(err, &apiErr) {
			return nil, apiErrors.NewErrNoRows("el usuario no está asignado a ninguna cuenta")
		}
		return nil, err
	}

	// Response
	res := echo.Map{
		"send_email": false,
		"username":   user.Username,
		"email":      user.Email,
	}

	// Verify if the user does not have email
	if user.Email == nil {
		return res, nil
	}

	// Generate token
	claims := auth.NewClaims(user)
	claims.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()
	claims.TokenType = auth.JWTPasswordResetToken

	token, err := auth.GenerateToken(claims)
	if err != nil {
		return nil, err
	}

	// Email message
	em := &mailer.EmailMessage{
		To:      mail.Address{Name: user.FirstName, Address: *user.Email},
		Subject: "Solicitud de recuperación de contraseña",
		Template: mailer.Template{
			Name: "accounts/password_reset_key_message.html",
			Context: echo.Map{
				"currentSite":   "Gestión consultas",
				"userFirstName": user.FirstName,
				"passwordResetUrl": fmt.Sprintf(
					"%s/password/reset/key/%s", config.Load("DEFAULT_DOMAIN"), token,
				),
			},
		},
	}

	// Send email
	ok, err := mailer.Send(em)
	if !ok || err != nil {
		return nil, err
	}

	res["send_email"] = ok
	return res, nil
}

func (s accountsService) PasswordResetConfirm(username, password string) (echo.Map, error) {
	pwd, err := s.auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	if err = s.accountsRepo.SetPasswordUser(username, pwd); err != nil {
		return nil, err
	}

	return echo.Map{
		"success": true,
		"message": "restablecimiento de contraseña completado",
	}, nil
}

// TokensAndUser returns the access and refresh JWT tokens and the user.
// For the user the attributes are shown depending on the role.
func (s accountsService) TokensAndUser(user *models.User) (auth.JWTResponse, error) {
	tokens, err := s.auth.TokenObtainPair(user)
	if err != nil {
		return auth.JWTResponse{}, err
	}

	return auth.JWTResponse{
		AccessToken:  tokens["access"],
		RefreshToken: tokens["refresh"],
		User:         responses.UserResponse(user.Role, user),
	}, nil
}
