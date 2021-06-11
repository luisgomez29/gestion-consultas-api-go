package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	apierrors "github.com/luisgomez29/gestion-consultas-api/api/errors"
)

// IsAdminUser creates a middleware that allows access only to models.UserAdmin users.
func IsAdminUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("user").(jwt.MapClaims)
		admin, ok := claims["admin"].(bool)
		if !admin || !ok {
			return apierrors.Forbidden("")
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
			return apierrors.Forbidden("")
		}
		return apierrors.Forbidden("")
	}
}
