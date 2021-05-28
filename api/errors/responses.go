package errors

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

// InvalidInput crea un objeto JSON con la clave errors y como valor los errores de la validación de las
// estructuras
func InvalidInput(err validation.Errors) error {
	e := map[string]error{"errors": err}
	return echo.NewHTTPError(http.StatusBadRequest, e)
}

func BadRequest(msg string) error {
	if msg == "" {
		msg = "Su solicitud está en un formato incorrecto."
	}
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func NotFound(msg string) error {
	if msg == "" {
		msg = "No se ha encontrado el recurso solicitado."
	}
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

func InternalServerError(msg string) error {
	if msg == "" {
		msg = "Se ha producido un error al procesar su solicitud."
	}
	return echo.NewHTTPError(http.StatusInternalServerError, msg)
}

func Unauthorized(msg string) error {
	if msg == "" {
		msg = "No está autenticado para realizar la acción solicitada."
	}
	return echo.NewHTTPError(http.StatusUnauthorized, msg)
}

func Forbidden(msg string) error {
	if msg == "" {
		msg = "No está autorizado a realizar la acción solicitada."
	}
	return echo.NewHTTPError(http.StatusForbidden, msg)
}
