package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/app/models"
	"github.com/luisgomez29/gestion-consultas-api/app/utils"
	"github.com/luisgomez29/gestion-consultas-api/pkg/config"
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
	User      *models.User
}

// NewClaims create the claims with values for the Id, IssuedAt and Username.
func NewClaims(u *models.User) *Claims {
	return &Claims{
		StandardClaims: jwt.StandardClaims{
			Id:       uuid.NewString(),
			IssuedAt: time.Now().Unix(),
		},
		User: u,
	}
}

// GenerateToken generate a JWT token from the claims.
func GenerateToken(c *Claims) (string, error) {
	claims := jwt.MapClaims{
		"token_type": c.TokenType,
		"username":   c.User.Username,
		"jti":        c.Id,
		"iat":        c.IssuedAt,
		"exp":        c.ExpiresAt,
	}

	// Add role
	if c.User.Role == models.UserAdmin.String() {
		claims["admin"] = true
	} else if c.User.Role == models.UserDoctor.String() {
		claims["doctor"] = true
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(
		[]byte(config.Load("JWT_SIGNING_KEY")),
	)

	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return token, nil
}

// VerifyToken verify that the token is valid. If a value is assigned to the `tokenType` parameter,
// the `token_type` of the claim is verified.
func VerifyToken(token string, tokenType ...string) (jwt.MapClaims, error) {
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

	claims, ok := tk.Claims.(jwt.MapClaims)
	if !ok || !tk.Valid {
		return nil, errJWTMissing
	}

	// Verify token type
	if tokenType != nil {
		if claims["token_type"] != tokenType[0] {
			return nil, errJWTInvalid
		}
		return claims, nil
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
func newAccessAndRefreshClaims(u *models.User) ([]*Claims, error) {
	atTime, err := utils.TimeDuration(config.Load("JWT_ACCESS_TOKEN_EXPIRATION_MINUTES"))
	if err != nil {
		return nil, errJWTimeSetting
	}

	rtTime, err := utils.TimeDuration(config.Load("JWT_REFRESH_TOKEN_EXPIRATION_DAYS"))
	if err != nil {
		return nil, errJWTimeSetting
	}

	acClaims := NewClaims(u)
	acClaims.ExpiresAt = time.Now().Add(time.Minute * atTime).Unix()
	acClaims.TokenType = JWTAccessToken

	rfClaims := NewClaims(u)
	rfClaims.ExpiresAt = time.Now().Add(time.Hour * 24 * rtTime).Unix()
	rfClaims.TokenType = JWTRefreshToken

	return []*Claims{acClaims, rfClaims}, nil
}
