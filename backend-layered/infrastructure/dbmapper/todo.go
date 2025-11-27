package dbmapper

import (
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/model"
)

func ToDBTodo(a *agg.Todo) *model.Todo {
	return &model.Todo{
		ID:          a.GetID(),
		Title:       a.GetTitle(),
		Description: a.GetDescription(),
		Status:      a.GetStatus(),
		Deadline:    a.GetDeadline(),
		Priority:    a.GetPriority(),
		UserId:      a.GetUserID(),
	}
}

func ToAggTodo(m *model.Todo) (*agg.Todo, error) {
	var deadline *vo.Deadline
	if m.Deadline != nil {
		d, err := vo.NewDeadline(*m.Deadline)
		if err != nil {
			return nil, err
		}
		deadline = &d
	}

	var priority *vo.Priority
	if m.Priority != nil {
		p, err := vo.NewPriority(*m.Priority)
		if err != nil {
			return nil, err
		}
		priority = &p
	}

	return agg.ReconstructTodo(
		vo.TodoID(m.ID),
		vo.Title(m.Title),
		m.Description,
		vo.Status(m.Status),
		deadline,
		priority,
		vo.UserID(m.UserId),
		m.CreatedAt,
		m.UpdatedAt,
	), nil
}
