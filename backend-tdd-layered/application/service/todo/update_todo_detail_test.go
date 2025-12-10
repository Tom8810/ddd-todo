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

func TestUpdateTodoDetail(t *testing.T) {
	repo := repository.NewRepositoriesMock()
	service := service.NewTodoService(repo)
	context := context.Background()

	validId := lib.GenerateID()
	notFoundId := "bb04f183-fb98-6d72-eae2-8633ed7a5c2d"
	errCaseInFindId := "c9f374b7-7c97-2529-de08-f091c9c9921c"
	errCaseInSaveId := "d4e5f6a7-b8c9-0d1e-2f3a-4b5c6d7e8f90"

	title := lib.StringPtr("Test Todo")
	description := lib.StringPtr("This is a test todo")
	deadline := lib.TimePtr(time.Now().Add(48 * time.Hour))
	priority := lib.StringPtr("HIGH")

	tests := []struct {
		name    string
		input   graphmodel.TodoUpdateInput
		wantErr bool
	}{
		{
			name: "正常系: タイトル、説明、期限、優先度が有効な場合、Todoが更新されること",
			input: graphmodel.TodoUpdateInput{
				ID:          validId,
				Title:       title,
				Description: description,
				Deadline:    deadline,
				Priority:    priority,
			},
			wantErr: false,
		},
		{
			name: "正常系: タイトルのみが有効な場合、Todoが更新されること",
			input: graphmodel.TodoUpdateInput{
				ID:          validId,
				Title:       title,
				Description: nil,
				Deadline:    nil,
				Priority:    nil,
			},
			wantErr: false,
		},
		{
			name: "異常系: 存在しないTodo IDの場合、エラーが返されること",
			input: graphmodel.TodoUpdateInput{
				ID:          notFoundId,
				Title:       title,
				Description: description,
				Deadline:    deadline,
				Priority:    priority,
			},
			wantErr: true,
		},
		{
			name: "異常系: 取得リポジトリでエラーが発生した場合、エラーが返されること",
			input: graphmodel.TodoUpdateInput{
				ID:          errCaseInFindId,
				Title:       title,
				Description: description,
				Deadline:    deadline,
				Priority:    priority,
			},
			wantErr: true,
		},
		{
			name: "異常系: 保存リポジトリでエラーが発生した場合、エラーが返されること",
			input: graphmodel.TodoUpdateInput{
				ID:          errCaseInSaveId,
				Title:       title,
				Description: description,
				Deadline:    deadline,
				Priority:    priority,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.UpdateTodoDetail(context, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
