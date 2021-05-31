package repositories

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
)

// UsersRepository encapsula la lógica para acceder a los usuarios desde la base de datos
type UsersRepository interface {
	All() ([]*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(u *models.User) (*models.User, error)
	Update(id uint, u *models.User) (*models.User, error)
	Delete(id uint) (uint, error)
}

type usersRepository struct {
	conn *pgxpool.Pool
}

// NewUsersRepository crea un nuevo repositorio de usuarios
func NewUsersRepository(db *pgxpool.Pool) UsersRepository {
	return usersRepository{conn: db}
}

func (r usersRepository) All() ([]*models.User, error) {
	query := `
		SELECT id, role, first_name, last_name, identification_type, identification_number, username, email, phone,
		picture, city, neighborhood, address, is_active, is_staff, is_superuser, last_login, created_at, updated_at
		FROM users;`

	var users []*models.User
	err := pgxscan.Select(context.Background(), r.conn, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r usersRepository) FindByUsername(username string) (*models.User, error) {
	query := `
		SELECT id, role, first_name, last_name, identification_type, identification_number, username, email, phone,
		picture, city, neighborhood, address, is_active, is_staff, is_superuser, last_login, created_at, updated_at
		FROM users WHERE username = $1;`

	user := new(models.User)
	err := pgxscan.Get(context.Background(), r.conn, user, query, &username)

	if err != nil {
		return nil, user.NotFound(err, "usuario no encontrado")
	}
	return user, nil
}

func (r usersRepository) Create(u *models.User) (*models.User, error) {
	//query := `
	//	INSERT INTO users(role, first_name, last_name, identification_type, identification_number, username, email,
	//    password, phone, picture, city, neighborhood, address, is_active, is_staff, last_login)
	//	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
	panic("implement me")
}

func (r usersRepository) Update(u uint, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (r usersRepository) Delete(u uint) (uint, error) {
	panic("implement me")
}
