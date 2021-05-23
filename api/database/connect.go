package database

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/config"
)

// ConnectDB permite conectarse a la base de datos
func ConnectDB() *pgxpool.Pool {
	cfg := new(config.Config)
	err := viper.Unmarshal(cfg)

	// postgres://username:password@url.com:port/dbName
	DSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=America/Bogota",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name,
	)

	DB, err := pgxpool.Connect(context.Background(), DSN)
	if err != nil {
		log.Fatal("Failed to connect data base: ", err)
	}
	return DB
}
