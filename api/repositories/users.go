package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/luisgomez29/gestion-consultas-api/api/models"
)

type UserRepository interface {
	All() ([]*models.User, error)
	FindById(uint) (*models.User, error)
	Create(*models.User) (*models.User, error)
	Update(uint, *models.User) (*models.User, error)
	Delete(uint) (uint, error)
}

type database struct {
	conn *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &database{db}
}

func (db *database) All() ([]*models.User, error) {
	query := `SELECT id, role, first_name, last_name, identification_type, identification_number, username, email,
		password, phone, picture, city, neighborhood, address, is_active, is_staff, last_login, created_at, updated_at 
		FROM users;`

	var users []*models.User
	err := pgxscan.Select(context.Background(), db.conn, &users, query)
	if err != nil {
		return nil, err
	}
	fmt.Printf("USERS => %v", users[0])
	us, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n\nJSON %s", string(us))
	return users, nil
}

func (db *database) FindById(u uint) (*models.User, error) {
	panic("implement me")
}

func (db *database) Create(user *models.User) (*models.User, error) {
	panic("implement me")
}

func (db *database) Update(u uint, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (db *database) Delete(u uint) (uint, error) {
	panic("implement me")
}
