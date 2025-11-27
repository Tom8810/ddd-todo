package domainport

import (
	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func ToDomainTodoFromCreateTodoInput(input dto.CreateTodoInput, id string) (*agg.Todo, error) {
	dTodoID, err := vo.NewTodoID(id)
	if err != nil {
		return nil, err
	}

	var deadline *vo.Deadline
	if input.Deadline != nil {
		d, err := vo.NewDeadline(*input.Deadline)
		if err != nil {
			return nil, err
		}
		deadline = &d
	}

	var priority *vo.Priority
	if input.Priority != nil {
		p, err := vo.NewPriority(*input.Priority)
		if err != nil {
			return nil, err
		}
		priority = &p
	}

	return agg.NewTodo(
		dTodoID,
		vo.Title(input.Title),
		input.Description,
		vo.StatusWaiting,
		deadline,
		priority,
		vo.UserID(input.UserID),
	), nil
}
