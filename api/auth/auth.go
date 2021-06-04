// Package auth contiene los tipos y funciones relacionadas con la autenticación de usuario.
package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	repo "github.com/luisgomez29/gestion-consultas-api/api/repositories"
)

type Auth interface {
	// HashPassword retorna el hash argon2id de las contraseña
	HashPassword(password string) (string, error)

	// VerifyPassword verifica que coincidan el hash de la contraseña en la base de datos con la
	// contraseña ingresada por el usuario.
	VerifyPassword(password, hashedPassword string) (bool, error)

	// UsernameFromContext obtiene el username del usuario de la solicitud.
	UsernameFromContext(c echo.Context) string

	// IsAuthenticated verifica si el usuario ha iniciado sesión.
	// Si el usuario ha iniciado sesión retorna AccessDetails y true.
	IsAuthenticated(c echo.Context) (*AccessDetails, bool)

	// UserPermissions obtiene los permisos que el usuario tiene en `user_permissions`.
	UserPermissions(u *models.User) ([]*models.Permission, error)

	// GroupPermissions obtiene los permisos que el usuario tiene de los grupos a los que pertenece.
	GroupPermissions(u *models.User) ([]*models.Permission, error)

	// AllPermissions obtiene todos los permisos del usuario ya sean de grupo o de usuario.
	AllPermissions(u *models.User) ([]*models.Permission, error)

	// HasPermission verifica si el usuario tiene un permiso.
	HasPermission(u *models.User, perm string) (bool, error)
}

type (
	auth struct {
		authRepo repo.AuthRepository
	}

	// AccessDetails representa el usuario que ha iniciado sesión.
	AccessDetails struct {
		TokenUuid string
		User      *models.User
	}

	// JWTResponse es la respuesta cuando el usuario inicia sesión o se registra.
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

func (r auth) UsernameFromContext(c echo.Context) string {
	user := c.Get("user")
	if user == nil {
		return ""
	}
	claims := user.(jwt.MapClaims)
	return claims["username"].(string)
}

func (r auth) IsAuthenticated(c echo.Context) (*AccessDetails, bool) {
	username := r.UsernameFromContext(c)
	if username == "" {
		return &AccessDetails{}, false
	}

	u := r.authRepo.UserLoggedIn(username)
	return &AccessDetails{User: u}, true
}

func (r auth) UserPermissions(u *models.User) ([]*models.Permission, error) {
	if u.Role == models.UserAdmin.String() {
		return r.authRepo.AllPermissions()
	}
	return r.authRepo.UserPermissions(u.Username)
}

func (r auth) GroupPermissions(u *models.User) ([]*models.Permission, error) {
	if u.Role == models.UserAdmin.String() {
		return r.authRepo.AllPermissions()
	}
	return r.authRepo.GroupPermissions(u.Username)
}

func (r auth) AllPermissions(u *models.User) ([]*models.Permission, error) {
	if u.Role == models.UserAdmin.String() {
		return r.authRepo.AllPermissions()
	}

	uPerms, err := r.UserPermissions(u)
	if err != nil {
		return nil, err
	}

	gPerms, err := r.GroupPermissions(u)
	if err != nil {
		return nil, err
	}

	uPerms = append(uPerms, gPerms...)
	return uPerms, nil
}

func (r auth) HasPermission(u *models.User, perm string) (bool, error) {
	perms, err := r.AllPermissions(u)
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
