package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// UsersController encapsula la lógica de negocio para los usuarios
type UsersController interface {
	UsersList(c echo.Context) error
	UsersRetrieve(c echo.Context) error
}

type usersController struct {
	usersRepo repositories.UsersRepository
}

// NewUsersController crea un nuevo controlador de usuarios
func NewUsersController(repo repositories.UsersRepository) UsersController {
	return usersController{usersRepo: repo}
}

func (ct usersController) UsersList(c echo.Context) error {
	if _, ok := auth.IsAuthenticated(c); ok {
		users, err := ct.usersRepo.All()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
	return echo.NewHTTPError(http.StatusBadRequest, "INICIE SESSION")
	//users, err := ct.usersRepo.All()
	//if err != nil {
	//	return err
	//}
	//return c.JSON(http.StatusOK, users)
}

func (ct usersController) UsersRetrieve(c echo.Context) error {
	user, err := ct.usersRepo.FindByUsername(c.Param("username"))

	u, ok := auth.IsAuthenticated(c)
	if ok {
		fmt.Printf("USER AUTH %#v\n", u)
	}
	if err != nil {
		return err
	}

	if user.Role != models.UserAdmin.String() {
		user = responses.UserResponse(user)
	}

	return c.JSON(http.StatusOK, user)
}
