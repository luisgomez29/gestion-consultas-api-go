package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
)

// UsersController encapsula la lógica de negocio para los usuarios
type UsersController interface {
	UsersList(c echo.Context) error
}

type usersController struct {
	repo repositories.UsersRepository
}

// NewUsersController crea un nuevo controlador de usuarios
func NewUsersController(repo repositories.UsersRepository) UsersController {
	return usersController{repo: repo}
}

func (ctrl usersController) UsersList(c echo.Context) error {
	if _, ok := auth.IsAuthenticated(c); ok {
		users, err := ctrl.repo.All()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
	return echo.NewHTTPError(http.StatusBadRequest, "INICIE SESSION")
	//users, err := ctrl.repo.All()
	//if err != nil {
	//	return err
	//}
	//return c.JSON(http.StatusOK, users)
}
