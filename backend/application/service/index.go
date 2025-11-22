package service

import (
	t "github.com/ddd-todo/project-backend/application/service/todo"
	u "github.com/ddd-todo/project-backend/application/service/user"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
)

type Services struct {
	TodoService *t.TodoService
	UserService *u.UserService
}

func NewServices(
	repo *repository.Repositories,
) *Services {
	return &Services{
		TodoService: t.NewTodoService(repo),
		UserService: u.NewUserService(repo),
	}
}
