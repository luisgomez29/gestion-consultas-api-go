package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ValidationErrorResponse crea un objeto JSON con la clave errors y como valor los errores de la validación de las
// estructuras
func ValidationErrorResponse(c echo.Context, err error) error {
	e := map[string]error{"errors": err}
	return c.JSON(http.StatusBadRequest, e)
}
