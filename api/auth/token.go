package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

// Tipo de autorización
const authorizationTypeBearer = "Bearer"

// jwtSecretKet es la clave para firmar los tokens
var jwtSecretKet = []byte(config.Load("JWT_ACCESS_SECRET_KEY"))

// Errores
var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "token faltante o tiene un formato incorrecto")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusUnauthorized, "token inválido o expirado")
)

// GenerateToken genera el token de acceso
func GenerateToken(username string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"type":     "access_token",
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * utils.Hours(config.Load("JWT_ACCESS_TOKEN_EXPIRATION_DAYS"))).Unix(),
	}).SignedString(jwtSecretKet)

	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return token, nil
}

// VerifyToken verifica que el token sea valido
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKet, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, ErrJWTMissing
			case jwt.ValidationErrorExpired, jwt.ValidationErrorSignatureInvalid:
				return nil, ErrJWTInvalid
			default:
				return nil, ErrJWTMissing
			}
		default:
			return nil, ErrJWTMissing
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	}

	return nil, ErrJWTMissing
}

// ExtractToken obtiene el token del header de la solicitud
func ExtractToken(authzHeader string) (string, error) {
	l := len(authorizationTypeBearer)
	if len(authzHeader) > l+1 && authzHeader[:l] == authorizationTypeBearer {
		return authzHeader[l+1:], nil
	}
	return "", ErrJWTMissing
}
