package day01

import (
	"math"
	"testing"
)

func TestFuelCalculator(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		n int
		want int
	}{
		{
			name: "Example 1",
			n: 1969,
			want: 966,
		},
		{
			name: "Example 2",
			n: 100756,
			want: 50346,
		},
		{
			name: "Extreme: zero",
			n: 0,
			want: 0,
		},
		{
			name: "Extreme: negative",
			n: -1,
			want: 0,
		},
		{
			name: "Extreme: max",
			n: math.MaxInt32,
			want: 1073741757,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Calc01b{}
			if got := m.RequiredFuel(tt.n); got != tt.want {
				t.Errorf("RequiredFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
