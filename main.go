package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisgomez29/gestion-consultas-api/api/routers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/gestion-consultas-api/api/config"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	routers.SetupRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))

}
