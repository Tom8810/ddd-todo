package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/drepository"
	"github.com/ddd-todo/project-backend/infrastructure/query"
	"github.com/ddd-todo/project-backend/port/dbport"
)

func (r *todoRepositoryImpl) List(ctx context.Context, f drepository.TodoFilter, p drepository.PageInput) ([]*agg.Todo, error) {
	q := query.Use(r.db)
	whereConds := buildWhereConditions(f, q)
	orderConds := buildOrderBy(p, q)

	baseQuery := q.Todo.WithContext(ctx).Where(whereConds...).Order(orderConds)
	baseQuery = applyPaging(baseQuery, p)

	dbTodos, err := baseQuery.Find()

	if err != nil {
		return nil, err
	}

	todos := make([]*agg.Todo, 0, len(dbTodos))
	for _, dbTodo := range dbTodos {
		todo, err := dbport.ToDomainTodo(dbTodo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *todoRepositoryMock) List(ctx context.Context, f drepository.TodoFilter, p drepository.PageInput) ([]*agg.Todo, error) {
	fmt.Println("Mock TodoRepository.List called")
	return []*agg.Todo{}, nil
}
