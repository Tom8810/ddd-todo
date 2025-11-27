package agg

import (
	"time"

	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

type Todo struct {
	ID          vo.TodoID
	Title       vo.Title
	Description *string
	Status      vo.Status
	Deadline    *vo.Deadline
	Priority    *vo.Priority
	UserID      vo.UserID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodo(
	id vo.TodoID,
	title vo.Title,
	description *string,
	status vo.Status,
	deadline *vo.Deadline,
	priority *vo.Priority,
	userID vo.UserID,
) *Todo {
	now := time.Now()

	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		Deadline:    deadline,
		Priority:    priority,
		UserID:      userID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func ReconstructTodo(
	id vo.TodoID,
	title vo.Title,
	description *string,
	status vo.Status,
	deadline *vo.Deadline,
	priority *vo.Priority,
	userID vo.UserID,
	createdAt time.Time,
	updatedAt time.Time,
) *Todo {
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		Deadline:    deadline,
		Priority:    priority,
		UserID:      userID,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func (t *Todo) GetID() string {
	return t.ID.Value()
}

func (t *Todo) GetTitle() string {
	return t.Title.Value()
}

func (t *Todo) GetDescription() *string {
	return t.Description
}

func (t *Todo) GetStatus() string {
	return t.Status.Value()
}

func (t *Todo) GetDeadline() *time.Time {
	if t.Deadline == nil {
		return nil
	}
	dl := t.Deadline.Value()
	return &dl
}

func (t *Todo) GetPriority() *string {
	if t.Priority == nil {
		return nil
	}
	pr := t.Priority.Value()
	return &pr
}

func (t *Todo) GetUserID() string {
	return t.UserID.Value()
}

func (t *Todo) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *Todo) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}

func (t *Todo) UpdateTitle(title vo.Title) {
	t.Title = title
	t.updateTimestamp()
}

func (t *Todo) UpdateDescription(description *string) {
	t.Description = description
	t.updateTimestamp()
}

func (t *Todo) UpdateDeadline(deadline *vo.Deadline) {
	t.Deadline = deadline
	t.updateTimestamp()
}

func (t *Todo) UpdatePriority(priority *vo.Priority) {
	t.Priority = priority
	t.updateTimestamp()
}

func (t *Todo) UpdateStatus(status vo.Status) error {
	currentStatus := t.Status.Value()
	if currentStatus == string(vo.StatusWaiting) && status.Value() == string(vo.StatusCompleted) ||
		currentStatus == string(vo.StatusCompleted) && status.Value() == string(vo.StatusWaiting) {
		return derr.ErrInvalidStatusTransition
	}
	t.Status = status
	t.updateTimestamp()
	return nil
}

func (t *Todo) updateTimestamp() {
	t.UpdatedAt = time.Now()
}
