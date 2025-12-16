package repository

import "gorm.io/gorm"

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

type userRepositoryMock struct{}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}
}
