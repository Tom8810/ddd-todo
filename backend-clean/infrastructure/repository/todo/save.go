package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/port/dbport"
)

func (r *todoRepositoryImpl) Save(ctx context.Context, todo *agg.Todo) error {
	dbTodo := dbport.ToDBTodo(todo)

	if err := r.db.WithContext(ctx).Save(dbTodo).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepositoryMock) Save(ctx context.Context, todo *agg.Todo) error {
	fmt.Println("Mock TodoRepository.Save called")
	return nil
}
