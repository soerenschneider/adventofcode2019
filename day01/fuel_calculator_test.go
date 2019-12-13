package day01

import (
	"math"
	"testing"
)

func TestProcessInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    int
		wantErr bool
	}{
		{
			name: "Basic",
			input: []int{12, 14, 1969, 100756},
			want: 34241,
			wantErr: false,
		},
		{
			name: "Extremes: nil",
			input: nil,
			want: -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

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


