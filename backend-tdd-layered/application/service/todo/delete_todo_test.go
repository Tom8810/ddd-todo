package service_test

import (
	"context"
	"testing"

	service "github.com/ddd-todo/project-backend/application/service/todo"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func TestDeleteTodo(t *testing.T) {
	repo := repository.NewRepositoriesMock()
	service := service.NewTodoService(repo)
	context := context.Background()

	validId := lib.GenerateID()
	notFoundId := "bb04f183-fb98-6d72-eae2-8633ed7a5c2d"
	errCaseId := "c9f374b7-7c97-2529-de08-f091c9c9921c"

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid Todo ID",
			input:   validId,
			wantErr: false,
		},
		{
			name:    "Invalid Todo ID",
			input:   "invalid-id",
			wantErr: true,
		},
		{
			name:    "Non-existent Todo ID",
			input:   notFoundId,
			wantErr: true,
		},
		{
			name:    "Repository error case",
			input:   errCaseId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.DeleteTodo(context, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
