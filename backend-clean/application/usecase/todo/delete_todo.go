package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *TodoUsecase) DeleteTodo(ctx context.Context, input dto.TodoMutationInput) (bool, error) {
	dID, err := vo.NewTodoID(input.TodoID)
	if err != nil {
		return false, err
	}

	if err := s.Repository.TodoRepository.Delete(ctx, dID); err != nil {
		return false, err
	}

	return true, nil
}
