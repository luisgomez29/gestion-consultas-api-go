// Package auth contiene los tipos y funciones relacionadas con la autenticación de usuario.
package auth

import (
	"errors"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
)

// AccessDetails representa el usuario que ha iniciado sesión
type AccessDetails struct {
	TokenUuid string
	Username  string
}

// JWTResponse es la respuesta cuando el usuario inicia sesión
type JWTResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	*models.User `json:"user"`
}

type AccessToken struct {
	Uuid      string
	Token     string
	ExpiresAt int64
}

type RefreshToken struct {
	Uuid      string
	Token     string
	ExpiresAt int64
}

// IsAuthenticated verifica si el usuario ha iniciado sesión.
// Si el usuario ha iniciado sesión retorna AccessDetails y true.
func IsAuthenticated(c echo.Context) (*AccessDetails, bool) {
	user := c.Get("user")
	if user != nil {
		return user.(*AccessDetails), true
	}
	return nil, false
}

// VerifyPassword verifica que coincidan el hash de la contraseña en la base de datos con la contraseña ingresada por
// el usuario
func VerifyPassword(hashedPassword, password string) error {
	if hashedPassword != password {
		return errors.New("las contraseñas no coinciden")
	}
	return nil
}
