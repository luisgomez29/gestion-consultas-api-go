// Package routes contiene la definición de los endpoints de la API.
package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/controllers"
	"github.com/luisgomez29/gestion-consultas-api/api/middlewares"
)

// AccountsHandlers establece las rutas para la autenticación y gestión de cuenta
func AccountsHandlers(g *echo.Group, ctrl controllers.AccountsController) {
	g.POST("/signup", ctrl.SignUp)
	g.POST("/login", ctrl.Login)
	g.POST("/verify-token", ctrl.VerifyToken)
}

// UsersHandlers establece las rutas para models.User
func UsersHandlers(g *echo.Group, ctrl controllers.UsersController) {
	g.Use(middlewares.Authentication(true))
	g.GET("/users", ctrl.UsersList)
	g.GET("/users/:username", ctrl.UsersRetrieve)
}
