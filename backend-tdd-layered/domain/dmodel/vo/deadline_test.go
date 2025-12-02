package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDeadline(t *testing.T) {
	now := time.Now()
	futureTime := now.Add(24 * time.Hour)
	pastTime := now.Add(-24 * time.Hour)

	tests := []struct {
		name         string
		value        time.Time
		wantErr      bool
		wantDeadline Deadline
	}{
		{
			name:         "valid deadline in future",
			value:        futureTime,
			wantErr:      false,
			wantDeadline: Deadline(futureTime),
		},
		{
			name:    "invalid deadline in past",
			value:   pastTime,
			wantErr: true,
		},
		{
			name:    "invalid deadline with zero time",
			value:   time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deadline, err := NewDeadline(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Deadline(time.Time{}), deadline)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantDeadline, deadline)
			}
		})
	}
}

func TestDeadline_Value(t *testing.T) {
	now := time.Now()
	futureTime := now.Add(24 * time.Hour)

	tests := []struct {
		name     string
		deadline Deadline
		expected time.Time
	}{
		{
			name:     "normal deadline",
			deadline: Deadline(futureTime),
			expected: futureTime,
		},
		{
			name:     "current time deadline",
			deadline: Deadline(now),
			expected: now,
		},
		{
			name:     "zero deadline",
			deadline: Deadline(time.Time{}),
			expected: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.deadline.Value()
			assert.Equal(t, tt.expected, result)
		})
	}
}
