package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPriority(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		wantErr      bool
		wantPriority Priority
	}{
		{
			name:         "valid priority LOW",
			value:        "LOW",
			wantErr:      false,
			wantPriority: PriorityLow,
		},
		{
			name:         "valid priority MEDIUM",
			value:        "MEDIUM",
			wantErr:      false,
			wantPriority: PriorityMedium,
		},
		{
			name:         "valid priority HIGH",
			value:        "HIGH",
			wantErr:      false,
			wantPriority: PriorityHigh,
		},
		{
			name:    "invalid priority with lowercase",
			value:   "low",
			wantErr: true,
		},
		{
			name:    "invalid priority with mixed case",
			value:   "Low",
			wantErr: true,
		},
		{
			name:    "invalid priority value",
			value:   "URGENT",
			wantErr: true,
		},
		{
			name:    "empty priority",
			value:   "",
			wantErr: true,
		},
		{
			name:    "priority with spaces",
			value:   "LOW ",
			wantErr: true,
		},
		{
			name:    "priority with special characters",
			value:   "LOW!",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			priority, err := NewPriority(tt.value)
			
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Priority(""), priority)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantPriority, priority)
			}
		})
	}
}

func TestPriority_Value(t *testing.T) {
	tests := []struct {
		name     string
		priority Priority
		expected string
	}{
		{
			name:     "LOW priority",
			priority: PriorityLow,
			expected: "LOW",
		},
		{
			name:     "MEDIUM priority",
			priority: PriorityMedium,
			expected: "MEDIUM",
		},
		{
			name:     "HIGH priority",
			priority: PriorityHigh,
			expected: "HIGH",
		},
		{
			name:     "empty priority",
			priority: Priority(""),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.priority.Value()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPriority_Constants(t *testing.T) {
	// Test that the priority constants are defined correctly
	assert.Equal(t, "LOW", PriorityLow.Value())
	assert.Equal(t, "MEDIUM", PriorityMedium.Value())
	assert.Equal(t, "HIGH", PriorityHigh.Value())
}