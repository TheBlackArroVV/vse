package utils_test

import (
	"testing"
	"vse/utils"
)

func TestTransformStrings(t *testing.T) {
	tests := []struct {
		name   string
		string string
		want   string
	}{
		{
			name:   "downcase all capital letters",
			string: "ASDF",
			want:   "asdf",
		},
		{
			name:   "repalces newlines with spaces",
			string: "asdf\nasdf\nzxcv",
			want:   "asdf asdf zxcv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.TransformStrings(tt.string)
			if tt.want != got {
				t.Errorf("TransformStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
