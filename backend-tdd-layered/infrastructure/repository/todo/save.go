package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/infrastructure/dbmapper"
)

func (r *todoRepositoryImpl) Save(ctx context.Context, todo *agg.Todo) error {
	dbTodo := dbmapper.ToDBTodo(todo)

	if err := r.db.WithContext(ctx).Save(dbTodo).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepositoryMock) Save(ctx context.Context, todo *agg.Todo) error {
	if todo.UserID.Value() == "c9f374b7-7c97-2529-de08-f091c9c9921c" {
		return fmt.Errorf("mocked error in Save")
	}

	if todo.ID.Value() == "d4e5f6a7-b8c9-0d1e-2f3a-4b5c6d7e8f90" {
		return fmt.Errorf("mocked error in Save")
	}

	return nil
}
