package utils

import (
	"regexp"
)

// Expresiones regulares
var (
	ReLettersOnly = regexp.MustCompile("^[a-zA-ZÁ-ÿ+ ?]*$")
	ReCellPhone   = regexp.MustCompile("^3[0-9]{2} ?[0-9]{3} ?[0-9]{4}$")
	ReUsername    = regexp.MustCompile("^[\\w.@+-]+\\z")
	ReDigit       = regexp.MustCompile("^[0-9]+$")
	ReLetters     = regexp.MustCompile("^[a-zA-z]+$")
)

// DatabaseValidationError define los métodos para verificar los errores de la base de datos
type DatabaseValidationError interface {

	// ValidatePgError verifica el error de la base de datos.
	// Se debe usar para verificar los errores de unicidad y otros de los campos.
	ValidatePgError(err error) error
}
