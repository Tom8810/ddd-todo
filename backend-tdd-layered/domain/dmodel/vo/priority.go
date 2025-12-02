package vo

import "github.com/ddd-todo/project-backend/domain/derr"

type Priority string

const (
	PriorityLow    Priority = "LOW"
	PriorityMedium Priority = "MEDIUM"
	PriorityHigh   Priority = "HIGH"
)

var validPriorities = map[Priority]bool{
	PriorityLow:    true,
	PriorityMedium: true,
	PriorityHigh:   true,
}

// Constructor
func NewPriority(value string) (Priority, error) {
	p := Priority(value)
	if !validPriorities[p] {
		return Priority(""), derr.ErrInvalidPriority
	}

	return p, nil
}

// Getter
func (p Priority) Value() string {
	return string(p)
}
