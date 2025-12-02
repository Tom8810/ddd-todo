package vo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTodoID(t *testing.T) {
	validUUID := uuid.New().String()

	tests := []struct {
		name       string
		value      string
		wantErr    bool
		wantTodoID TodoID
	}{
		{
			name:       "valid UUID v4",
			value:      validUUID,
			wantErr:    false,
			wantTodoID: TodoID(validUUID),
		},
		{
			name:       "valid UUID v1",
			value:      "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			wantErr:    false,
			wantTodoID: TodoID("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		},
		{
			name:       "valid nil UUID",
			value:      "00000000-0000-0000-0000-000000000000",
			wantErr:    false,
			wantTodoID: TodoID("00000000-0000-0000-0000-000000000000"),
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
			todoID, err := NewTodoID(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, TodoID(""), todoID)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantTodoID, todoID)
			}
		})
	}
}

func TestTodoID_Value(t *testing.T) {
	validUUID := uuid.New().String()

	tests := []struct {
		name     string
		todoID   TodoID
		expected string
	}{
		{
			name:     "normal UUID",
			todoID:   TodoID(validUUID),
			expected: validUUID,
		},
		{
			name:     "empty TodoID",
			todoID:   TodoID(""),
			expected: "",
		},
		{
			name:     "nil UUID",
			todoID:   TodoID("00000000-0000-0000-0000-000000000000"),
			expected: "00000000-0000-0000-0000-000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.todoID.Value()
			assert.Equal(t, tt.expected, result)
		})
	}
}
