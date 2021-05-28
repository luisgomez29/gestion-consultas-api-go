// Package errors implementa los tipos y funciones para manipular los errores en la API
package errors

// DatabaseValidationError define los métodos para verificar los errores de la base de datos
type DatabaseValidationError interface {

	// ValidatePgError verifica el error de la base de datos.
	// Se debe usar para verificar los errores de unicidad y otros de los campos.
	ValidatePgError(err error) error
}
