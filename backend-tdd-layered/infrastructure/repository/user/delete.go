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

	notFoundId := "bb04f183-fb98-6d72-eae2-8633ed7a5c2d"
	errCaseId := "c9f374b7-7c97-2529-de08-f091c9c9921c"

	if id.Value() == errCaseId {
		return fmt.Errorf("mocked repository error")
	}
	if id.Value() == notFoundId {
		return fmt.Errorf("user not found")
	}
	return nil
}
