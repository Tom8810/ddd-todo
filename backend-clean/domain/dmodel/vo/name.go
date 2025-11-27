package vo

import "github.com/ddd-todo/project-backend/domain/derr"

type Name string

// Constructor
func NewName(value string) (Name, error) {
	if len(value) == 0 {
		value = "ゲスト"
	}
	if len(value) > 64 {
		return Name(""), derr.ErrInvalidName
	}
	return Name(value), nil
}

// Getter
func (n Name) Value() string {
	return string(n)
}
