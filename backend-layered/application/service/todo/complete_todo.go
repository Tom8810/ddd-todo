package service

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *TodoService) CompleteTodo(ctx context.Context, todoID string) (bool, error) {
	dTodoID, err := vo.NewTodoID(todoID)
	if err != nil {
		return false, err
	}

	dTodo, err := s.Repository.TodoRepository.FindByID(ctx, dTodoID)
	if err != nil {
		return false, err
	}

	dTodo.UpdateStatus(vo.StatusCompleted)

	err = s.Repository.TodoRepository.Save(ctx, dTodo)
	if err != nil {
		return false, err
	}

	return true, nil
}
