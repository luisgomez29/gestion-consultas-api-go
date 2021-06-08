package utils

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4"

	apierrors "github.com/luisgomez29/gestion-consultas-api/api/errors"
)

// Regular expressions
var (
	ReLettersOnly = regexp.MustCompile("^[a-zA-ZÁ-ÿ+ ?]*$")
	ReCellPhone   = regexp.MustCompile("^3[0-9]{2} ?[0-9]{3} ?[0-9]{4}$")
	ReUsername    = regexp.MustCompile("^[\\w.@+-]+\\z")
	ReDigit       = regexp.MustCompile("^[0-9]+$")
	ReLetters     = regexp.MustCompile("^[a-zA-z]+$")
)

// Model includes the ID, CreatedAt, and UpdatedAt fields.
// It is used in the definition of models.
type Model struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ValidateErrNoRows checks if the error is of type pgx.ErrNoRows.
func ValidateErrNoRows(err error, msg string) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return apierrors.NewErrNoRows(msg)
	}
	return err
}

// TimeDuration convert a string to time.Duration.
func TimeDuration(t string) (time.Duration, error) {
	tc, err := strconv.Atoi(t)
	if err != nil {
		//log.Printf("Error converting in Integer %v", err)
		return 0, err
	}
	return time.Duration(tc), nil
}
