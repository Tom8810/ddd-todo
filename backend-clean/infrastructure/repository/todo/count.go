package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/drepository"
	"github.com/ddd-todo/project-backend/infrastructure/query"
)

func (r *todoRepositoryImpl) Count(ctx context.Context, filter drepository.TodoFilter) (int, error) {
	q := query.Use(r.db)
	whereConds := buildWhereConditions(filter, q)
	baseQuery := q.Todo.WithContext(ctx).Where(whereConds...)

	count, err := baseQuery.Count()

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *todoRepositoryMock) Count(ctx context.Context, filter drepository.TodoFilter) (int, error) {
	fmt.Println("Mock TodoRepository.List called")
	return 0, nil
}
