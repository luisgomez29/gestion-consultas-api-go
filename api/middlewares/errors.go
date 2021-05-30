package middlewares

import (
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// ErrorHandler crea un middleware que gestiona los pánicos y errores encontrados durante el procesamiento de
// las peticiones HTTP.
func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			res := buildErrorResponse(err)
			if res.(*echo.HTTPError).Code == http.StatusInternalServerError {
				log.Printf("encountered internal server error: %v", err)
			}
			return res
		}
		return nil
	}
}

// buildErrorResponse construye una respuesta de error a partir de un error.
func buildErrorResponse(err error) error {
	if vErr, ok := err.(validation.Errors); ok {
		return responses.InvalidInput(vErr)
	}
	return err
	//return InternalServerError("")
}
