package gqlport

import "github.com/ddd-todo/project-backend/application/usecase/todo/dto"

func ToDtoGetTodoDetailInput(todoID string) dto.GetTodoDetailInput {
	return dto.GetTodoDetailInput{
		TodoID: todoID,
	}
}
