package service

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func (s *UserService) UpdateUserPassword(ctx context.Context, userID string, password string) (bool, error) {
	dID, err := vo.NewUserID(userID)
	if err != nil {
		return false, err
	}
	dPassword, err := vo.NewPassword(password)
	if err != nil {
		return false, err
	}
	hashedDPassword, err := lib.HashPassword(dPassword)
	if err != nil {
		return false, err
	}
	user, err := s.repo.UserRepository.FindByID(ctx, dID)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, derr.ErrUserNotFound
	}

	user.UpdatePassword(hashedDPassword)

	err = s.repo.UserRepository.Save(ctx, user)
	if err != nil {
		return false, err
	}

	return true, nil
}
