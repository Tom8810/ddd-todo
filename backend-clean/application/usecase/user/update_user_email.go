package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *UserUsecase) UpdateUserEmail(ctx context.Context, input dto.UserUpdateInput) (bool, error) {
	dID, err := vo.NewUserID(input.UserID)
	if err != nil {
		return false, err
	}
	dEmail, err := vo.NewEmail(input.UpdateField)
	if err != nil {
		return false, err
	}
	user, err := s.Repository.UserRepository.FindByID(ctx, dID)
	if err != nil {
		return false, err
	}

	user.UpdateEmail(dEmail)

	err = s.Repository.UserRepository.Save(ctx, user)
	if err != nil {
		return false, err
	}

	return true, nil
}
