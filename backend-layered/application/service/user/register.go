package service

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func (s *UserService) Register(ctx context.Context, input graphmodel.UserCreateInput) (bool, error) {
	id := lib.GenerateID()
	dID, err := vo.NewUserID(id)
	if err != nil {
		return false, err
	}

	dName, err := vo.NewName(input.Name)
	if err != nil {
		return false, err
	}

	dEmail, err := vo.NewEmail(input.Email)
	if err != nil {
		return false, err
	}

	hashedPassword, err := lib.HashPassword(input.Password)
	if err != nil {
		return false, err
	}
	dPassword, err := vo.NewPassword(hashedPassword)
	if err != nil {
		return false, err
	}

	dUser := agg.NewUser(dID, dName, dEmail, dPassword)

	err = s.Repository.UserRepository.Save(ctx, dUser)
	if err != nil {
		return false, err
	}

	return true, nil
}
