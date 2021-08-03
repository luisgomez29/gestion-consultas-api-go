package requests

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/luisgomez29/gestion-consultas-api/app/models"
	"github.com/luisgomez29/gestion-consultas-api/app/utils"
)

// Validation of user model fields entered by user.
var (
	firstNameRule = []validation.Rule{
		validation.Required.Error("el nombre es requerido"),
		validation.Length(2, 40).Error("el nombre debe tener entre 2 y 40 caracteres"),
	}

	lastNameRule = []validation.Rule{
		validation.Required.Error("los apellidos son requeridos"),
		validation.Length(2, 40).Error("el nombre debe tener entre 5 y 40 caracteres"),
	}

	identificationTypeRule = []validation.Rule{
		validation.Required.Error("el tipo de identificación es requerido"),
		validation.In(models.IdentificationTypeCC.String(), models.IdentificationTypeCE.String()).Error(
			fmt.Sprintf(
				"el tipo de identificación debe ser %s o %s",
				models.IdentificationTypeCC, models.IdentificationTypeCE,
			),
		),
	}

	identificationNumberRule = []validation.Rule{
		validation.Required.Error("el número de identificación es requerido"),
		is.Digit.Error("su identificación debe ser un valor numérico"),
		validation.Length(6, 10).Error(
			"su identificación debe tener entre 6 y 10 caracteres",
		),
	}

	usernameRule = []validation.Rule{
		validation.Required.Error("el nombre de usuario es requerido"),
		validation.Length(3, 60).Error("su usuario debe tener máximo 60 caracteres"),
		validation.Match(utils.ReUsername).Error("su usuario puede tener letras o números"),
	}

	emailRule = []validation.Rule{
		is.Email.Error("ingrese una dirección de correo electrónico válida"),
	}

	phoneRule = []validation.Rule{
		validation.Required.Error("el teléfono es requerido"),
		validation.Match(utils.ReCellPhone).Error("el teléfono ingresado no es válido"),
	}

	cityRule = []validation.Rule{
		validation.Required.Error("la ciudad es requerida"),
		validation.Match(utils.ReLettersOnly).Error("el nombre de la ciudad debe tener solo letras (A-Z)"),
	}

	neighborhoodRule = []validation.Rule{
		validation.Length(3, 40).Error(
			"el nombre del barrio debe tener entre 3 y 40 caracteres",
		),
	}

	addressRule = []validation.Rule{
		validation.Length(3, 60).Error(
			"la dirección debe tener entre 3 y 60 caracteres",
		),
	}

	passwordRule = []validation.Rule{
		validation.Required.Error("la contraseña es requerida"),
		validation.Length(8, 25).Error(
			"la contraseña debe tener entre 8 y 25 caracteres",
		),
		validation.By(utils.PasswordValidator),
	}

	passwordConfirmationRule = []validation.Rule{
		validation.Required.Error("la contraseña de confirmación es requerida"),
	}
)

// TokenRule verify token
var tokenRule = []validation.Rule{
	validation.Required.Error("el token es requerido"),
}

// Passwords represents the password and the confirmation password.
type Passwords struct {
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"password_confirm,omitempty"`
}

func (p *Passwords) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Password, passwordRule...),
		validation.Field(&p.PasswordConfirm, passwordConfirmationRule...),
	)
}

// ------ REQUESTS

// SignUpRequest represents the user's request for the creation of an account.
type SignUpRequest struct {
	*UserBaseRequest
	Passwords
}

func (rs *SignUpRequest) Validate() error {
	if err := validation.ValidateStruct(rs,
		validation.Field(&rs.FirstName, firstNameRule...),
		validation.Field(&rs.LastName, lastNameRule...),
		validation.Field(&rs.IdentificationType, identificationTypeRule...),
		validation.Field(&rs.IdentificationNumber, identificationNumberRule...),
		validation.Field(&rs.Username, usernameRule...),
		validation.Field(&rs.Email, emailRule...),
		validation.Field(&rs.Phone, phoneRule...),
		validation.Field(&rs.City, cityRule...),
		validation.Field(&rs.Neighborhood, neighborhoodRule...),
		validation.Field(&rs.Address, addressRule...),
	); err != nil {
		return err
	}

	if err := rs.Passwords.Validate(); err != nil {
		return err
	}

	return nil
}

// LoginRequest represents the user's login request.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func (rs *LoginRequest) Validate() error {
	return validation.ValidateStruct(rs,
		validation.Field(&rs.Username, usernameRule...),
		validation.Field(&rs.Password, validation.Required.Error("la contraseña es requerida")),
	)
}

// TokenRequest represents the request to verify or update a JWT token.
type TokenRequest struct {
	Token string `json:"token"`
}

func (rs *TokenRequest) Validate() error {
	return validation.ValidateStruct(rs,
		validation.Field(&rs.Token, tokenRule...),
	)
}

// PasswordResetRequest represents the request to look up the user and send the password reset email.
type PasswordResetRequest struct {
	Username string `json:"username"`
}

func (rs *PasswordResetRequest) Validate() error {
	return validation.ValidateStruct(rs,
		validation.Field(&rs.Username, usernameRule...),
	)
}

// PasswordResetConfirmRequest represents the request to reset the password.
type PasswordResetConfirmRequest struct {
	TokenRequest
	Passwords
}

func (rs *PasswordResetConfirmRequest) Validate() error {
	if err := rs.TokenRequest.Validate(); err != nil {
		return err
	}
	if err := rs.Passwords.Validate(); err != nil {
		return err
	}
	return nil
}
