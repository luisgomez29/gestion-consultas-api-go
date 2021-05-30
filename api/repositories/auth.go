package repositories

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jinzhu/copier"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// AuthRepository encapsula la lógica para acceder a los usuarios desde la base de datos
type AuthRepository interface {
	SignUp(res *responses.SignUpResponse) (*models.User, error)
	Login(res *responses.LoginResponse) (*models.User, error)
}

type authRepository struct {
	conn *pgxpool.Pool
}

// NewAuthRepository crea un nuevo repositorio de autenticación
func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return authRepository{conn: db}
}

func (db authRepository) SignUp(res *responses.SignUpResponse) (*models.User, error) {
	query := `
		INSERT INTO users(role, first_name, last_name, identification_type, identification_number, username, email,
		password, phone, city, neighborhood, address, is_active, is_staff, is_superuser)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING id, created_at, updated_at`

	user := new(models.User)
	if err := copier.Copy(user, res); err != nil {
		return nil, err
	}

	user.Role = models.UserDefault
	user.IsActive = true

	err := db.conn.QueryRow(
		context.Background(), query, user.Role, user.FirstName, user.LastName, user.IdentificationType,
		user.IdentificationNumber, user.Username, user.Email, user.Password, user.Phone, user.City,
		user.Neighborhood, user.Address, user.IsActive, user.IsStaff, user.IsSuperuser,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, user.ValidatePgError(err)
	}
	return responses.UserResponse(user), nil
}

func (db authRepository) Login(res *responses.LoginResponse) (*models.User, error) {
	query := `
		SELECT id, role, first_name, last_name, identification_type, identification_number, username, email, password,
		phone, picture, city, neighborhood, address, is_active, is_staff, is_superuser, last_login, created_at, updated_at
		FROM users WHERE username = $1;`

	user := new(models.User)
	err := pgxscan.Get(context.Background(), db.conn, user, query, &res.Username)
	if err != nil {
		return nil, user.ValidatePgError(err)
	}
	return user, nil
}
