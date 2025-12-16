package service

import "github.com/ddd-todo/project-backend/infrastructure/repository"

type UserService struct {
	repo *repository.Repositories
}

func NewUserService(repo *repository.Repositories) *UserService {
	return &UserService{
		repo: repo,
	}
}
