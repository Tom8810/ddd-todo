package vo

import (
	"testing"
)

func TestNewTitle(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid title",
			input:   "My Todo Title",
			wantErr: false,
		},
		{
			name:    "Empty title",
			input:   "",
			wantErr: true,
		},
		{
			name:    "max length 64",
			input:   string(make([]byte, 64)),
			wantErr: false,
		},
		{
			name:    "over max length 65",
			input:   string(make([]byte, 65)),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTitle(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTitle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
