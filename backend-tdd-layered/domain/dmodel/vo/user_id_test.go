package vo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserID(t *testing.T) {
	validUUID := uuid.New().String()

	tests := []struct {
		name       string
		value      string
		wantErr    bool
		wantUserID UserID
	}{
		{
			name:       "valid UUID v4",
			value:      validUUID,
			wantErr:    false,
			wantUserID: UserID(validUUID),
		},
		{
			name:       "valid UUID v1",
			value:      "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			wantErr:    false,
			wantUserID: UserID("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		},
		{
			name:       "valid nil UUID",
			value:      "00000000-0000-0000-0000-000000000000",
			wantErr:    false,
			wantUserID: UserID("00000000-0000-0000-0000-000000000000"),
		},
		{
			name:    "invalid UUID format (wrong length)",
			value:   "6ba7b810-9dad-11d1-80b4-00c04fd430c",
			wantErr: true,
		},
		{
			name:    "invalid UUID with extra characters",
			value:   "6ba7b810-9dad-11d1-80b4-00c04fd430c8x",
			wantErr: true,
		},
		{
			name:    "empty string",
			value:   "",
			wantErr: true,
		},
		{
			name:    "invalid UUID with letters out of range",
			value:   "6ba7b810-9dad-11d1-80b4-00c04fd430cg",
			wantErr: true,
		},
		{
			name:    "random string",
			value:   "not-a-uuid",
			wantErr: true,
		},
		{
			name:    "UUID with spaces",
			value:   "6ba7b810-9dad-11d1-80b4-00c04fd430c8 ",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userID, err := NewUserID(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, UserID(""), userID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantUserID, userID)
			}
		})
	}
}

func TestUserID_Value(t *testing.T) {
	validUUID := uuid.New().String()

	tests := []struct {
		name     string
		userID   UserID
		expected string
	}{
		{
			name:     "normal UUID",
			userID:   UserID(validUUID),
			expected: validUUID,
		},
		{
			name:     "empty UserID",
			userID:   UserID(""),
			expected: "",
		},
		{
			name:     "nil UUID",
			userID:   UserID("00000000-0000-0000-0000-000000000000"),
			expected: "00000000-0000-0000-0000-000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.userID.Value()
			assert.Equal(t, tt.expected, result)
		})
	}
}
