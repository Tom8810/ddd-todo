package gqlport

import (
	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToDtoUpdateTodoInput(m graphmodel.TodoUpdateInput) dto.UpdateTodoInput {
	return dto.UpdateTodoInput{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Deadline:    m.Deadline,
		Priority:    m.Priority,
	}
}
