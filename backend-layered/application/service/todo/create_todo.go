package service

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func (s *TodoService) CreateTodo(ctx context.Context, input graphmodel.TodoCreateInput) (bool, error) {
	id := lib.GenerateID()
	dTodoID, err := vo.NewTodoID(id)
	if err != nil {
		return false, err
	}

	dTitle, err := vo.NewTitle(input.Title)
	if err != nil {
		return false, err
	}

	var dDeadline vo.Deadline
	if input.Deadline != nil {
		dDeadline, err = vo.NewDeadline(*input.Deadline)
		if err != nil {
			return false, err
		}
	}

	var dPriority vo.Priority
	if input.Priority != nil {
		dPriority, err = vo.NewPriority(*input.Priority)
		if err != nil {
			return false, err
		}
	}

	dUserID, err := vo.NewUserID(input.UserID)
	if err != nil {
		return false, err
	}

	newTodo := agg.NewTodo(
		dTodoID,
		dTitle,
		input.Description,
		vo.StatusWaiting,
		&dDeadline,
		&dPriority,
		dUserID,
	)

	err = s.Repository.TodoRepository.Save(ctx, newTodo)
	if err != nil {
		return false, err
	}

	return true, nil
}
