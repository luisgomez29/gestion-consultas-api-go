package routers

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/controllers"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api")

	// API V1
	v1 := api.Group("/v1")

	// Accounts
	accounts := v1.Group("/accounts")
	accounts.POST("/signup", controllers.SignUp)

	//	Users
	users := v1.Group("/users")
	users.GET("/", controllers.UserList)
}
