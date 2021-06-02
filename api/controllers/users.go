package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	"github.com/luisgomez29/gestion-consultas-api/api/models"
	repo "github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// UsersController encapsula la lógica de negocio para los usuarios
type UsersController interface {
	UsersList(c echo.Context) error
	UsersRetrieve(c echo.Context) error
}

type usersController struct {
	auth      auth.Auth
	usersRepo repo.UsersRepository
}

// NewUsersController crea un nuevo controlador de usuarios
func NewUsersController(at auth.Auth, u repo.UsersRepository) UsersController {
	return usersController{auth: at, usersRepo: u}
}

func (ct usersController) UsersList(c echo.Context) error {
	users, err := ct.usersRepo.All()
	if err != nil {
		return err
	}
	r := map[string][]*models.User{"results": users}
	return c.JSON(http.StatusOK, r)
}

func (ct usersController) UsersRetrieve(c echo.Context) error {
	user, err := ct.usersRepo.FindByUsername(c.Param("username"))
	if err != nil {
		return err
	}

	if user.Role != models.UserAdmin.String() {
		user = responses.UserResponse(user)
	}

	return c.JSON(http.StatusOK, user)
}
