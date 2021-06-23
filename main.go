package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/luisgomez29/gestion-consultas-api/app/middlewares"
	"github.com/luisgomez29/gestion-consultas-api/app/routers"
	"github.com/luisgomez29/gestion-consultas-api/pkg/config"
	"github.com/luisgomez29/gestion-consultas-api/pkg/database"
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
	routers.Setup(db, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Load("SERVER_PORT"))))
}
