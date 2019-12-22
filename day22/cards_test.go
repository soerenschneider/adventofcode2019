package day22

import (
	"reflect"
	"testing"
)

func Test_cards_Cut(t *testing.T) {
	type fields struct {
		deck []int
	}
	tests := []struct {
		name   string
		fields fields
		args   int
		want   []int
	}{
		{
			fields: fields{
				deck: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			args: 3,
			want: []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
		{
			fields: fields{
				deck: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			args: -4,
			want: []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cards{
				deck: tt.fields.deck,
			}
			c.Cut(tt.args)
			got := c.deck
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cards_Deal(t *testing.T) {
	type fields struct {
		deck []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			fields: fields{
				deck: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cards{
				deck: tt.fields.deck,
			}
			c.Deal()
			got := c.deck
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cards_Increment(t *testing.T) {
	type fields struct {
		deck []int
	}
	tests := []struct {
		name   string
		fields fields
		args   int
		want   []int
	}{
		{
			fields: fields{
				deck: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			args: 3,
			want: []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cards{
				deck: tt.fields.deck,
			}
			c.Increment(tt.args)
			got := c.deck
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseArg(t *testing.T) {
	type args struct {
		instr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				instr: "deal with increment 20",
			},
			want: 20,
		},
		{
			args: args{
				instr: "cut -2565",
			},
			want: -2565,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseArg(tt.args.instr); got != tt.want {
				t.Errorf("parseArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_perform(t *testing.T) {
	type args struct {
		file string
		n    int
	}
	tests := []struct {
		name string
		args args
		want *cards
	}{
		{
			args: args{
				file: "../resources/day22/test1.txt",
				n: 10,
			},
			want: &cards{
					deck: []int {0,3,6,9,2,5,8,1,4,7},
			},
		},
		{
			args: args{
				file: "../resources/day22/test2.txt",
				n: 10,
			},
			want: &cards{
				deck: []int {3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
			},
		},
		{
			args: args{
				file: "../resources/day22/test3.txt",
				n: 10,
			},
			want: &cards{
				deck: []int {6,3,0,7,4,1,8,5,2,9},
			},
		},
		{
			args: args{
				file: "../resources/day22/test4.txt",
				n: 10,
			},
			want: &cards{
				deck: []int {9,2,5,8,1,4,7,0,3,6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadInstructionsAndShuffle(tt.args.file, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadInstructionsAndShuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}