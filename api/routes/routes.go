// Package routes contains the definition of the API endpoints.
package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/controllers"
	"github.com/luisgomez29/gestion-consultas-api/api/middlewares"
)

// AccountsHandlers defines the endpoints for authentication and account management.
func AccountsHandlers(g *echo.Group, ctrl controllers.AccountsController) {
	g.POST("/signup", ctrl.SignUp)
	g.POST("/login", ctrl.Login)
	g.POST("/token/verify", ctrl.VerifyToken)
	g.POST("/password/reset", ctrl.PasswordReset)
	g.POST("/password/reset/confirm", ctrl.PasswordResetConfirm)
}

// UsersHandlers defines the endpoints for users management.
func UsersHandlers(g *echo.Group, ctrl controllers.UsersController) {
	g.Use(middlewares.Authentication(true))
	g.GET("/users", ctrl.All, middlewares.IsAdminOrDoctorUser)
	g.GET("/users/:username", ctrl.Get, middlewares.IsAdminOrDoctorUser)
}
