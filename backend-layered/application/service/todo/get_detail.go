package service

import (
	"context"

	"github.com/ddd-todo/project-backend/application/gqlmapper"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func (s *TodoService) GetTodoDetail(ctx context.Context, todoID string) (*graphmodel.Todo, error) {
	dID, err := vo.NewTodoID(todoID)
	if err != nil {
		return nil, err
	}

	todo, err := s.Repository.TodoRepository.FindByID(ctx, dID)
	if err != nil {
		return nil, err
	}

	return gqlmapper.ToGqlTodo(todo), nil
}
