package router

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/controllers"
)

// AuthHandlers establece las rutas para la autenticación
func AuthHandlers(g *echo.Group, ctrl controllers.AuthController) {
	g.POST("/signup", ctrl.SignUp)
}

// UsersHandlers establece las rutas para models.User
func UsersHandlers(g *echo.Group, ctrl controllers.UsersController) {
	g.GET("/users", ctrl.UserList)
}
