package day01

import "testing"


func NaiveTestRequiredFuel(t *testing.T) {
	tests := []struct {
		name string
		arg int
		want int
	}{
		{
			name: "Example 1",
			arg: 12,
			want: 2,
		},
		{
			name: "Example 2",
			arg: 14,
			want: 2,
		},
		{
			name: "Example 3",
			arg: 1969,
			want: 654,
		},
		{
			name: "Example 4",
			arg: 100756,
			want: 33583,
		},
		{
			name: "Extremes: zero",
			arg: 0,
			want: -2,
		},
		{
			name: "Extremes: negative",
			arg: -1,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run("Acceptance", func(t *testing.T) {
			m := &Calc01{}
			if got := m.RequiredFuel(tt.arg); got != tt.want {
				t.Errorf("RequiredFuelNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}