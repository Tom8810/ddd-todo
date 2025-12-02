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

func NewTodo(id vo.TodoID, title vo.Title, description *string, status vo.Status, priority *vo.Priority, deadline *vo.Deadline, userID vo.UserID) Todo {
	return Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
		Deadline:    deadline,
		UserID:      userID,
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

func (t *Todo) UpdateStatus(newStatus vo.Status) error {
	validTransitions := map[vo.Status][]vo.Status{
		vo.StatusWaiting:   {vo.StatusDoing},
		vo.StatusDoing:     {vo.StatusCompleted},
		vo.StatusCompleted: {},
	}

	allowedStatuses := validTransitions[t.Status]
	for _, allowedStatus := range allowedStatuses {
		if newStatus == allowedStatus {
			t.Status = newStatus
			t.updateTimestamp()
			return nil
		}
	}

	return derr.ErrInvalidStatusTransition
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

func (t *Todo) updateTimestamp() {
	t.UpdatedAt = time.Now()
}
