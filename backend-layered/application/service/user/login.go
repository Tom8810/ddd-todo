package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/ddd-todo/project-backend/application/gqlmapper"
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
	"github.com/ddd-todo/project-backend/infrastructure/auth"
)

func (s *UserService) Login(ctx context.Context, input graphmodel.LoginInput) (*graphmodel.LoginResponse, error) {
	emailVO, err := vo.NewEmail(input.Email)
	if err != nil {
		return &graphmodel.LoginResponse{}, err
	}

	dUser, err := s.Repository.UserRepository.FindByEmail(ctx, emailVO)
	if err != nil {
		return &graphmodel.LoginResponse{}, err
	}

	if !s.verifyPassword(dUser, input.Password) {
		return &graphmodel.LoginResponse{}, err
	}

	token, err := auth.GenerateToken(dUser.ID.Value(), dUser.Email.Value())
	if err != nil {
		return &graphmodel.LoginResponse{}, err
	}

	user := gqlmapper.ToGqlUser(dUser)

	return &graphmodel.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *UserService) verifyPassword(user *agg.User, plainPassword string) bool {
	hashedPassword := user.Password.Value()
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
