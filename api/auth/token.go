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

// Authorization type
const authorizationTypeBearer = "Bearer"

// JWT tokens type
const (
	JWTAccessToken        = "access"
	JWTRefreshToken       = "refresh"
	JWTPasswordResetToken = "password_reset"
)

// Errors
var (
	errJWTMissing    = echo.NewHTTPError(http.StatusBadRequest, "token faltante o tiene un formato incorrecto")
	errJWTInvalid    = echo.NewHTTPError(http.StatusUnauthorized, "token inválido o expirado")
	errJWTimeSetting = echo.NewHTTPError(http.StatusInternalServerError, "Invalid time definition in .env file")
)

// Claims defines the username of the user and the standard claims to generate the JWT token.
type Claims struct {
	jwt.StandardClaims

	TokenType string
	Username  string
}

// NewClaims create the claims with values for the Id, IssuedAt and Username.
func NewClaims(username string) *Claims {
	return &Claims{
		StandardClaims: jwt.StandardClaims{
			Id:       uuid.NewString(),
			IssuedAt: time.Now().Unix(),
		},
		Username: username,
	}
}

// GenerateToken generate a JWT token from the claims.
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

// VerifyToken does not verify the "token_type" claim. This is useful when performing general
// validation of a token's signature.
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

// VerifyTokenWithType verify that the token is valid and the "token_type" of the claim.
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

// ExtractToken get the token from the request header.
func ExtractToken(authzHeader string) (string, error) {
	l := len(authorizationTypeBearer)
	if len(authzHeader) > l+1 && authzHeader[:l] == authorizationTypeBearer {
		return authzHeader[l+1:], nil
	}
	return "", errJWTMissing
}

// newAccessAndRefreshClaims defines the claims of the access and refresh JWT token.
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
