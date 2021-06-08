package models

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/echo/v4"
)

// Group represents the data about an group.
type Group struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (*Group) ValidatePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return validation.Errors{"name": errors.New("ya existe un grupo con este nombre")}
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}

// UserGroup represents the data about an user-group.
type UserGroup struct {
	ID      uint `json:"id,omitempty"`
	UserID  uint `json:"user,omitempty"`
	GroupID uint `json:"group,omitempty"`
}

func (*UserGroup) ValidatePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return validation.Errors{"group_id": errors.New("el usuario ya esta asignado a este grupo")}
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}
