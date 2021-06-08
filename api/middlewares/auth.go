package middlewares

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
)

// Authentication is the authentication middleware for endpoints. It must be indicated if user
// authentication with JWT token is required.
//
// If user authentication is required and the token is valid, the auth.AccessDetails are stored in the
// context of the request under the user key.
//
//If authentication is not required, the user accesses the public data. If the user authenticates,
// it will deal with the required authentication.
func Authentication(isRequired bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authzHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			if isRequired || authzHeader != "" {

				tokenString, err := auth.ExtractToken(authzHeader)
				if err != nil {
					return err
				}

				claims, err := auth.VerifyTokenWithType(tokenString, auth.JWTAccessToken)
				if err != nil {
					return err
				}

				c.Set("user", claims)
			}
			return next(c)
		}
	}
}
