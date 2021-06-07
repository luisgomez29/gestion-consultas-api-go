package middlewares

import (
	"errors"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"

	apierrors "github.com/luisgomez29/gestion-consultas-api/api/errors"
)

// ErrorHandler crea un middleware que gestiona los pánicos y errores encontrados durante el procesamiento de
// las peticiones HTTP.
func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			buildErr := buildErrorResponse(err)
			if buildErr.(*echo.HTTPError).Code == http.StatusInternalServerError {
				log.Printf("encountered internal server error: %v", err)
			}
			return buildErr
		}
		return nil
	}
}

// buildErrorResponse construye una respuesta de error a partir de un error.
func buildErrorResponse(err error) error {
	switch err.(type) {
	case validation.Errors:
		return apierrors.InvalidInput(err.(validation.Errors))
	case *echo.HTTPError:
		switch err.(*echo.HTTPError).Code {
		case http.StatusNotFound:
			return apierrors.NotFound("")
		case http.StatusInternalServerError:
			return apierrors.InternalServerError("")
		case http.StatusForbidden:
			return apierrors.Forbidden("")
		default:
			return err
		}
	}

	var errNoRows *apierrors.ErrNoRows
	if errors.As(err, &errNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, errNoRows.Error())
	}

	return apierrors.InternalServerError("")
}
