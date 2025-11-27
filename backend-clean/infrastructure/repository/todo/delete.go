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
	return nil
}
