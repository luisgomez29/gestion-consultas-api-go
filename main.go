package main

import (
	"fmt"
	"log"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
	"github.com/luisgomez29/gestion-consultas-api/api/database"
)

func main() {
	cfg, err := config.Load(".")

	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	database.ConnectDB(cfg.Database)
	defer database.DB.Close()

	fmt.Println("CONFIG", cfg)
}
