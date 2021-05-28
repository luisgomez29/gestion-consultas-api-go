package errors

import (
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

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

func buildErrorResponse(err error) error {
	switch err.(type) {
	case validation.Errors:
		return InvalidInput(err.(validation.Errors))
	case *echo.HTTPError:
		switch err.(*echo.HTTPError).Code {
		case http.StatusBadRequest:
			return BadRequest("")
		case http.StatusInternalServerError:
			return InternalServerError("")
		case http.StatusNotFound:
			return NotFound("")
		case http.StatusUnauthorized:
			return Unauthorized("")
		case http.StatusForbidden:
			return Forbidden("")
		default:
			return err
		}
	}
	return InternalServerError("")
}
