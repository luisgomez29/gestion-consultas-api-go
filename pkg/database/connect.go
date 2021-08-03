package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/pkg/config"
)

// ConnectDB connect to the database
func ConnectDB() *pgxpool.Pool {
	// postgres://username:password@url.com:port/dbName
	DSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=America/Bogota",
		config.Load("DB_USER"), config.Load("DB_PWD"), config.Load("DB_HOST"),
		config.Load("DB_PORT"), config.Load("DB_NAME"),
	)

	DB, err := pgxpool.Connect(context.Background(), DSN)
	if err != nil {
		log.Fatal("Failed to connect data base: ", err)
	}
	return DB
}
