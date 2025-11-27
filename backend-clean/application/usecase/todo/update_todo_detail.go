package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *TodoUsecase) UpdateTodoDetail(ctx context.Context, input dto.UpdateTodoInput) (bool, error) {
	dTodoID, err := vo.NewTodoID(input.ID)
	if err != nil {
		return false, err
	}

	dTodo, err := s.Repository.TodoRepository.FindByID(ctx, dTodoID)
	if err != nil {
		return false, err
	}

	if input.Title != nil {
		dTitle, err := vo.NewTitle(*input.Title)
		if err != nil {
			return false, err
		}
		dTodo.UpdateTitle(dTitle)
	}

	if input.Description != nil {
		dTodo.UpdateDescription(input.Description)
	}

	if input.Deadline != nil {
		dDeadline, err := vo.NewDeadline(*input.Deadline)
		if err != nil {
			return false, err
		}
		dTodo.UpdateDeadline(&dDeadline)
	}

	if input.Priority != nil {
		dPriority, err := vo.NewPriority(*input.Priority)
		if err != nil {
			return false, err
		}
		dTodo.UpdatePriority(&dPriority)
	}

	err = s.Repository.TodoRepository.Save(ctx, dTodo)
	if err != nil {
		return false, err
	}

	return true, nil
}
