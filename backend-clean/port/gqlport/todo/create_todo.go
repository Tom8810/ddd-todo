package gqlport

import (
	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToDtoCreateTodoInput(m graphmodel.TodoCreateInput) dto.CreateTodoInput {
	return dto.CreateTodoInput{
		Title:       m.Title,
		Description: m.Description,
		Deadline:    m.Deadline,
		Priority:    m.Priority,
		UserID:      m.UserID,
	}
}
