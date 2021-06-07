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

// SignUpResponse representa la solicitud del usuario para la creación de una cuenta.
type SignUpResponse struct {
	*UserBaseResponse

	Password             string `json:"password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
}

// Validate valida los campos de SignUpResponse
func (rs *SignUpResponse) Validate() error {
	return validation.ValidateStruct(rs,
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
		validation.Field(&rs.Password, passwordRule...),
		validation.Field(&rs.PasswordConfirmation, passwordConfirmationRule...),
	)
}

// LoginResponse representa la solicitud para el inicio de sesión del usuario.
type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

// Validate valida los campos de LoginResponse
func (rs *LoginResponse) Validate() error {
	return validation.ValidateStruct(rs,
		validation.Field(&rs.Username, usernameRule...),
		validation.Field(&rs.Password, validation.Required.Error("la contraseña es requerida")),
	)
}

// TokenResponse representa la solicitud para verificar o actualizar un token JWT.
type TokenResponse struct {
	Token string `json:"token"`
}

// Validate valida los campos de TokenResponse.
func (rs *TokenResponse) Validate() error {
	return validation.ValidateStruct(rs,
		validation.Field(&rs.Token, validation.Required.Error("el token es requerido")),
	)
}
