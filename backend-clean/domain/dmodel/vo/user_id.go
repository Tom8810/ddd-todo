package vo

import (
	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/google/uuid"
)

type UserID string

// Constructor
func NewUserID(value string) (UserID, error) {
	if _, err := uuid.Parse(value); err != nil {
		return UserID(""), derr.ErrInvalidUserID
	}
	return UserID(value), nil
}

// Getter
func (id UserID) Value() string {
	return string(id)
}
