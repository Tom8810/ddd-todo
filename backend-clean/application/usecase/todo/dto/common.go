package dto

import "time"

type TodoMutationInput struct {
	TodoID string
}

type TodoOutput struct {
	ID          string
	Title       string
	Description *string
	Status      string
	Deadline    *time.Time
	Priority    *string
	UserID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TodoFilterInput struct {
	Keyword        *string
	Status         *string
	Priority       *string
	DeadlineBefore *time.Time
	DeadlineAfter  *time.Time
	UserID         *string
}

type TodoPageInput struct {
	Limit         int
	Offset        int
	SortKey       *string
	SortDirection SortDirectionInput
}

type SortDirectionInput string

const (
	SortDirectionAsc  SortDirectionInput = "ASC"
	SortDirectionDesc SortDirectionInput = "DESC"
)
