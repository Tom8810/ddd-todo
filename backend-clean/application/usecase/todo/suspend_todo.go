package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *TodoUsecase) SuspendTodo(ctx context.Context, input dto.TodoMutationInput) (bool, error) {
	dTodoID, err := vo.NewTodoID(input.TodoID)
	if err != nil {
		return false, err
	}

	dTodo, err := s.Repository.TodoRepository.FindByID(ctx, dTodoID)
	if err != nil {
		return false, err
	}
	if dTodo == nil {
		return false, derr.ErrTodoNotFound
	}

	dTodo.UpdateStatus(vo.StatusWaiting)

	err = s.Repository.TodoRepository.Save(ctx, dTodo)
	if err != nil {
		return false, err
	}

	return true, nil
}
