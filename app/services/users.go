package services

import (
	"github.com/labstack/echo/v4"

	"github.com/luisgomez29/gestion-consultas-api/app/auth"
	"github.com/luisgomez29/gestion-consultas-api/app/models"
	"github.com/luisgomez29/gestion-consultas-api/app/repositories"
	"github.com/luisgomez29/gestion-consultas-api/app/resources/api/responses"
)

// UsersService encapsulates usecase logic for users.
type UsersService interface {
	All(c echo.Context) (map[string][]*models.User, error)
	Get(c echo.Context) (*models.User, error)
}

type usersService struct {
	auth      auth.Auth
	usersRepo repositories.UserRepository
}

// NewUsersService create a new users service.
func NewUsersService(at auth.Auth, u repositories.UserRepository) UsersService {
	return usersService{auth: at, usersRepo: u}
}

func (ct usersService) All(c echo.Context) (map[string][]*models.User, error) {
	users, err := ct.usersRepo.All()
	if err != nil {
		return nil, err
	}

	ad, _ := ct.auth.IsAuthenticated(c)
	return map[string][]*models.User{"results": responses.UserManyResponse(ad.User.Role, users)}, nil
}

func (ct usersService) Get(c echo.Context) (*models.User, error) {
	user, err := ct.usersRepo.Get(c.Param("username"))
	if err != nil {
		return nil, err
	}

	ad, _ := ct.auth.IsAuthenticated(c)
	return responses.UserResponse(ad.User.Role, user), nil
}
