package utils_test

import (
	"elastic_go/utils"
	"testing"
)

func TestAbsInt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		number int
		want   int
	}{
		{
			name:   "when number is 0",
			number: 0,
			want:   0,
		},
		{
			name:   "when number is positive",
			number: 10,
			want:   10,
		},
		{
			name:   "when number is negative",
			number: -10,
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.AbsInt(tt.number)
			if got != tt.want {
				t.Errorf("AbsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
