package derr

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidTodoID   = errors.New("invalid todo ID")
	ErrInvalidUserID   = errors.New("invalid user ID")
	ErrUserNotFound    = errors.New("user not found")
)
