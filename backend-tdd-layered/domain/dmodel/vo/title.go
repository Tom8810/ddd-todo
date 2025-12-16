package vo

import "github.com/ddd-todo/project-backend/domain/derr"

type Title string

func NewTitle(value string) (Title, error) {
	if len(value) == 0 || len(value) > 64 {
		return Title(""), derr.ErrInvalidTitle
	}
	return Title(value), nil
}

func (t Title) Value() string {
	return string(t)
}
