package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// ConnectDB permite conectarse a la base de datos
func ConnectDB(cfg map[string]string) *pgxpool.Pool {
	// postgres://username:password@url.com:port/dbName
	DSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=America/Bogota",
		cfg["DB_USER"], cfg["DB_PWD"], cfg["DB_HOST"], cfg["DB_PORT"], cfg["DB_NAME"],
	)

	DB, err := pgxpool.Connect(context.Background(), DSN)
	if err != nil {
		log.Fatal("Failed to connect data base: ", err)
	}
	return DB
}
