package day13

import "testing"

func Test_showScore(t *testing.T) {
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{0, 0},
			want: false,
		},
		{
			args: args{-1, 0},
			want: true,
		},
		{
			args: args{0, -1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := showScore(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("showScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getJoystickInput(t *testing.T) {
	type args struct {
		paddle int
		x      int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				paddle: 0,
				x: 0,
			},
			want: 0,
		},
		{
			args: args{
				paddle: 1,
				x: 0,
			},
			want: -1,
		},
		{
			args: args{
				paddle: 1,
				x: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getJoystickInput(tt.args.paddle, tt.args.x); got != tt.want {
				t.Errorf("getJoystickInput() = %v, want %v", got, tt.want)
			}
		})
	}
}