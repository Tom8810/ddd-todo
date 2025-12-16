package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		wantErr  bool
		wantName Name
	}{
		{
			name:     "valid name with 1 character",
			value:    "A",
			wantErr:  false,
			wantName: Name("A"),
		},
		{
			name:     "valid name with Japanese characters",
			value:    "田中太郎",
			wantErr:  false,
			wantName: Name("田中太郎"),
		},
		{
			name:     "valid name with special characters",
			value:    "O'Connor-Smith",
			wantErr:  false,
			wantName: Name("O'Connor-Smith"),
		},
		{
			name:     "valid name with numbers",
			value:    "John2",
			wantErr:  false,
			wantName: Name("John2"),
		},
		{
			name:     "max length 64",
			value:    string(make([]byte, 64)),
			wantErr:  false,
			wantName: Name(string(make([]byte, 64))),
		},
		{
			name:    "over max length 65",
			value:   string(make([]byte, 65)),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, err := NewName(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Name(""), name)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantName, name)
			}
		})
	}
}

func TestName_Value(t *testing.T) {
	tests := []struct {
		name     string
		nameVO   Name
		expected string
	}{
		{
			name:     "normal name",
			nameVO:   Name("John Doe"),
			expected: "John Doe",
		},
		{
			name:     "empty name",
			nameVO:   Name(""),
			expected: "",
		},
		{
			name:     "name with special characters",
			nameVO:   Name("O'Connor"),
			expected: "O'Connor",
		},
		{
			name:     "Japanese name",
			nameVO:   Name("田中太郎"),
			expected: "田中太郎",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.nameVO.Value()
			assert.Equal(t, tt.expected, result)
		})
	}
}
