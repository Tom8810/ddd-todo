package repository

import (
	"gorm.io/gorm"
)

type todoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepositoryImpl(db *gorm.DB) *todoRepositoryImpl {
	return &todoRepositoryImpl{db: db}
}

type todoRepositoryMock struct{}

func NewTodoRepositoryMock() *todoRepositoryMock {
	return &todoRepositoryMock{}
}
