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

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) error {
	if msg == "" {
		msg = "su solicitud está en un formato incorrecto."
	}
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

// NotFound crea una nueva respuesta de error que representa un error de recurso no encontrado (HTTP 404)
func NotFound(msg string) error {
	if msg == "" {
		msg = "no se ha encontrado el recurso solicitado."
	}
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

// InternalServerError crea una nueva respuesta de error que representa un error interno del servidor (HTTP 500)
func InternalServerError(msg string) error {
	if msg == "" {
		msg = "se ha producido un error al procesar su solicitud."
	}
	return echo.NewHTTPError(http.StatusInternalServerError, msg)
}

// Unauthorized crea una nueva respuesta de error que representa un fallo de autenticación/autorización (HTTP 401)
func Unauthorized(msg string) error {
	if msg == "" {
		msg = "no está autenticado para realizar la acción solicitada."
	}
	return echo.NewHTTPError(http.StatusUnauthorized, msg)
}

// Forbidden crea una nueva respuesta de error que representa un fallo de autorización (HTTP 403)
func Forbidden(msg string) error {
	if msg == "" {
		msg = "no está autorizado a realizar la acción solicitada."
	}
	return echo.NewHTTPError(http.StatusForbidden, msg)
}
