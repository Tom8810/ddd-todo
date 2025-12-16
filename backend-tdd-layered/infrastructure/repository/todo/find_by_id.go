package repository

import (
	"context"
	"fmt"
	"time"

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

	// Not found case
	if id.Value() == "bb04f183-fb98-6d72-eae2-8633ed7a5c2d" {
		return nil, nil
	}

	// error case
	if id.Value() == "c9f374b7-7c97-2529-de08-f091c9c9921c" {
		return nil, fmt.Errorf("mocked repository error")
	}

	// Valid case
	return &agg.Todo{
		ID:        id,
		Title:     vo.Title("Mocked Todo Title"),
		Status:    vo.StatusWaiting,
		UserID:    vo.UserID("mocked-user-id"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
