package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
)

// DB conector a la base de datos
var DB *pgxpool.Pool

// ConnectDB permite conectarse a la base de datos
func ConnectDB(cfg config.DatabaseConfig) {
	var err error
	// postgres://username:password@url.com:port/dbName
	var DSN = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	DB, err = pgxpool.Connect(context.Background(), DSN)
	if err != nil {
		log.Fatal("Failed to connect data base: ", err)
	}
}
