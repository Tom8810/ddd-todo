package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/model"
)

func (r *todoRepositoryImpl) Delete(ctx context.Context, id vo.TodoID) error {
	if err := r.db.WithContext(ctx).Delete(&model.Todo{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepositoryMock) Delete(ctx context.Context, id vo.TodoID) error {
	fmt.Println("Mock TodoRepository.Delete called")

	// Not found case
	if id.Value() == "bb04f183-fb98-6d72-eae2-8633ed7a5c2d" {
		return fmt.Errorf("todo not found")
	}

	// error case
	if id.Value() == "c9f374b7-7c97-2529-de08-f091c9c9921c" {
		return fmt.Errorf("mocked repository error")
	}

	return nil
}
