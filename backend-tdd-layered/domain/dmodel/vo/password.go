package vo

import (
	"github.com/ddd-todo/project-backend/domain/derr"
)

type Password string

// Constructor
func NewPassword(value string) (Password, error) {
	if len(value) == 0 || len(value) > 16 {
		return Password(""), derr.ErrInvalidPassword
	}

	return Password(value), nil
}

func NewHashedPassword(value string) Password {
	return Password(value)
}

// Getter
func (p Password) Value() string {
	return string(p)
}
