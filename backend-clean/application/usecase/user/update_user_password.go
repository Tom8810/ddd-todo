package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func (s *UserUsecase) UpdateUserPassword(ctx context.Context, input dto.UserUpdateInput) (bool, error) {
	dID, err := vo.NewUserID(input.UserID)
	if err != nil {
		return false, err
	}
	dPassword, err := vo.NewPassword(input.UpdateField)
	if err != nil {
		return false, err
	}
	hashedDPassword, err := lib.HashPassword(dPassword)
	if err != nil {
		return false, err
	}
	user, err := s.Repository.UserRepository.FindByID(ctx, dID)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, derr.ErrUserNotFound
	}

	user.UpdatePassword(hashedDPassword)

	err = s.Repository.UserRepository.Save(ctx, user)
	if err != nil {
		return false, err
	}

	return true, nil
}
