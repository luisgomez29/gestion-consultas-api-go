package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

// Tipo de autorización
const authorizationTypeBearer = "Bearer"

// JWT tokens type
const (
	JWTAccessToken        = "access"
	JWTRefreshToken       = "refresh"
	JWTPasswordResetToken = "password_reset"
)

// Errores
var (
	errJWTMissing    = echo.NewHTTPError(http.StatusBadRequest, "token faltante o tiene un formato incorrecto")
	errJWTInvalid    = echo.NewHTTPError(http.StatusUnauthorized, "token inválido o expirado")
	errJWTimeSetting = echo.NewHTTPError(http.StatusInternalServerError, "Invalid time definition in .env file")
)

// Claims define el username del usuario y las claims estándar para generar el JWT token.
type Claims struct {
	jwt.StandardClaims

	TokenType string
	Username  string
}

// NewClaims crea la claims con valores para el Id, IssuedAt y Username.
func NewClaims(username string) *Claims {
	return &Claims{
		StandardClaims: jwt.StandardClaims{
			Id:       uuid.NewString(),
			IssuedAt: time.Now().Unix(),
		},
		Username: username,
	}
}

// GenerateToken genera un JWT token a partir de las claims.
func GenerateToken(c *Claims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token_type": c.TokenType,
		"username":   c.Username,
		"jti":        c.Id,
		"iat":        c.IssuedAt,
		"exp":        c.ExpiresAt,
	}).SignedString([]byte(config.Load("JWT_SIGNING_KEY")))

	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return token, nil
}

// VerifyToken no verifica el "token_type" de la claim. Es útil cuando se realiza la
// validación general de la firma de un token.
func VerifyToken(token string) (jwt.MapClaims, error) {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Load("JWT_SIGNING_KEY")), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, errJWTMissing
			case jwt.ValidationErrorExpired, jwt.ValidationErrorSignatureInvalid:
				return nil, errJWTInvalid
			default:
				return nil, errJWTMissing
			}
		default:
			return nil, errJWTMissing
		}
	}

	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		return claims, nil
	}

	return nil, errJWTMissing
}

// VerifyTokenWithType verifica que el token sea valido y el "token_type" de la claim.
func VerifyTokenWithType(token string, tokenType string) (jwt.MapClaims, error) {
	claims, err := VerifyToken(token)
	if err != nil {
		return nil, err
	}

	// Verificar el tipo de token
	if claims["token_type"] != tokenType {
		return nil, errJWTInvalid
	}
	return claims, nil
}

// ExtractToken obtiene el token del header de la solicitud
func ExtractToken(authzHeader string) (string, error) {
	l := len(authorizationTypeBearer)
	if len(authzHeader) > l+1 && authzHeader[:l] == authorizationTypeBearer {
		return authzHeader[l+1:], nil
	}
	return "", errJWTMissing
}

// newAccessAndRefreshClaims define las claims de los JWT token de acceso y actualización.
func newAccessAndRefreshClaims(username string) ([]*Claims, error) {
	atTime, err := utils.TimeDuration(config.Load("JWT_ACCESS_TOKEN_EXPIRATION_MINUTES"))
	if err != nil {
		return nil, errJWTimeSetting
	}

	rtTime, err := utils.TimeDuration(config.Load("JWT_REFRESH_TOKEN_EXPIRATION_DAYS"))
	if err != nil {
		return nil, errJWTimeSetting
	}

	acClaims := NewClaims(username)
	acClaims.ExpiresAt = time.Now().Add(time.Minute * atTime).Unix()
	acClaims.TokenType = JWTAccessToken

	rfClaims := NewClaims(username)
	rfClaims.ExpiresAt = time.Now().Add(time.Hour * 24 * rtTime).Unix()
	rfClaims.TokenType = JWTRefreshToken

	claims := []*Claims{acClaims, rfClaims}
	return claims, nil
}
