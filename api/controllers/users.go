package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
)

// UsersController encapsula la lógica de negocio para los usuarios
type UsersController interface {
	UserList(c echo.Context) error
}

type usersController struct {
	repo repositories.UsersRepository
}

// NewUsersController crea un nuevo controlador de usuarios
func NewUsersController(repo repositories.UsersRepository) UsersController {
	return usersController{repo: repo}
}

func (ctrl usersController) UserList(c echo.Context) error {
	users, err := ctrl.repo.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
