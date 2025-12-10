package service

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *TodoService) DeleteTodo(ctx context.Context, todoID string) (bool, error) {
	dID, err := vo.NewTodoID(todoID)
	if err != nil {
		return false, err
	}

	if err := s.repo.TodoRepository.Delete(ctx, dID); err != nil {
		return false, err
	}

	return true, nil
}
