// Package routes contiene la definición de los endpoints de la API.
package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/controllers"
	"github.com/luisgomez29/gestion-consultas-api/api/middlewares"
)

// AuthHandlers establece las rutas para la autenticación
func AuthHandlers(g *echo.Group, ctrl controllers.AuthController) {
	g.POST("/signup", ctrl.SignUp)
	g.POST("/login", ctrl.Login)
}

// UsersHandlers establece las rutas para models.User
func UsersHandlers(g *echo.Group, ctrl controllers.UsersController) {
	//g.Use(middlewares.Authentication(false))
	g.GET("/users", ctrl.UsersList, middlewares.Authentication(false))
	g.GET("/users/:username", ctrl.UsersRetrieve, middlewares.Authentication(false))
}
