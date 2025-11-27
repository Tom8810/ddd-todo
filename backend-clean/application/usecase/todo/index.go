package usecase

import "github.com/ddd-todo/project-backend/infrastructure/repository"

type TodoUsecase struct {
	Repository repository.Repositories
}

func NewTodoUsecase(repo *repository.Repositories) *TodoUsecase {
	if repo == nil {
		panic("repository.Repositories is nil")
	}
	return &TodoUsecase{
		Repository: *repo,
	}
}
