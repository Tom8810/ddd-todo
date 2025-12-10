package service

import "github.com/ddd-todo/project-backend/infrastructure/repository"

type TodoService struct {
	repo *repository.Repositories
}

func NewTodoService(repo *repository.Repositories) *TodoService {
	return &TodoService{
		repo: repo,
	}
}
