package responses

import (
	"time"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

// UserResponse lists the fields to return for the user types.
func UserResponse(u *models.User) *models.User {
	if u.Role != models.UserAdmin.String() {
		u.IsActive = false
	}
	u.Password = ""
	return u
}

// UserBaseResponse contains the user fields that are returned regardless of role.
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
