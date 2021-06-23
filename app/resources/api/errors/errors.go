//Package errors contains the types and functions related to errors
package errors

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// DatabaseValidationError defines methods to check for database errors.
type DatabaseValidationError interface {

	// ValidatePgError check the database error.
	// It should be used to check for uniqueness and other errors in the fields.
	ValidatePgError(err error) error
}

// PasswordMismatch occurs when passwords do not match
var PasswordMismatch = validation.Errors{"password": fmt.Errorf("las contraseñas ingresadas no coinciden")}

// ErrNoRows Customize the error message when the error of type pgx.ErrNoRows occurs.
// Used as an error return value for utils.ValidateErrNoRows.
type ErrNoRows struct {
	msg string
}

func (e *ErrNoRows) Error() string {
	return e.msg
}

func NewErrNoRows(msg string) *ErrNoRows {
	return &ErrNoRows{msg}
}
