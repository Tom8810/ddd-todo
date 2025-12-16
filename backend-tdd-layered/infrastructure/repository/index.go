package repository

import (
	"github.com/ddd-todo/project-backend/domain/drepository"
	t "github.com/ddd-todo/project-backend/infrastructure/repository/todo"
	u "github.com/ddd-todo/project-backend/infrastructure/repository/user"

	"gorm.io/gorm"
)

type Repositories struct {
	TodoRepository drepository.TodoRepository
	UserRepository drepository.UserRepository
	DB             *gorm.DB
}

func NewRepositoriesImpl(db *gorm.DB) *Repositories {
	todoRepo := t.NewTodoRepositoryImpl(db)
	userRepo := u.NewUserRepositoryImpl(db)
	return &Repositories{
		TodoRepository: todoRepo,
		UserRepository: userRepo,
		DB:             db,
	}
}

func NewRepositoriesMock() *Repositories {
	return &Repositories{
		TodoRepository: t.NewTodoRepositoryMock(),
		UserRepository: u.NewUserRepositoryMock(),
	}
}
