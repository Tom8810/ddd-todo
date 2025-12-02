package vo

import "testing"

func TestStatus(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid status - waiting",
			input:   "WAITING",
			wantErr: false,
		},
		{
			name:    "Valid status - doing",
			input:   "DOING",
			wantErr: false,
		},
		{
			name:    "Valid status - completed",
			input:   "COMPLETED",
			wantErr: false,
		},
		{
			name:    "Invalid status - empty",
			input:   "",
			wantErr: true,
		},
		{
			name:    "Invalid status - unknown",
			input:   "UNKNOWN",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewStatus(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
