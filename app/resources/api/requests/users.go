package requests

// UserBaseRequest contains the user fields that are returned regardless of role.
type UserBaseRequest struct {
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

// UserDefaultRequest represents the request to create a user with a models.UserDefault or models.UserDoctor role.
type UserDefaultRequest struct {
	*UserBaseRequest

	Role string `json:"role"`
}

// UserAdminRequest represents the request to create a user with the models.UserAdmin role.
type UserAdminRequest struct {
	*UserDefaultRequest

	IsActive    bool `json:"is_active"`
	IsStaff     bool `json:"is_staff"`
	IsSuperuser bool `json:"is_superuser"`
}
