package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		wantErr      bool
		wantPassword Password
	}{
		{
			name:         "valid password with 16 characters",
			value:        "1234567890123456",
			wantErr:      false,
			wantPassword: Password("1234567890123456"),
		},
		{
			name:         "valid password with special characters",
			value:        "Pass@123!",
			wantErr:      false,
			wantPassword: Password("Pass@123!"),
		},
		{
			name:         "valid password with numbers and letters",
			value:        "password123",
			wantErr:      false,
			wantPassword: Password("password123"),
		},
		{
			name:    "empty password",
			value:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := NewPassword(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Password(""), password)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantPassword, password)
			}
		})
	}
}

func TestPassword_Value(t *testing.T) {
	tests := []struct {
		name     string
		password Password
		expected string
	}{
		{
			name:     "normal password",
			password: Password("password123"),
			expected: "password123",
		},
		{
			name:     "empty password",
			password: Password(""),
			expected: "",
		},
		{
			name:     "password with special characters",
			password: Password("Pass@123!"),
			expected: "Pass@123!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.password.Value()
			assert.Equal(t, tt.expected, result)
		})
	}
}
