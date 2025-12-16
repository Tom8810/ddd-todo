package vo

import "github.com/ddd-todo/project-backend/domain/derr"

type Status string

const (
	StatusWaiting   Status = "WAITING"
	StatusDoing     Status = "DOING"
	StatusCompleted Status = "COMPLETED"
)

func NewStatus(s string) (Status, error) {
	switch Status(s) {
	case StatusWaiting, StatusDoing, StatusCompleted:
		return Status(s), nil
	default:
		return "", derr.ErrInvalidStatus
	}
}

func (s Status) Value() string {
	return string(s)
}
