package vo

import (
	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/google/uuid"
)

type TodoID string

// Constructor
func NewTodoID(value string) (TodoID, error) {
	if _, err := uuid.Parse(value); err != nil {
		return TodoID(""), derr.ErrInvalidTodoID
	}
	return TodoID(value), nil
}

// Getter
func (id TodoID) Value() string {
	return string(id)
}
