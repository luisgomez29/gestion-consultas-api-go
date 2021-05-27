package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jinzhu/copier"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// AuthRepository encapsula la lógica para acceder a los usuarios desde la base de datos
type AuthRepository interface {
	SignUp(res *responses.SignUpResponse) (*responses.SignUpResponse, error)
}

type authRepository struct {
	conn *pgxpool.Pool
}

// NewAuthRepository crea un nuevo repositorio de autenticación
func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return authRepository{conn: db}
}

func (db authRepository) SignUp(res *responses.SignUpResponse) (*responses.SignUpResponse, error) {
	query := `
		INSERT INTO users(role, first_name, last_name, identification_type, identification_number, username, email,
		password, phone, city, neighborhood, address, is_active, is_staff, is_superuser)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	user := new(models.User)

	if err := copier.Copy(user, res); err != nil {
		return nil, err
	}
	user.Role = models.UserDefault
	user.IsActive = true
	fmt.Printf("\n\n%#v\n\n", user)

	_, err := db.conn.Exec(
		context.Background(), query, user.Role, user.FirstName, user.LastName, user.IdentificationType,
		user.IdentificationNumber, user.Username, user.Email, user.Password, user.Phone, user.City,
		user.Neighborhood, user.Address, user.IsActive, user.IsStaff, user.IsSuperuser,
	)

	if err != nil {
		return nil, user.ValidatePgError(err)
	}
	return res, nil
}
