package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	apiErrors "github.com/luisgomez29/gestion-consultas-api/app/resources/api/errors"
)

// IsAdminUser creates a middleware that allows access only to models.UserAdmin users.
func IsAdminUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("user").(jwt.MapClaims)
		admin, ok := claims["admin"].(bool)
		if !admin || !ok {
			return apiErrors.Forbidden("")
		}
		return next(c)
	}
}

// IsAdminOrDoctorUser creates a middleware that allows access only to models.UserAdmin or models.UserDoctor users.
func IsAdminOrDoctorUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("user").(jwt.MapClaims)
		admin, ok := claims["admin"].(bool)
		doctor, ok1 := claims["doctor"].(bool)
		if ok || ok1 {
			if admin || doctor {
				return next(c)
			}
			return apiErrors.Forbidden("")
		}
		return apiErrors.Forbidden("")
	}
}
