package models

import (
	"errors"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

const (
	// UserAdmin usuario con rol administrador
	UserAdmin = "ADMIN"

	// UserDoctor usuario con rol doctor
	UserDoctor = "DOC"

	// UserDefault usuario paciente
	UserDefault = "USR"

	// IdentificationTypeCC tipo de identificación cedula de ciudadanía (CC)
	IdentificationTypeCC = "CC"

	// IdentificationTypeCE tipo de identificación cedula de extranjería (CE)
	IdentificationTypeCE = "CE"
)

// User representa la tabla users en la DB
type User struct {
	utils.Model
	Role                 string     `json:"role"`
	FirstName            string     `json:"first_name"`
	LastName             string     `json:"last_name"`
	IdentificationType   string     `json:"identification_type"`
	IdentificationNumber string     `json:"identification_number"`
	Username             string     `json:"username"`
	Email                *string    `json:"email"`
	Password             string     `json:"password,omitempty"`
	Phone                string     `json:"phone"`
	Picture              *string    `json:"picture"`
	City                 string     `json:"city"`
	Neighborhood         *string    `json:"neighborhood"`
	Address              *string    `json:"address"`
	IsActive             bool       `json:"is_active,omitempty"`
	IsStaff              bool       `json:"is_staff,omitempty"`
	IsSuperuser          bool       `json:"is_superuser,omitempty"`
	LastLogin            *time.Time `json:"last_login"`
}

func (*User) ValidatePgError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		if pgErr.Code == pgerrcode.UniqueViolation {
			switch pgErr.ConstraintName {
			case "users_identification_number_key":
				e := errors.New("ya existe un usuario con este número de identificación")
				return validation.Errors{"identification_number": e}

			case "users_username_key":
				e := errors.New("ya existe un usuario con este nombre de usuario")
				return validation.Errors{"username": e}
			}
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err)
}
