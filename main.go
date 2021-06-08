package main

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	"github.com/luisgomez29/gestion-consultas-api/api/config"
	"github.com/luisgomez29/gestion-consultas-api/api/controllers"
	"github.com/luisgomez29/gestion-consultas-api/api/database"
	"github.com/luisgomez29/gestion-consultas-api/api/middlewares"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/routes"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	e := echo.New()

	// Middlewares
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.Secure(),
		middlewares.ErrorHandler,
	)

	// Routes
	setupRoutes(db, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Load("SERVER_PORT"))))
}

// setupRoutes sets the available API endpoints.
func setupRoutes(db *pgxpool.Pool, e *echo.Echo) {
	// API V1
	api := e.Group("/api")
	v1 := api.Group("/v1")

	// Repositories
	permRepo := repositories.NewPermissionRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	usersRepo := repositories.NewUserRepository(db)
	authRepo := repositories.NewAuthRepository(db, permRepo, usersRepo)
	accountsRepo := repositories.NewAccountRepository(db, groupRepo)

	// Authentication service
	authn := auth.NewAuth(authRepo)

	// Accounts
	routes.AccountsHandlers(v1, controllers.NewAccountsController(authn, accountsRepo))

	// Users
	routes.UsersHandlers(v1, controllers.NewUsersController(authn, usersRepo))
}
