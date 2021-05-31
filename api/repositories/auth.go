package repositories

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
)

type AuthRepository interface {
	IsAuthenticated(c jwt.MapClaims) (*auth.AccessDetails, bool)
}

type authRepository struct {
	conn  *pgxpool.Pool
	users usersRepository
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return authRepository{conn: db}
}

func (r authRepository) IsAuthenticated(c jwt.MapClaims) (*auth.AccessDetails, bool) {

	//claims := c.Get("user").(jwt.MapClaims)

	fmt.Printf("USER_CLAIMS %#v\n\n", c)

	username := c["username"].(string)
	//user, _ := r.users.FindByUsername(username)
	fmt.Printf("USERNAME %#v\n\n", username)

	if username == "" {
		return nil, false
	}
	fmt.Printf("ENTRO\n\n")
	//fmt.Printf("USER_DB %#v\n\n", user)
	return &auth.AccessDetails{}, true
}
