package vo

import (
	"time"

	"github.com/ddd-todo/project-backend/domain/derr"
)

type Deadline time.Time

// Constructor
func NewDeadline(value time.Time) (Deadline, error) {
	now := time.Now()
	if value.Before(now) {
		return Deadline{}, derr.ErrInvalidDeadline
	}
	return Deadline(value), nil
}

// Getter
func (d Deadline) Value() time.Time {
	return time.Time(d)
}
