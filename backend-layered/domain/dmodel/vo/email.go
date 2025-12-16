package vo

import (
	"regexp"

	"github.com/ddd-todo/project-backend/domain/derr"
)

type Email string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Constructor
func NewEmail(value string) (Email, error) {
	if !emailRegex.MatchString(value) {
		return Email(""), derr.ErrInvalidEmail
	}

	return Email(value), nil
}

// Getter
func (e Email) Value() string {
	return string(e)
}
