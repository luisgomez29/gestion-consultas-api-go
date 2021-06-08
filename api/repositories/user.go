package repositories

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

// UserRepository encapsulates the logic to access users from the data source.
type UserRepository interface {
	All() ([]*models.User, error)
	Get(username string) (*models.User, error)
	Create(u *models.User) (*models.User, error)
	Update(id uint, u *models.User) (*models.User, error)
	Delete(id uint) (uint, error)
}

//userRepository persists users in database.
type userRepository struct {
	conn *pgxpool.Pool
}

// NewUserRepository  creates a new user repository.
func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return userRepository{conn: db}
}

func (r userRepository) All() ([]*models.User, error) {
	query := `
		SELECT id, role, first_name, last_name, identification_type, identification_number, username, email, phone,
		picture, city, neighborhood, address, is_active, is_staff, is_superuser, last_login, created_at, updated_at
		FROM users WHERE is_active = true ORDER BY created_at DESC`

	var users []*models.User
	if err := pgxscan.Select(context.Background(), r.conn, &users, query); err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepository) Get(username string) (*models.User, error) {
	query := `
		SELECT id, role, first_name, last_name, identification_type, identification_number, username, email, phone,
		picture, city, neighborhood, address, is_active, is_staff, is_superuser, last_login, created_at, updated_at
		FROM users WHERE username = $1 AND is_active = true`

	user := new(models.User)
	if err := pgxscan.Get(context.Background(), r.conn, user, query, &username); err != nil {
		return nil, utils.ValidateErrNoRows(err, "usuario no encontrado")
	}
	return user, nil
}

func (r userRepository) Create(u *models.User) (*models.User, error) {
	//query := `
	//	INSERT INTO user(role, first_name, last_name, identification_type, identification_number, username, email,
	//    password, phone, picture, city, neighborhood, address, is_active, is_staff, last_login)
	//	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
	panic("implement me")
}

func (r userRepository) Update(u uint, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (r userRepository) Delete(u uint) (uint, error) {
	panic("implement me")
}
