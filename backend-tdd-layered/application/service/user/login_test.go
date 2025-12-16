package service_test

import (
	"context"
	"testing"

	service "github.com/ddd-todo/project-backend/application/service/user"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
)

func TestLogin(t *testing.T) {
	repo := repository.NewRepositoriesMock()
	service := service.NewUserService(repo)
	ctx := context.Background()

	validEmail := "demo@example.com"
	invalidEmail := "invalid-email"
	notfoundEmail := "notfound@example.com"

	validPassword := "password"
	invalidPassword := "wrongpassword"

	tests := []struct {
		name    string
		input   graphmodel.LoginInput
		wantErr bool
	}{
		{
			name: "Valid Login",
			input: graphmodel.LoginInput{
				Email:    validEmail,
				Password: validPassword,
			},
			wantErr: false,
		},
		{
			name: "Invalid Email Format",
			input: graphmodel.LoginInput{
				Email:    invalidEmail,
				Password: validPassword,
			},
			wantErr: true,
		},
		{
			name: "Non-existent Email",
			input: graphmodel.LoginInput{
				Email:    notfoundEmail,
				Password: validPassword,
			},
			wantErr: true,
		},
		{
			name: "Incorrect Password",
			input: graphmodel.LoginInput{
				Email:    validEmail,
				Password: invalidPassword,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.Login(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
