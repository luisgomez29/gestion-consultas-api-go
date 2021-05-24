package repositories

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

type AccountsRepository interface {
	SignUp(res *responses.SignUpResponse) (*responses.SignUpResponse, error)
}

type accountsRepositoryDB struct {
	conn *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) AccountsRepository {
	return &accountsRepositoryDB{conn: db}
}

func (db *accountsRepositoryDB) SignUp(res *responses.SignUpResponse) (*responses.SignUpResponse, error) {
	query := `
		INSERT INTO users(role, first_name, last_name, identification_type, identification_number, username, email,
	   password, phone, city, neighborhood, address, is_active, is_staff, is_superuser)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	var user = new(models.User)

	if err := copier.Copy(user, res); err != nil {
		return nil, err
	}

	user.Role = string(models.UserRole.ADMIN)
	user.IsActive = true
	fmt.Printf("\n\n%#v\n\n", user)

	_, err := db.conn.Exec(
		context.Background(), query, user.Role, user.FirstName, user.LastName, user.IdentificationType,
		user.IdentificationNumber, user.Username, user.Email, user.Password, user.Phone, user.City, user.Neighborhood,
		user.Address, user.IsActive, user.IsStaff, user.IsSuperuser,
	)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				var msg string
				switch pgErr.ConstraintName {
				case "users_identification_number_key":
					msg = "Ya existe un usuario con este número de identificación"
				case "users_username_key":
					msg = "Ya existe un usuario con este nombre de usuario"
				}
				return nil, echo.NewHTTPError(http.StatusOK, msg)
			}
		}
	}

	return res, nil
}
