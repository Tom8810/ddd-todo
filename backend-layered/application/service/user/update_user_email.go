package service

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *UserService) UpdateUserEmail(ctx context.Context, userID string, email string) (bool, error) {
	dID, err := vo.NewUserID(userID)
	if err != nil {
		return false, err
	}
	dEmail, err := vo.NewEmail(email)
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
