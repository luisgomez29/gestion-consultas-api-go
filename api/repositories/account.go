package repositories

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jinzhu/copier"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

// AccountRepository encapsulates the logic to access users from the data source.
type AccountRepository interface {
	CreateUser(res *responses.SignUpResponse) (*models.User, error)
	FindUser(username string) (*models.User, error)
}

type accountRepository struct {
	conn  *pgxpool.Pool
	group GroupRepository
}

// NewAccountRepository creates a new account repository.
func NewAccountRepository(db *pgxpool.Pool, g GroupRepository) AccountRepository {
	return accountRepository{conn: db, group: g}
}

func (r accountRepository) CreateUser(res *responses.SignUpResponse) (*models.User, error) {
	query := `
		INSERT INTO users(role, first_name, last_name, identification_type, identification_number, username, email,
		password, phone, city, neighborhood, address, is_active, is_staff, is_superuser)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING id, created_at, updated_at`

	user := new(models.User)
	if err := copier.Copy(user, res); err != nil {
		return nil, err
	}

	user.Role = models.UserDefault.String()
	user.IsActive = true

	err := r.conn.QueryRow(
		context.Background(), query, user.Role, user.FirstName, user.LastName, user.IdentificationType,
		user.IdentificationNumber, user.Username, user.Email, user.Password, user.Phone, user.City,
		user.Neighborhood, user.Address, user.IsActive, user.IsStaff, user.IsSuperuser,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, user.ValidatePgError(err)
	}

	// Add user to `Users` group
	if err := r.group.SetUser(user.ID, 1); err != nil {
		return nil, err
	}

	return user, nil
}

func (r accountRepository) FindUser(username string) (*models.User, error) {
	query := `
		SELECT id, role, first_name, last_name, identification_type, identification_number, username, email, password,
		phone, picture, city, neighborhood, address, is_active, is_staff, is_superuser, last_login, created_at, updated_at
		FROM users WHERE username = $1 AND is_active = true`

	user := new(models.User)
	if err := pgxscan.Get(context.Background(), r.conn, user, query, &username); err != nil {
		return nil, utils.ValidateErrNoRows(err, "usuario o contraseña incorrectos")
	}
	return user, nil
}
