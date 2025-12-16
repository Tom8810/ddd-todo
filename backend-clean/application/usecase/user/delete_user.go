package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *UserUsecase) DeleteUser(ctx context.Context, input dto.DeleteUserInput) (bool, error) {
	dID, err := vo.NewUserID(input.UserID)
	if err != nil {
		return false, err
	}

	if err := s.Repository.UserRepository.Delete(ctx, dID); err != nil {
		return false, err
	}

	return true, nil
}
