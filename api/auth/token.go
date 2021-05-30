package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// jwtSecretKet es la clave para firmar los tokens
var jwtSecretKet = []byte(os.Getenv("JWT_ACCESS_SECRET"))

// GenerateToken genera el token de acceso
func GenerateToken(username string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
		"type":     "access_token",
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
				return nil, responses.BadRequest("Token con formato incorrecto")
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

// ExtractAccessDetails obtiene el usuario del token y retorna un AccessDetails
func ExtractAccessDetails(token *jwt.Token) (*AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &AccessDetails{Username: claims["username"].(string)}, nil
	}
	return nil, errors.New("no se pudo obtener los valores del token")
}
