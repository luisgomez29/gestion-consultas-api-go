package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
)

// ConnectDB permite conectarse a la base de datos
func ConnectDB(cfg config.DatabaseConfig) *pgxpool.Pool {
	// postgres://username:password@url.com:port/dbName
	DSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=America/Bogota",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	DB, err := pgxpool.Connect(context.Background(), DSN)
	if err != nil {
		log.Fatal("Failed to connect data base: ", err)
	}
	return DB
}
