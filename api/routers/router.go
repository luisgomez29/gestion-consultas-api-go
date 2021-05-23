package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/gestion-consultas-api/api/controllers"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api")

	//	Users
	users := api.Group("/users")
	users.GET("/", controllers.UserList)
}
