package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (r *userRepositoryImpl) Delete(ctx context.Context, id vo.UserID) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id.Value()).Delete(&struct{}{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryMock) Delete(ctx context.Context, id vo.UserID) error {
	fmt.Println("Mock UserRepository.Delete called")
	return nil
}
