package dto

import "time"

type CreateTodoInput struct {
	Title       string
	Description *string
	Deadline    *time.Time
	Priority    *string
	UserID      string
}
