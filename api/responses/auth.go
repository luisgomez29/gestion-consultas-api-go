package responses

import (
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
	"github.com/luisgomez29/gestion-consultas-api/api/utils"
)

// Validación de los campos del modelo user que ingresa el usuario
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
		validation.In(models.IdentificationTypeCC, models.IdentificationTypeCE).Error(
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
		validation.By(PasswordValidator),
	}
	passwordConfirmationRule = []validation.Rule{
		validation.Required.Error("la contraseña es requerida"),
	}
)

func PasswordValidator(value interface{}) error {
	s, _ := value.(string)
	if utils.ReDigit.Match([]byte(s)) || utils.ReLetters.Match([]byte(s)) {
		return errors.New("la contraseña debe tener letras y números")
	}
	return nil
}

// SignUpResponse define los campos para que un usuario se registre
type SignUpResponse struct {
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
	Password             string  `json:"password,omitempty"`
	PasswordConfirmation string  `json:"password_confirmation,omitempty"`
}

// Validate valida los campos de SignUpResponse
func (s *SignUpResponse) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.FirstName, firstNameRule...),
		validation.Field(&s.LastName, lastNameRule...),
		validation.Field(&s.IdentificationType, identificationTypeRule...),
		validation.Field(&s.IdentificationNumber, identificationNumberRule...),
		validation.Field(&s.Username, usernameRule...),
		validation.Field(&s.Email, emailRule...),
		validation.Field(&s.Phone, phoneRule...),
		validation.Field(&s.City, cityRule...),
		validation.Field(&s.Neighborhood, neighborhoodRule...),
		validation.Field(&s.Address, addressRule...),
		validation.Field(&s.Password, passwordRule...),
		validation.Field(&s.PasswordConfirmation, passwordConfirmationRule...),
	)
}

// LoginResponse define los campos para que un usuario inicie sesión
type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

// Validate valida los campos de LoginResponse
func (l *LoginResponse) Validate() error {
	return validation.ValidateStruct(l,
		validation.Field(&l.Username, usernameRule...),
		validation.Field(&l.Password, passwordRule...),
	)
}
