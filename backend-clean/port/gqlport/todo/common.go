package gqlport

import (
	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToDtoTodoMutationInput(todoID string) dto.TodoMutationInput {
	return dto.TodoMutationInput{
		TodoID: todoID,
	}
}

func ToDtoTodoFilterInput(filter graphmodel.TodoFilter) dto.TodoFilterInput {
	return dto.TodoFilterInput{
		Keyword:        filter.Keyword,
		Status:         filter.Status,
		Priority:       filter.Priority,
		DeadlineBefore: filter.DeadlineBefore,
		DeadlineAfter:  filter.DeadlineAfter,
		UserID:         filter.UserID,
	}
}

func ToDtoTodoPageInput(page graphmodel.PageInput) dto.TodoPageInput {
	return dto.TodoPageInput{
		Limit:         page.Limit,
		Offset:        page.Offset,
		SortKey:       page.SortKey,
		SortDirection: dto.SortDirectionInput(page.SortDirection),
	}
}

func ToGqlTodo(todo dto.TodoOutput) *graphmodel.Todo {
	return &graphmodel.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		Deadline:    todo.Deadline,
		Priority:    todo.Priority,
		UserID:      todo.UserID,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
