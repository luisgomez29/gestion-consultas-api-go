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

type (
	UserRole               int
	UserIdentificationType int
)

const (
	// UserAdmin user with administrator role (ADMIN)
	UserAdmin UserRole = iota + 1

	// UserDoctor user with doctor role (DOC)
	UserDoctor

	// UserDefault patient user or default user (USER)
	UserDefault
)

const (
	// IdentificationTypeCC type of identification `cedula de ciudadanía` (CC)
	IdentificationTypeCC UserIdentificationType = iota + 1

	// IdentificationTypeCE type of identification `cedula de extranjería` (CE)
	IdentificationTypeCE
)

func (u UserRole) String() string {
	val := [...]string{"ADMIN", "DOC", "USR"}

	if u < UserAdmin || u > UserDefault {
		return ""
	}
	return val[u-1]
}

func (u UserIdentificationType) String() string {
	val := [...]string{"CC", "CE"}

	if u < IdentificationTypeCC || u > IdentificationTypeCE {
		return ""
	}
	return val[u-1]
}

// User represents the data about an user.
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
