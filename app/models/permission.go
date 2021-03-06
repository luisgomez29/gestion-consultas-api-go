package models

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/echo/v4"
)

// Permission represents the data about an permission.
type Permission struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Codename      string `json:"codename"`
	ContentTypeID uint   `json:"content_type,omitempty"`
}

func (*Permission) ValidatePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return validation.Errors{"codename": errors.New("ya existe un permiso con este codename")}
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err)
}
