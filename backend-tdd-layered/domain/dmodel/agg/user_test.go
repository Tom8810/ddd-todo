package agg

import (
	"testing"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func TestNewUser(t *testing.T) {
	id, _ := vo.NewUserID("user-123")
	email, _ := vo.NewEmail("test@example.com")
	password, _ := vo.NewPassword("password")
	name, _ := vo.NewName("Test User")
	tests := []struct {
		name  string
		input struct {
			id       vo.UserID
			email    vo.Email
			password vo.Password
			name     vo.Name
		}
		want *User
	}{
		{
			name: "Valid user creation",
			input: struct {
				id       vo.UserID
				email    vo.Email
				password vo.Password
				name     vo.Name
			}{
				id:       id,
				email:    email,
				password: password,
				name:     name,
			},
			want: &User{
				ID:       id,
				Email:    email,
				Password: password,
				Name:     name,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := NewUser(tt.input.id, tt.input.name, tt.input.email, tt.input.password)

			if got.ID != tt.want.ID {
				t.Errorf("NewUser() ID = %v, want %v", got.ID, tt.want.ID)
			}
			if got.Email != tt.want.Email {
				t.Errorf("NewUser() Email = %v, want %v", got.Email, tt.want.Email)
			}
			if got.Password != tt.want.Password {
				t.Errorf("NewUser() Password = %v, want %v", got.Password, tt.want.Password)
			}
			if got.Name != tt.want.Name {
				t.Errorf("NewUser() Name = %v, want %v", got.Name, tt.want.Name)
			}
		})
	}
}
