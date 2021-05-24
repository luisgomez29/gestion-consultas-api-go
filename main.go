package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
	"github.com/luisgomez29/gestion-consultas-api/api/routers"
)

func main() {
	cfg, err := config.Load(".")

	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routers.SetupRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))

}
