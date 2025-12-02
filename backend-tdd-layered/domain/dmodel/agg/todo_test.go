package agg

import (
	"testing"
	"time"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func TestUpdateStatus(t *testing.T) {
	timeNow := time.Now()
	id, _ := vo.NewTodoID(lib.GenerateID())
	title, _ := vo.NewTitle("Test Todo")
	userID := vo.UserID(lib.GenerateID())

	testTodoWaiting := *ReconstructTodo(
		id,
		title,
		nil,
		vo.StatusWaiting,
		nil,
		nil,
		userID,
		timeNow,
		timeNow,
	)
	testTodoDoing := *ReconstructTodo(
		id,
		title,
		nil,
		vo.StatusDoing,
		nil,
		nil,
		userID,
		timeNow,
		timeNow,
	)
	testTodoCompleted := *ReconstructTodo(
		id,
		title,
		nil,
		vo.StatusCompleted,
		nil,
		nil,
		userID,
		timeNow,
		timeNow,
	)

	tests := []struct {
		name       string
		beforeTodo Todo
		input      vo.Status
		wantErr    bool
		wantTodo   Todo
	}{
		{
			name:       "Valid status - waiting to doing",
			beforeTodo: testTodoWaiting,
			input:      vo.StatusDoing,
			wantErr:    false,
			wantTodo:   testTodoDoing,
		},
		{
			name:       "Valid status - doing to completed",
			beforeTodo: testTodoDoing,
			input:      vo.StatusCompleted,
			wantErr:    false,
			wantTodo:   testTodoCompleted,
		},
		{
			name:       "Invalid status - completed to waiting",
			beforeTodo: testTodoCompleted,
			input:      vo.StatusWaiting,
			wantErr:    true,
		},
		{
			name:       "Invalid status - waiting to completed",
			beforeTodo: testTodoWaiting,
			input:      vo.StatusCompleted,
			wantErr:    true,
		},
		{
			name:       "Invalid status - empty",
			beforeTodo: testTodoWaiting,
			input:      vo.Status(""),
			wantErr:    true,
		},
		{
			name:       "Invalid status - unknown",
			beforeTodo: testTodoWaiting,
			input:      vo.Status("UNKNOWN"),
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.beforeTodo.UpdateStatus(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.beforeTodo.Status != tt.wantTodo.Status {
				t.Errorf("ChangeStatus() status = %v, want %v", tt.beforeTodo.Status, tt.wantTodo.Status)
			}
			if !tt.wantErr && !tt.beforeTodo.UpdatedAt.After(timeNow) {
				t.Errorf("ChangeStatus() UpdatedAt is not updated properly")
			}
		})
	}
}

func TestNewTodo(t *testing.T) {
	id, _ := vo.NewTodoID(lib.GenerateID())
	title, _ := vo.NewTitle("Test Todo")
	userID, _ := vo.NewUserID(lib.GenerateID())

	description := "This is a test todo"
	priority, _ := vo.NewPriority("HIGH")
	deadline, _ := vo.NewDeadline(time.Now().Add(24 * time.Hour))

	tests := []struct {
		name  string
		input struct {
			id          vo.TodoID
			title       vo.Title
			description *string
			status      vo.Status
			priority    *vo.Priority
			deadline    *vo.Deadline
			userID      vo.UserID
		}
		want Todo
	}{
		{
			name: "Normal case",
			input: struct {
				id          vo.TodoID
				title       vo.Title
				description *string
				status      vo.Status
				priority    *vo.Priority
				deadline    *vo.Deadline
				userID      vo.UserID
			}{
				id:          id,
				title:       title,
				description: nil,
				status:      vo.StatusWaiting,
				priority:    nil,
				deadline:    nil,
				userID:      userID,
			},
			want: Todo{
				ID:          id,
				Title:       title,
				Description: nil,
				Status:      vo.StatusWaiting,
				Priority:    nil,
				Deadline:    nil,
				UserID:      userID,
			},
		},
		{
			name: "With description, priority and deadline",
			input: struct {
				id          vo.TodoID
				title       vo.Title
				description *string
				status      vo.Status
				priority    *vo.Priority
				deadline    *vo.Deadline
				userID      vo.UserID
			}{
				id:          id,
				title:       title,
				description: &description,
				status:      vo.StatusDoing,
				priority:    &priority,
				deadline:    &deadline,
				userID:      userID,
			},
			want: Todo{
				ID:          id,
				Title:       title,
				Description: &description,
				Status:      vo.StatusDoing,
				Priority:    &priority,
				Deadline:    &deadline,
				UserID:      userID,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := NewTodo(
				tt.input.id,
				tt.input.title,
				tt.input.description,
				tt.input.status,
				tt.input.priority,
				tt.input.deadline,
				tt.input.userID,
			)
			if got.ID != tt.want.ID ||
				got.Title != tt.want.Title ||
				got.Description != tt.want.Description ||
				got.Status != tt.want.Status ||
				got.Priority != tt.want.Priority ||
				got.Deadline != tt.want.Deadline ||
				got.UserID != tt.want.UserID {
				t.Errorf("NewTodo() = %v, want %v", got, tt.want)
			}
		},
		)
	}
}
