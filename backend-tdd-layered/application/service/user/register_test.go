package service_test

import (
	"context"
	"testing"

	service "github.com/ddd-todo/project-backend/application/service/user"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
)

func TestRegister(t *testing.T) {
	repo := repository.NewRepositoriesMock()
	service := service.NewUserService(repo)
	ctx := context.Background()

	name := "Test User"
	email := "demo@example.com"
	existingEmail := "existing@example.com"
	errorEmail := "error@example.com"
	password := "password"

	tests := []struct {
		name    string
		input   graphmodel.UserCreateInput
		wantErr bool
	}{
		{
			name: "Successful Registration",
			input: graphmodel.UserCreateInput{
				Name:     name,
				Email:    email,
				Password: password,
			},
			wantErr: false,
		},
		{
			name: "Registration with Existing Email",
			input: graphmodel.UserCreateInput{
				Name:     name,
				Email:    existingEmail,
				Password: password,
			},
			wantErr: true,
		},
		{
			name: "Repository Error",
			input: graphmodel.UserCreateInput{
				Name:     name,
				Email:    errorEmail,
				Password: password,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.Register(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
