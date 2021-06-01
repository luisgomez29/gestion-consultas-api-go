package main

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	"github.com/luisgomez29/gestion-consultas-api/api/config"
	ctrl "github.com/luisgomez29/gestion-consultas-api/api/controllers"
	"github.com/luisgomez29/gestion-consultas-api/api/database"
	"github.com/luisgomez29/gestion-consultas-api/api/middlewares"
	repo "github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/routes"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	e := echo.New()

	// Middleware
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.Secure(),
		middlewares.ErrorHandler,
	)

	// Auth
	//ctrl.NewAuthController(repo.NewAuthRepository(db))

	// Routes
	setupRoutes(db, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Load("SERVER_PORT"))))
}

// setupRoutes establece las rutas disponibles de la API
func setupRoutes(db *pgxpool.Pool, e *echo.Echo) {
	// API V1
	api := e.Group("/api")
	v1 := api.Group("/v1")

	usersRepo := repo.NewUsersRepository(db)
	authRepo := repo.NewAuthRepository(db, usersRepo)
	accountsRepo := repo.NewAccountsRepository(db)

	// Auth
	ath := auth.NewAuth(authRepo)

	// Accounts
	routes.AccountsHandlers(v1, ctrl.NewAccountsController(accountsRepo, ath))

	// Users
	routes.UsersHandlers(v1, ctrl.NewUsersController(usersRepo, ath))
}
