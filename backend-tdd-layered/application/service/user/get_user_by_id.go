package service

import (
	"context"

	"github.com/ddd-todo/project-backend/application/gqlmapper"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*graphmodel.User, error) {
	dID, err := vo.NewUserID(userID)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.UserRepository.FindByID(ctx, dID)
	if err != nil {
		return nil, err
	}

	return gqlmapper.ToGqlUser(user), nil
}
