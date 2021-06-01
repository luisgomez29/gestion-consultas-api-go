package repositories

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
)

type AuthRepository interface {
	User(username string) *models.User
}

type authRepository struct {
	conn  *pgxpool.Pool
	users UsersRepository
}

func NewAuthRepository(db *pgxpool.Pool, u UsersRepository) AuthRepository {
	return authRepository{conn: db, users: u}
}

func (r authRepository) User(username string) *models.User {
	user, _ := r.users.FindByUsername(username)
	return user
}
