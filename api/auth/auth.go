// Package auth contiene los tipos y funciones relacionadas con la autenticación de usuario.
package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	repo "github.com/luisgomez29/gestion-consultas-api/api/repositories"
)

type Auth interface {
	IsAuthenticated(c echo.Context) (*AccessDetails, bool)
	VerifyPassword(hashedPassword, password string) error
}

type (
	auth struct {
		authRepo repo.AuthRepository
	}

	// AccessDetails representa el usuario que ha iniciado sesión
	AccessDetails struct {
		TokenUuid string
		User      *models.User
	}

	// JWTResponse es la respuesta cuando el usuario inicia sesión o se registra
	JWTResponse struct {
		Token        string       `json:"token"`
		RefreshToken string       `json:"refresh_token"`
		User         *models.User `json:"user"`
	}

	//AccessToken struct {
	//	Uuid      string
	//	Token     string
	//	ExpiresAt int64
	//}
	//
	//RefreshToken struct {
	//	Uuid      string
	//	Token     string
	//	ExpiresAt int64
	//}
)

func NewAuth(at repo.AuthRepository) Auth {
	return auth{authRepo: at}
}

// IsAuthenticated verifica si el usuario ha iniciado sesión.
// Si el usuario ha iniciado sesión retorna AccessDetails y true.
func (r auth) IsAuthenticated(c echo.Context) (*AccessDetails, bool) {
	user := c.Get("user")
	if user == nil {
		return &AccessDetails{}, false
	}

	claims := user.(jwt.MapClaims)
	username := claims["username"].(string)
	u := r.authRepo.User(username)
	return &AccessDetails{User: u}, true
}

// VerifyPassword verifica que coincidan el hash de la contraseña en la base de datos con la contraseña ingresada por
// el usuario
func (auth) VerifyPassword(hashedPassword, password string) error {
	if hashedPassword != password {
		return errors.New("las contraseñas no coinciden")
	}
	return nil
}
