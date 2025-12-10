package service

import (
	"context"

	"github.com/ddd-todo/project-backend/application/gqlmapper"
	application_utils "github.com/ddd-todo/project-backend/application/utils"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func (s *TodoService) ListTodos(ctx context.Context, filter graphmodel.TodoFilter, pageInput graphmodel.PageInput) (*graphmodel.TodoConnection, error) {
	dFilter, err := gqlmapper.ToAggTodoFilter(&filter)
	if err != nil {
		return nil, err
	}
	dPaging, err := gqlmapper.ToAggPageInput(&pageInput)
	if err != nil {
		return nil, err
	}

	todos, err := s.repo.TodoRepository.List(ctx, dFilter, dPaging)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.TodoRepository.Count(ctx, dFilter)
	if err != nil {
		return nil, err
	}

	pagingInfo := application_utils.GetPagingInfo(count, pageInput)

	return gqlmapper.ToGqlTodoConnection(todos, *pagingInfo), nil
}
