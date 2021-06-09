package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/api/auth"
	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/repositories"
	"github.com/luisgomez29/gestion-consultas-api/api/responses"
)

// UsersController represents endpoints for users.
type UsersController interface {
	All(c echo.Context) error
	Get(c echo.Context) error
}

type usersController struct {
	auth      auth.Auth
	usersRepo repositories.UserRepository
}

// NewUsersController create a new users controller.
func NewUsersController(at auth.Auth, u repositories.UserRepository) UsersController {
	return usersController{auth: at, usersRepo: u}
}

func (ct usersController) All(c echo.Context) error {
	users, err := ct.usersRepo.All()
	if err != nil {
		return err
	}

	ad, _ := ct.auth.IsAuthenticated(c)

	if ad.User.Role != models.UserAdmin.String() {
		for i, user := range users {
			users[i] = responses.UserResponse(user)
		}
	}

	r := map[string][]*models.User{"results": users}
	return c.JSON(http.StatusOK, r)
}

func (ct usersController) Get(c echo.Context) error {
	user, err := ct.usersRepo.Get(c.Param("username"))
	if err != nil {
		return err
	}

	if user.Role != models.UserAdmin.String() {
		user = responses.UserResponse(user)
	}

	return c.JSON(http.StatusOK, user)
}
