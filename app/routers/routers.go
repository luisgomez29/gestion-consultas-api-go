package routers

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/app/auth"
	"github.com/luisgomez29/gestion-consultas-api/app/controllers"
	"github.com/luisgomez29/gestion-consultas-api/app/repositories"
	"github.com/luisgomez29/gestion-consultas-api/app/services"
)

// Setup sets the available API endpoints.
func Setup(db *pgxpool.Pool, e *echo.Echo) {
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

	// Services
	usersService := services.NewUsersService(authn, usersRepo)
	accountsService := services.NewAccountsService(authn, accountsRepo)

	// Routes
	Accounts(v1, controllers.NewAccountsController(authn, accountsService))
	Users(v1, controllers.NewUsersController(usersService))
}
