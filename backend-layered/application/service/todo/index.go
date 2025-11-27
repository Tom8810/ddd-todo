package service

import "github.com/ddd-todo/project-backend/infrastructure/repository"

type TodoService struct {
	Repository repository.Repositories
}

func NewTodoService(repo *repository.Repositories) *TodoService {
	if repo == nil {
		panic("repository.Repositories is nil")
	}
	return &TodoService{
		Repository: *repo,
	}
}
