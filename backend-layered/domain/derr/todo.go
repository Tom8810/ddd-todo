package derr

import "errors"

var (
	ErrInvalidTitle               = errors.New("invalid title")
	ErrInvalidStatus              = errors.New("invalid status")
	ErrInvalidDeadline            = errors.New("invalid deadline")
	ErrInvalidPriority            = errors.New("invalid priority")
	ErrInvalidName                = errors.New("invalid name")
	ErrInvalidEmail               = errors.New("invalid email")
	ErrInvalidStatusTransition    = errors.New("invalid status transition")
	ErrInvalidTodoFitlerDateRange = errors.New("invalid todo filter date range")
	ErrInvalidTodoSortField       = errors.New("invalid todo sort field")
)
