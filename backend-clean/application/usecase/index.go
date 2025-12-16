package usecase

import (
	"github.com/ddd-todo/project-backend/application/service"
	t "github.com/ddd-todo/project-backend/application/usecase/todo"
	u "github.com/ddd-todo/project-backend/application/usecase/user"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
)

type Usecases struct {
	TodoUsecase *t.TodoUsecase
	UserUsecase *u.UserUsecase
}

func NewUsecases(
	repo *repository.Repositories,
	authService service.AuthService,
) *Usecases {
	return &Usecases{
		TodoUsecase: t.NewTodoUsecase(repo),
		UserUsecase: u.NewUserUsecase(repo, authService),
	}
}
