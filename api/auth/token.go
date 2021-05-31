package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

// jwtSecretKet es la clave para firmar los tokens
var jwtSecretKet = []byte(config.Load("JWT_ACCESS_SECRET_KEY"))

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
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKet, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, responses.BadRequest("token faltante o tiene un formato incorrecto")
			case jwt.ValidationErrorExpired:
				return nil, responses.Unauthorized("su token a expirado")
			case jwt.ValidationErrorSignatureInvalid:
				return nil, responses.Unauthorized("la firma del token no coincide")
			default:
				return nil, responses.BadRequest("su token no es valido")
			}
		default:
			return nil, responses.BadRequest("su token no es valido")
		}
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	}

	return token, nil
}

// ExtractToken obtiene el token del header de la solicitud
func ExtractToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	strArr := strings.Split(auth, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// TokenPayload obtiene el payload del token
func TokenPayload(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token claims could not be obtained")
}
