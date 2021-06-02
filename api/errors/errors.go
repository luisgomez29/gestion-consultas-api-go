//Package errors
package errors

// DatabaseValidationError define los métodos para verificar los errores de la base de datos
type DatabaseValidationError interface {

	// ValidatePgError verifica el error de la base de datos.
	// Se debe usar para verificar los errores de unicidad y otros de los campos.
	ValidatePgError(err error) error
}

// ErrNoRows personaliza el mensaje de error cuando ocurre el error de tipo pgx.ErrNoRows.
// Se usa como valor de retorno de error para utils.ValidateErrNoRows
type ErrNoRows struct {
	msg string
}

func (e *ErrNoRows) Error() string {
	return e.msg
}

func NewErrNoRows(msg string) *ErrNoRows {
	return &ErrNoRows{msg}
}
