package service

import "github.com/ddd-todo/project-backend/infrastructure/repository"

type UserService struct {
	Repository repository.Repositories
}

func NewUserService(repo *repository.Repositories) *UserService {
	if repo == nil {
		panic("repository.Repositories is nil")
	}
	return &UserService{
		Repository: *repo,
	}
}
