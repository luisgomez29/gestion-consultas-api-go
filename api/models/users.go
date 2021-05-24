package models

import (
	"time"
)

// Role tipo personalizado para representar el rol del usuario
type Role string

// UserRole roles del usuario
var UserRole = struct {
	ADMIN  Role
	DOCTOR Role
	USER   Role
}{
	ADMIN:  "ADMIN",
	DOCTOR: "DOC",
	USER:   "USR",
}

// User representa la tabla users en la DB
type User struct {
	ID                   uint      `json:"id"`
	Role                 string    `json:"role"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	IdentificationType   string    `json:"identification_type"`
	IdentificationNumber string    `json:"identification_number"`
	Username             string    `json:"username"`
	Email                string    `json:"email,omitempty"`
	Password             string    `json:"password,omitempty"`
	Phone                string    `json:"phone"`
	Picture              string    `json:"picture"`
	City                 string    `json:"city"`
	Neighborhood         string    `json:"neighborhood"`
	Address              string    `json:"address"`
	IsActive             bool      `json:"is_active"`
	IsStaff              bool      `json:"is_staff,omitempty"`
	IsSuperuser          bool      `json:"is_superuser,omitempty"`
	LastLogin            *string   `json:"last_login"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
