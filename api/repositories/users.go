package repositories

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
)

type UsersRepository interface {
	All() ([]*models.User, error)
	FindById(uint) (*models.User, error)
	Create(*models.User) (*models.User, error)
	Update(uint, *models.User) (*models.User, error)
	Delete(uint) (uint, error)
}

type userRepositoryDB struct {
	conn *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) UsersRepository {
	return &userRepositoryDB{conn: db}
}

func (db *userRepositoryDB) All() ([]*models.User, error) {
	query := `
		SELECT id, role, first_name, last_name, identification_type, identification_number, username, email,
		password, phone, picture, city, neighborhood, address, is_active, is_staff, last_login, created_at, updated_at 
		FROM users;`

	var users []*models.User
	err := pgxscan.Select(context.Background(), db.conn, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *userRepositoryDB) FindById(u uint) (*models.User, error) {
	panic("implement me")
}

func (db *userRepositoryDB) Create(user *models.User) (*models.User, error) {
	//query := `
	//	INSERT INTO users(role, first_name, last_name, identification_type, identification_number, username, email,
	//    password, phone, picture, city, neighborhood, address, is_active, is_staff, last_login)
	//	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
	panic("implement me")
}

func (db *userRepositoryDB) Update(u uint, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (db *userRepositoryDB) Delete(u uint) (uint, error) {
	panic("implement me")
}
