package service_test

import (
	"context"
	"testing"

	service "github.com/ddd-todo/project-backend/application/service/todo"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func TestGetTodoDetail(t *testing.T) {
	repo := repository.NewRepositoriesMock()
	service := service.NewTodoService(repo)
	ctx := context.Background()

	validId := lib.GenerateID()
	notFoundId := "bb04f183-fb98-6d72-eae2-8633ed7a5c2d"
	errCaseId := "c9f374b7-7c97-2529-de08-f091c9c9921c"

	tests := []struct {
		name    string
		input   string
		wantErr bool
		want    int
	}{
		{
			name:    "Valid Todo ID",
			input:   validId,
			wantErr: false,
			want:    1,
		},
		{
			name:    "Invalid Todo ID",
			input:   "invalid-id",
			wantErr: true,
			want:    0,
		},
		{
			name:    "Non-existent Todo ID",
			input:   notFoundId,
			wantErr: false,
			want:    0,
		},
		{
			name:    "Repository error case",
			input:   errCaseId,
			wantErr: true,
			want:    0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetTodoDetail(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTodoDetail() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want > 0 && got == nil {
				t.Errorf("GetTodoDetail() got = nil, want non-nil")
				return
			}
		})
	}
}
