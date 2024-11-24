package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       float64
		wantErr    bool
	}{
		{
			name:       "simple addition",
			expression: "2 + 3",
			want:       5,
			wantErr:    false,
		},
		{
			name:       "complex expression",
			expression: "3 * (4 + 2) / 2",
			want:       9,
			wantErr:    false,
		},
		{
			name:       "negative numbers",
			expression: "-5 + 3",
			want:       -2,
			wantErr:    false,
		},
		{
			name:       "decimal numbers",
			expression: "2.5 * 4",
			want:       10,
			wantErr:    false,
		},
		{
			name:       "invalid expression",
			expression: "2 + * 3",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "empty expression",
			expression: "",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "division by zero",
			expression: "5 / 0",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "unmatched parentheses",
			expression: "2 * (3 + 4",
			want:       0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calc(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
