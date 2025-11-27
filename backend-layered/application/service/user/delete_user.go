package service

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func (s *UserService) DeleteUser(ctx context.Context, userID string) (bool, error) {
	dID, err := vo.NewUserID(userID)
	if err != nil {
		return false, err
	}

	if err := s.Repository.UserRepository.Delete(ctx, dID); err != nil {
		return false, err
	}

	return true, nil
}
