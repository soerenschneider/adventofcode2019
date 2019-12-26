package util

import "testing"

func Test_IsAscii(t *testing.T) {
	tests := []struct {
		name string
		args int64
		want bool
	}{
		{
			args: -1,
			want: false,
		},
		{
			args: 0,
			want: true,
		},
		{
			args: 32,
			want: true,
		},
		{
			args: 128,
			want: true,
		},
		{
			args: 129,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAscii(tt.args); got != tt.want {
				t.Errorf("isAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

