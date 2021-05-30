package responses

import (
	"time"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

// UserResponse lista los campos a retornar para los tipos de usuario
func UserResponse(u *models.User) *models.User {
	if u.Role == models.UserDefault || u.Role == models.UserDoctor {
		u.IsActive = false
	}
	u.Password = ""
	return u
}

// UserBaseResponse contiene los campos del usuario que se retornan sin importar el rol
type UserBaseResponse struct {
	FirstName            string  `json:"first_name"`
	LastName             string  `json:"last_name"`
	IdentificationType   string  `json:"identification_type"`
	IdentificationNumber string  `json:"identification_number"`
	Username             string  `json:"username"`
	Email                *string `json:"email"`
	Phone                string  `json:"phone"`
	City                 string  `json:"city"`
	Neighborhood         *string `json:"neighborhood"`
	Address              *string `json:"address"`
}

type UserDefaultResponse struct {
	*UserBaseResponse
	utils.Model
	Role      string     `json:"role"`
	LastLogin *time.Time `json:"last_login"`
}

type UserDoctorResponse struct {
	*UserDefaultResponse
}

type UserAdminResponse struct {
	*UserDefaultResponse
	IsActive    bool `json:"is_active"`
	IsStaff     bool `json:"is_staff"`
	IsSuperuser bool `json:"is_superuser"`
}
