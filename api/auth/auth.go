// Package auth contains the types and functions related to user authentication.
package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
)

type Auth interface {
	// HashPassword get the argon2id hash of the password
	HashPassword(password string) (string, error)

	// VerifyPassword verifies that the password hash matches the password entered by the user.
	VerifyPassword(password, hashedPassword string) (bool, error)

	// TokenObtainPair generates the JWT access and refresh tokens.
	TokenObtainPair(username string) (map[string]string, error)

	// VerifyToken verify that the token is valid.
	VerifyToken(token string) (map[string]interface{}, error)

	// UsernameFromContext get the username of the request user.
	UsernameFromContext(c echo.Context) string

	// IsAuthenticated check if the user is logged in. If the user is logged in, it returns AccessDetails and true.
	IsAuthenticated(c echo.Context) (*AccessDetails, bool)

	// UserPermissions get the permissions that the user has in `user_permissions`.
	UserPermissions(u *models.User) ([]*models.Permission, error)

	// UserGroupPermissions get the permissions that this user has through their groups..
	UserGroupPermissions(u *models.User) ([]*models.Permission, error)

	// UserAllPermissions Get all the user's group and user permissions.
	UserAllPermissions(u *models.User) ([]*models.Permission, error)

	// HasPermission check if the user has the specified permission.
	HasPermission(u *models.User, perm string) (bool, error)
}

type (
	// AccessDetails represents the user who is logged in.
	AccessDetails struct {
		TokenUuid string
		User      *models.User
	}

	// JWTResponse is the response when the user logs in or register.
	JWTResponse struct {
		AccessToken  string       `json:"access_token"`
		RefreshToken string       `json:"refresh_token"`
		User         *models.User `json:"user"`
	}
)

type auth struct {
	authRepo repositories.AuthRepository
}

func NewAuth(at repositories.AuthRepository) Auth {
	return auth{authRepo: at}
}

func (auth) HashPassword(password string) (string, error) {
	c := &passwordConfig{
		memory:      102400,
		iterations:  2,
		parallelism: 8,
		saltLength:  16,
		keyLength:   16,
	}
	return generatePassword(c, password)
}

func (auth) VerifyPassword(password, hashedPassword string) (bool, error) {
	return comparePasswordAndHash(password, hashedPassword)
}

func (auth) TokenObtainPair(username string) (map[string]string, error) {
	claims, err := newAccessAndRefreshClaims(username)
	if err != nil {
		return nil, err
	}

	accessToken, err := GenerateToken(claims[0])
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateToken(claims[1])
	if err != nil {
		return nil, err
	}

	tokens := map[string]string{
		"access":  accessToken,
		"refresh": refreshToken,
	}
	return tokens, nil
}

func (auth) VerifyToken(token string) (map[string]interface{}, error) {
	claims, err := VerifyToken(token)
	if err != nil {
		return nil, err
	}

	r := map[string]interface{}{
		"success": true,
		"payload": claims,
	}
	return r, nil
}

func (a auth) UsernameFromContext(c echo.Context) string {
	user := c.Get("user")
	if user == nil {
		return ""
	}
	claims := user.(jwt.MapClaims)
	return claims["username"].(string)
}

func (a auth) IsAuthenticated(c echo.Context) (*AccessDetails, bool) {
	username := a.UsernameFromContext(c)
	if username == "" {
		return &AccessDetails{}, false
	}

	u := a.authRepo.UserLoggedIn(username)
	return &AccessDetails{User: u}, true
}

func (a auth) UserPermissions(u *models.User) ([]*models.Permission, error) {
	if u.Role == models.UserAdmin.String() {
		return a.authRepo.AllPermissions()
	}
	return a.authRepo.UserPermissions(u.Username)
}

func (a auth) UserGroupPermissions(u *models.User) ([]*models.Permission, error) {
	if u.Role == models.UserAdmin.String() {
		return a.authRepo.AllPermissions()
	}
	return a.authRepo.UserGroupPermissions(u.Username)
}

func (a auth) UserAllPermissions(u *models.User) ([]*models.Permission, error) {
	if u.Role == models.UserAdmin.String() {
		return a.authRepo.AllPermissions()
	}

	uPerms, err := a.UserPermissions(u)
	if err != nil {
		return nil, err
	}

	gPerms, err := a.UserGroupPermissions(u)
	if err != nil {
		return nil, err
	}

	uPerms = append(uPerms, gPerms...)
	return uPerms, nil
}

func (a auth) HasPermission(u *models.User, perm string) (bool, error) {
	perms, err := a.UserAllPermissions(u)
	if err != nil {
		return false, err
	}

	for _, p := range perms {
		if p.Codename == perm {
			return true, nil
		}
	}
	return false, nil
}
