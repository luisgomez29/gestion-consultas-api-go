package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/database"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

func SignUp(c echo.Context) error {
	db := database.ConnectDB()
	defer db.Close()

	signup := new(responses.SignUpResponse)
	if err := c.Bind(signup); err != nil {
		return err
	}

	repo := repositories.NewAccountRepository(db)

	res, err := repo.SignUp(signup)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}
