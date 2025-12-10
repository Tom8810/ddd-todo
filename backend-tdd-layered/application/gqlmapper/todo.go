package gqlmapper

import (
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/domain/drepository"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToGqlTodo(a *agg.Todo) *graphmodel.Todo {
	if a == nil {
		return nil
	}
	return &graphmodel.Todo{
		ID:          a.GetID(),
		Title:       a.GetTitle(),
		Description: a.GetDescription(),
		Status:      a.GetStatus(),
		Deadline:    a.GetDeadline(),
		Priority:    a.GetPriority(),
		UserID:      a.GetUserID(),
		CreatedAt:   a.GetCreatedAt(),
		UpdatedAt:   a.GetUpdatedAt(),
	}
}

func ToGqlTodoConnection(todos []*agg.Todo, pagingInfo graphmodel.PagingInfo) *graphmodel.TodoConnection {
	items := make([]*graphmodel.Todo, len(todos))
	for i, todo := range todos {
		items[i] = ToGqlTodo(todo)
	}

	return &graphmodel.TodoConnection{
		Items:      items,
		PagingInfo: &pagingInfo,
	}
}

func ToAggTodoFilter(m *graphmodel.TodoFilter) (drepository.TodoFilter, error) {
	var err error
	var dUserID vo.UserID
	if m.UserID != nil {
		dUserID, err = vo.NewUserID(*m.UserID)
		if err != nil {
			return drepository.TodoFilter{}, err
		}
	}

	var status vo.Status
	if m.Status != nil {
		status, err = vo.NewStatus(*m.Status)
		if err != nil {
			return drepository.TodoFilter{}, err
		}
	}

	var priority vo.Priority
	if m.Priority != nil {
		priority, err = vo.NewPriority(*m.Priority)
		if err != nil {
			return drepository.TodoFilter{}, err
		}
	}

	return drepository.NewTodoFilter(
		m.Keyword,
		&dUserID,
		&status,
		&priority,
		m.DeadlineAfter,
		m.DeadlineBefore,
	)
}

func ToAggTodo(m *graphmodel.Todo) (*agg.Todo, error) {
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

	todo := agg.NewTodo(
		vo.TodoID(m.ID),
		vo.Title(m.Title),
		m.Description,
		vo.Status(m.Status),
		priority,
		deadline,
		vo.UserID(m.UserID),
	)

	return &todo, nil
}
