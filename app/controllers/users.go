package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/app/services"
)

// UsersController represents endpoints for users.
type UsersController interface {
	All(c echo.Context) error
	Get(c echo.Context) error
}

type usersController struct {
	usersService services.UsersService
}

// NewUsersController create a new users controller.
func NewUsersController(us services.UsersService) UsersController {
	return usersController{usersService: us}
}

func (ct usersController) All(c echo.Context) error {
	r, err := ct.usersService.All(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (ct usersController) Get(c echo.Context) error {
	r, err := ct.usersService.Get(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}
