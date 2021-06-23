// Package routers contains the definition of the API endpoints.
package routers

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/app/controllers"
	"github.com/luisgomez29/gestion-consultas-api/app/middlewares"
)

// Accounts defines the endpoints for authentication and account management.
func Accounts(g *echo.Group, ctrl controllers.AccountsController) {
	g.POST("/signup", ctrl.SignUp)
	g.POST("/login", ctrl.Login)
	g.POST("/token/verify", ctrl.VerifyToken)
	g.POST("/password/reset", ctrl.PasswordReset)
	g.POST("/password/reset/confirm", ctrl.PasswordResetConfirm)
}

// Users defines the endpoints for users management.
func Users(g *echo.Group, ctrl controllers.UsersController) {
	g.Use(middlewares.Authentication(true))
	g.GET("/users", ctrl.All, middlewares.IsAdminOrDoctorUser)
	g.GET("/users/:username", ctrl.Get, middlewares.IsAdminOrDoctorUser)
}
