package main

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
	ctrl "github.com/luisgomez29/gestion-consultas-api/api/controllers"
	"github.com/luisgomez29/gestion-consultas-api/api/database"
	repo "github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/router"
)

func main() {
	cfg, err := config.Load(".")

	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	db := database.ConnectDB(cfg.Database)
	defer db.Close()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	setupRoutes(db, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}

// setupRoutes establece las rutas disponibles de la API
func setupRoutes(db *pgxpool.Pool, e *echo.Echo) {
	// API V1
	api := e.Group("/api")
	v1 := api.Group("/v1")

	// Auth
	router.AuthHandlers(v1, ctrl.NewAuthController(repo.NewAuthRepository(db)))
	// Users
	router.UsersHandlers(v1, ctrl.NewUsersController(repo.NewUsersRepository(db)))
}
