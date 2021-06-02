package utils

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4"

	api "github.com/luisgomez29/gestion-consultas-api/api/errors"
)

// Expresiones regulares
var (
	ReLettersOnly = regexp.MustCompile("^[a-zA-ZÁ-ÿ+ ?]*$")
	ReCellPhone   = regexp.MustCompile("^3[0-9]{2} ?[0-9]{3} ?[0-9]{4}$")
	ReUsername    = regexp.MustCompile("^[\\w.@+-]+\\z")
	ReDigit       = regexp.MustCompile("^[0-9]+$")
	ReLetters     = regexp.MustCompile("^[a-zA-z]+$")
)

// Model incluye los campos ID, CreatedAt, UpdatedAt.
// Es usada en la definición de los modelos.
type Model struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ValidateErrNoRows verifica si el error es de tipo pgx.ErrNoRows
func ValidateErrNoRows(err error, msg string) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return api.NewErrNoRows(msg)
	}
	return err
}

//Hours convierte el número de dias a horas
func Hours(days string) time.Duration {
	d, err := strconv.Atoi(days)
	if err != nil {
		log.Printf("Error converting in Integer %v", err)
	}
	return time.Duration(24 * d)
}
