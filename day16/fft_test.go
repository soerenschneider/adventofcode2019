package day16

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{"12345"},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			args: args{""},
			want: []int{},
		},
		{
			args: args{"asd"},
			want: []int{},
		},
		{
			args: args{"80871224585914546619083218645595"},
			want: []int{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		input     []int
		iteration int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				iteration: 1,
			},
			want: 4,
		},
		{
			args: args{
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				iteration: 2,
			},
			want: 8,
		},
		{
			args: args{
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				iteration: 3,
			},
			want: 2,
		},
		{
			args: args{
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				iteration: 4,
			},
			want: 2,
		},
		{
			args: args{
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				iteration: 5,
			},
			want: 6,
		},
		{
			args: args{
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				iteration: 6,
			},
			want: 1,
		},
		{
			args: args{
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				iteration: 7,
			},
			want: 5,
		},
		{
			args: args{
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				iteration: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Apply(tt.args.input, tt.args.iteration); got != tt.want {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhases(t *testing.T) {
	type args struct {
		input  []int
		phases int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				phases: 1,
			},
			want: []int{4, 8, 2, 2, 6, 1, 5, 8},
		},
		{
			args: args{
				input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				phases: 2,
			},
			want: []int{3, 4, 0, 4, 0, 4, 3, 8},
		},
		{
			args: args{
				input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				phases: 3,
			},
			want: []int{0, 3, 4, 1, 5, 5, 1, 8},
		},
		{
			args: args{
				input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				phases: 4,
			},
			want: []int{0, 1, 0, 2, 9, 4, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Phases(tt.args.input, tt.args.phases)
			if !reflect.DeepEqual(tt.args.input, tt.want) {
				t.Errorf("Wanted %v, got %v", tt.want, tt.args.input)
			}
		})
	}
}

func TestAnswer16(t *testing.T) {
	type args struct {
		input  []int
		phases int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				input:  []int{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5},
				phases: 100,
			},
			want: "24176176",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanSignal(tt.args.input, tt.args.phases); got != tt.want {
				t.Errorf("CleanSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}
