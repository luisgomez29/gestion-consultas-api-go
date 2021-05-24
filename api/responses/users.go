package responses

// SignUpResponse define los campos para que un usuario se registre
type SignUpResponse struct {
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	IdentificationType   string `json:"identification_type"`
	IdentificationNumber string `json:"identification_number"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	City                 string `json:"city"`
	Neighborhood         string `json:"neighborhood"`
	Address              string `json:"address"`
	Password             string `json:"password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
}

// LoginResponse define los campos para que un usuario inicie sesión
type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}
