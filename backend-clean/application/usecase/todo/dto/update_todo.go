package dto

import "time"

type UpdateTodoInput struct {
	ID          string
	Title       *string
	Description *string
	Deadline    *time.Time
	Priority    *string
}
