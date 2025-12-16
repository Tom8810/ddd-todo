package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/domain/drepository"
	"github.com/ddd-todo/project-backend/infrastructure/dbmapper"
	"github.com/ddd-todo/project-backend/infrastructure/query"
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
		todo, err := dbmapper.ToAggTodo(dbTodo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *todoRepositoryMock) List(ctx context.Context, f drepository.TodoFilter, p drepository.PageInput) ([]*agg.Todo, error) {
	fmt.Println("Mock TodoRepository.List called")

	baseTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

	firstTodoID := vo.TodoID("bb04f183-fb98-6d72-eae2-8633ed7a5c2d")
	firstTitle := vo.Title("First Todo")
	firstDescription := "This is the first todo"
	firstStatus := vo.Status("WAITING")
	firstDeadline := vo.Deadline(baseTime.Add(24 * time.Hour))
	firstPriority := vo.Priority("HIGH")
	firstUserID := vo.UserID("c9f374b7-7c97-2529-de08-f091c9c9921c")
	firstCreatedAt := baseTime.Add(-1 * time.Hour)
	firstUpdatedAt := baseTime

	secondTodoID := vo.TodoID("c9f374b7-7c97-2529-de08-f091c9c9921c")
	secondTitle := vo.Title("Second Todo")
	secondDescription := "This is the second todo"
	secondStatus := vo.Status("COMPLETED")
	secondDeadline := vo.Deadline(baseTime.Add(48 * time.Hour))
	secondPriority := vo.Priority("LOW")
	secondUserID := vo.UserID("bb04f183-fb98-6d72-eae2-8633ed7a5c2d")
	secondCreatedAt := baseTime.Add(-2 * time.Hour)
	secondUpdatedAt := baseTime.Add(-30 * time.Minute)

	allItems := []*agg.Todo{
		{
			ID:          firstTodoID,
			Title:       firstTitle,
			Description: &firstDescription,
			Status:      firstStatus,
			Deadline:    &firstDeadline,
			Priority:    &firstPriority,
			UserID:      firstUserID,
			CreatedAt:   firstCreatedAt,
			UpdatedAt:   firstUpdatedAt,
		},
		{
			ID:          secondTodoID,
			Title:       secondTitle,
			Description: &secondDescription,
			Status:      secondStatus,
			Deadline:    &secondDeadline,
			Priority:    &secondPriority,
			UserID:      secondUserID,
			CreatedAt:   secondCreatedAt,
			UpdatedAt:   secondUpdatedAt,
		},
	}
	oneItem := []*agg.Todo{
		{
			ID:          firstTodoID,
			Title:       firstTitle,
			Description: &firstDescription,
			Status:      firstStatus,
			Deadline:    &firstDeadline,
			Priority:    &firstPriority,
			UserID:      firstUserID,
			CreatedAt:   firstCreatedAt,
			UpdatedAt:   firstUpdatedAt,
		},
	}

	if f.Keyword != nil && *f.Keyword == "First" {
		return oneItem, nil
	}

	if f.Keyword != nil && *f.Keyword == "NonExistent" {
		return []*agg.Todo{}, nil
	}

	if f.Keyword != nil && *f.Keyword == "ListError" {
		return nil, fmt.Errorf("simulated list error")
	}

	return allItems, nil
}
