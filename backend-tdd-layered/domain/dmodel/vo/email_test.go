package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantErr   bool
		wantEmail Email
	}{
		{
			name:      "valid email with standard format",
			value:     "test@example.com",
			wantErr:   false,
			wantEmail: Email("test@example.com"),
		},
		{
			name:      "valid email with subdomain",
			value:     "user@mail.example.com",
			wantErr:   false,
			wantEmail: Email("user@mail.example.com"),
		},
		{
			name:      "valid email with numbers",
			value:     "user123@example123.com",
			wantErr:   false,
			wantEmail: Email("user123@example123.com"),
		},
		{
			name:      "valid email with special characters",
			value:     "user.name+tag@example.com",
			wantErr:   false,
			wantEmail: Email("user.name+tag@example.com"),
		},
		{
			name:    "invalid email without @",
			value:   "testexample.com",
			wantErr: true,
		},
		{
			name:    "invalid email without domain",
			value:   "test@",
			wantErr: true,
		},
		{
			name:    "invalid email without local part",
			value:   "@example.com",
			wantErr: true,
		},
		{
			name:    "invalid email without TLD",
			value:   "test@example",
			wantErr: true,
		},
		{
			name:    "empty email",
			value:   "",
			wantErr: true,
		},
		{
			name:    "invalid email with spaces",
			value:   "test @example.com",
			wantErr: true,
		},
		{
			name:    "invalid email with multiple @",
			value:   "test@@example.com",
			wantErr: true,
		},
		{
			name:    "invalid email with invalid TLD",
			value:   "test@example.c",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := NewEmail(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Email(""), email)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantEmail, email)
			}
		})
	}
}

func TestEmail_Value(t *testing.T) {
	tests := []struct {
		name     string
		email    Email
		expected string
	}{
		{
			name:     "normal email",
			email:    Email("test@example.com"),
			expected: "test@example.com",
		},
		{
			name:     "empty email",
			email:    Email(""),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.email.Value()
			assert.Equal(t, tt.expected, result)
		})
	}
}
