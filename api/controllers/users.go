package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/gestion-consultas-api/api/database"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"net/http"
)

func UserList(c echo.Context) error {
	db := database.ConnectDB()
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	users, err := repo.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
