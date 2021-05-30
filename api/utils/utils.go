package utils

import (
	"log"
	"regexp"
	"strconv"
	"time"
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

//Hours convierte el número de dias a horas
func Hours(days string) time.Duration {
	d, err := strconv.Atoi(days)
	if err != nil {
		log.Printf("error al convertir configuraciones a tipo int %v", err)
	}
	return time.Duration(24 * d)
}
