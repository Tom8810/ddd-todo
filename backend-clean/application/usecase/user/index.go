package usecase

import (
	"github.com/ddd-todo/project-backend/application/service"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
)

type UserUsecase struct {
	Repository  repository.Repositories
	AuthService service.AuthService
}

func NewUserUsecase(repo *repository.Repositories, authService service.AuthService) *UserUsecase {
	if repo == nil {
		panic("repository.Repositories is nil")
	}
	return &UserUsecase{
		Repository:  *repo,
		AuthService: authService,
	}
}
