package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/dbmapper"
	"github.com/ddd-todo/project-backend/infrastructure/query"
)

func (r *todoRepositoryImpl) FindByID(ctx context.Context, id vo.TodoID) (*agg.Todo, error) {
	q := query.Use(r.db)

	dbTodo, err := q.Todo.WithContext(ctx).Where(q.Todo.ID.Eq(id.Value())).First()
	if err != nil {
		return nil, err
	}

	return dbmapper.ToAggTodo(dbTodo)
}

func (r *todoRepositoryMock) FindByID(ctx context.Context, id vo.TodoID) (*agg.Todo, error) {
	fmt.Println("Mock TodoRepository.FindByID called")
	return &agg.Todo{}, nil
}
