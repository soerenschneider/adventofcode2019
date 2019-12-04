package day04

import "testing"

func Test_isAdjacent(t *testing.T) {
	tests := []struct {
		name string
		candidate string
		want bool
	}{
		{
			name: "Positive",
			candidate: "0123455",
			want: true,
		},
		{
			name: "Positive",
			candidate: "0023455",
			want: true,
		},
		{
			name: "Positive",
			candidate: "01233456",
			want: true,
		},
		{
			name: "Negative",
			candidate: "0123456",
			want: false,
		},
		{
			name: "Extremes",
			candidate: "",
			want: false,
		},
		{
			name: "Extremes",
			candidate: "aa123456",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdjacent(tt.candidate); got != tt.want {
				t.Errorf("isAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIncreasing(t *testing.T) {
	tests := []struct {
		name string
		candidate string
		want bool
	}{
		{
			name: "Positive",
			candidate: "123456",
			want: true,
		},
		{
			name: "Positive",
			candidate: "1234556",
			want: true,
		},
		{
			name: "Positive",
			candidate: "0000000",
			want: true,
		},
		{
			name: "Negative",
			candidate: "1234546",
			want: false,
		},
		{
			name: "Extremes",
			candidate: "a12",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIncreasing(tt.candidate); got != tt.want {
				t.Errorf("isIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_meetsRules(t *testing.T) {
	tests := []struct {
		name string
		candidate string
		want bool
	}{
		{
			name: "Positive",
			candidate: "111111",
			want: true,
		},
		{
			name: "Negative",
			candidate: "223450",
			want: false,
		},
		{
			name: "Negative",
			candidate: "123789",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := meetsRules1st(tt.candidate); got != tt.want {
				t.Errorf("meetsRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_meetsRulesb(t *testing.T) {
	tests := []struct {
		name string
		candidate string
		want bool
	}{
		{
			name: "Positive",
			candidate: "111111",
			want: false,
		},
		{
			name: "Negative",
			candidate: "223450",
			want: false,
		},
		{
			name: "Negative",
			candidate: "123789",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := meetsRules2nd(tt.candidate); got != tt.want {
				t.Errorf("meetsRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAdjacentb(t *testing.T) {
	tests := []struct {
		name string
		candidate string
		want bool
	}{
		{
			name: "Positive",
			candidate: "112233",
			want: true,
		},
		{
			name: "Positive",
			candidate: "111122",
			want: true,
		},
		{
			name: "Positive",
			candidate: "133122",
			want: true,
		},
		{
			name: "Positive",
			candidate: "13322",
			want: true,
		},
		{
			name: "Positive",
			candidate: "123444",
			want: false,
		},
		{
			name: "Positive",
			candidate: "123444",
			want: false,
		},
		{
			name: "Positive",
			candidate: "0008577",
			want: true,
		},
		{
			name: "Positive",
			candidate: "0085777",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdjacent2nd(tt.candidate); got != tt.want {
				t.Errorf("isAdjacent2nd() = %v, want %v", got, tt.want)
			}
		})
	}
}