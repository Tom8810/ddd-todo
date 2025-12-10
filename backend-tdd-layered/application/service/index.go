package service

import (
	t "github.com/ddd-todo/project-backend/application/service/todo"
	u "github.com/ddd-todo/project-backend/application/service/user"
)

type Services struct {
	TodoService *t.TodoService
	UserService *u.UserService
}

func NewServices(userService *u.UserService, todoService *t.TodoService) *Services {
	return &Services{
		TodoService: todoService,
		UserService: userService,
	}
}
