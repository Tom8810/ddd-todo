package service_test

import (
	"context"
	"testing"

	service "github.com/ddd-todo/project-backend/application/service/user"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func TestUpdateProperty(t *testing.T) {
	repo := repository.NewRepositoriesMock()
	service := service.NewUserService(repo)
	ctx := context.Background()

	validId := lib.GenerateID()
	notFoundId := "bb04f183-fb98-6d72-eae2-8633ed7a5c2d"
	errCaseId := "c9f374b7-7c97-2529-de08-f091c9c9921c"

	tests := []struct {
		name  string
		input struct {
			userID   string
			property string
		}
		field   string
		wantErr bool
	}{
		{
			name: "Email Valid User ID and Property",
			input: struct {
				userID   string
				property string
			}{
				userID:   validId,
				property: "demo@example.com",
			},
			field:   "email",
			wantErr: false,
		},
		{
			name: "Username Valid User ID and Property",
			input: struct {
				userID   string
				property string
			}{
				userID:   validId,
				property: "new_username",
			},
			field:   "username",
			wantErr: false,
		},
		{
			name: "Password Valid User ID and Property",
			input: struct {
				userID   string
				property string
			}{
				userID:   validId,
				property: "newpassword",
			},
			field:   "password",
			wantErr: false,
		},
		{
			name: "Invalid User ID",
			input: struct {
				userID   string
				property string
			}{
				userID:   "invalid-id",
				property: "some_value",
			},
			field:   "all",
			wantErr: true,
		},
		{
			name: "Non-existent User ID",
			input: struct {
				userID   string
				property string
			}{
				userID:   notFoundId,
				property: "some_value",
			},
			field:   "all",
			wantErr: true,
		},
		{
			name: "Repository error case",
			input: struct {
				userID   string
				property string
			}{
				userID:   errCaseId,
				property: "some_value",
			},
			field:   "all",
			wantErr: true,
		},
		{
			name: "Empty Property",
			input: struct {
				userID   string
				property string
			}{
				userID:   validId,
				property: "",
			},
			field:   "all",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.field == "email" || tt.field == "all" {
				_, err = service.UpdateUserEmail(ctx, tt.input.userID, tt.input.property)
			}
			if tt.field == "username" || tt.field == "all" {
				_, err = service.UpdateUserProfile(ctx, tt.input.userID, tt.input.property)
			}
			if tt.field == "password" || tt.field == "all" {
				_, err = service.UpdateUserPassword(ctx, tt.input.userID, tt.input.property)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProperty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
