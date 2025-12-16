package vo

import "github.com/ddd-todo/project-backend/domain/derr"

type Status string

const (
	StatusWaiting   Status = "WAITING"
	StatusDoing     Status = "DOING"
	StatusCompleted Status = "COMPLETED"
)

var validStatuses = map[Status]bool{
	StatusWaiting:   true,
	StatusDoing:     true,
	StatusCompleted: true,
}

// Constructor
func NewStatus(value string) (Status, error) {
	s := Status(value)
	if !validStatuses[s] {
		return Status(""), derr.ErrInvalidStatus
	}

	return s, nil
}

// Getter
func (s Status) Value() string {
	return string(s)
}
