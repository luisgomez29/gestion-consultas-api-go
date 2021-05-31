package middlewares

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
)

// Authentication es el middleware de autenticación para los endpoints. Se debe indicar si la autenticación del usuario
// con JWT token es necesaria.
//
// Si se requiere de la autenticación del usuario y el token es valido se almacena los detalles auth.AccessDetails en el
// contexto de la solicitud bajo la clave user.
//
// Si no se requiere autenticación el usuario accede a los datos públicos. Si el usuario se autentica se tratara con la
// autenticación requerida.
func Authentication(isRequired bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			if isRequired || header != "" {
				token, err := auth.VerifyToken(c.Request())
				if err != nil {
					return err
				}
				claims, err := auth.TokenPayload(token)
				log.Printf("CLAIMS %#v\n", claims)
				if err != nil {
					return err
				}
				c.Set("user", claims)
			}
			return next(c)
		}
	}
}
