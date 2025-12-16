package service

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *UserService) UpdateUserProfile(ctx context.Context, userID string, name string) (bool, error) {
	dID, err := vo.NewUserID(userID)
	if err != nil {
		return false, err
	}
	dName, err := vo.NewName(name)
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

	user.UpdateName(dName)

	err = s.repo.UserRepository.Save(ctx, user)
	if err != nil {
		return false, err
	}

	return true, nil
}
