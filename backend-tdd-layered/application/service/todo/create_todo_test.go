package service_test

import (
	"context"
	"testing"
	"time"

	service "github.com/ddd-todo/project-backend/application/service/todo"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func TestCreateTodoTest(t *testing.T) {
	repo := repository.NewRepositoriesMock()
	service := service.NewTodoService(repo)
	ctx := context.Background()

	validUserId := lib.GenerateID()
	errCaseId := "c9f374b7-7c97-2529-de08-f091c9c9921c"
	title := "Test Todo"
	description := lib.StringPtr("This is a test todo")
	deadline := lib.TimePtr(time.Now().Add(48 * time.Hour))
	errCaseDeadline := lib.TimePtr(time.Now().Add(-24 * time.Hour))
	priority := lib.StringPtr("HIGH")

	tests := []struct {
		name    string
		input   graphmodel.TodoCreateInput
		wantErr bool
	}{
		{
			name: "正常系: タイトル、説明、期限、優先度、ユーザーIDが有効な場合、Todoが作成されること",
			input: graphmodel.TodoCreateInput{
				Title:       title,
				Description: description,
				Deadline:    deadline,
				Priority:    priority,
				UserID:      validUserId,
			},
			wantErr: false,
		},
		{
			name: "正常系: タイトルとユーザーIDのみが有効な場合、Todoが作成されること",
			input: graphmodel.TodoCreateInput{
				Title:       title,
				Description: nil,
				Deadline:    nil,
				Priority:    nil,
				UserID:      validUserId,
			},
		},
		{
			name: "異常系: タイトルが空の場合、エラーが返されること",
			input: graphmodel.TodoCreateInput{
				Title:       "",
				Description: description,
				Deadline:    deadline,
				Priority:    priority,
				UserID:      validUserId,
			},
			wantErr: true,
		},
		{
			name: "異常系: ユーザーIDが無効な場合、エラーが返されること",
			input: graphmodel.TodoCreateInput{
				Title:       title,
				Description: description,
				Deadline:    deadline,
				Priority:    priority,
				UserID:      "invalid-user-id",
			},
			wantErr: true,
		},
		{
			name: "異常系: 期限が過去の日付の場合、エラーが返されること",
			input: graphmodel.TodoCreateInput{
				Title:       title,
				Description: description,
				Deadline:    errCaseDeadline,
				Priority:    priority,
				UserID:      validUserId,
			},
			wantErr: true,
		},
		{
			name: "異常系: 優先度が無効な場合、エラーが返されること",
			input: graphmodel.TodoCreateInput{
				Title:       title,
				Description: description,
				Deadline:    deadline,
				Priority:    lib.StringPtr("INVALID_PRIORITY"),
				UserID:      validUserId,
			},
			wantErr: true,
		},
		{
			name: "異常系: リポジトリでエラーが発生した場合、エラーが返されること",
			input: graphmodel.TodoCreateInput{
				Title:       title,
				Description: description,
				Deadline:    deadline,
				Priority:    priority,
				UserID:      errCaseId,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.CreateTodo(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
