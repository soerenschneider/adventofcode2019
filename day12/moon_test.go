package day12

import (
	"reflect"
	"testing"
)

func TestVector_Sum(t *testing.T) {
	type fields struct {
		X int
		Y int
		Z int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				X: -1,
				Y: -2,
				Z: -3,
			},
			want: 6,
		},
		{
			fields: fields{
				X: 1,
				Y: 2,
				Z: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoon_Gravity(t *testing.T) {
	tests := []struct {
		name      string
		moon      Moon
		other     Moon
		wantMoon  Moon
		wantOther Moon
	}{
		{
			moon: Moon{
				position: Vector{1, 2, 3},
				velocity: Vector{0, 0, 0},
			},
			other: Moon{
				position: Vector{10, -20, 30},
				velocity: Vector{0, 0, 0},
			},
			wantMoon: Moon{
				position: Vector{1, 2, 3},
				velocity: Vector{1, -1, 1},
			},
			wantOther: Moon{
				position: Vector{10, -20, 30},
				velocity: Vector{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.moon.ApplyGravity(tt.other)

			if !reflect.DeepEqual(tt.moon, tt.wantMoon) {
				t.Errorf("moon, got = %v, want %v", tt.moon, tt.wantMoon)
			}

			if !reflect.DeepEqual(tt.other, tt.wantOther) {
				t.Errorf("other, got = %v, want %v", tt.other, tt.wantOther)
			}
		})
	}
}

func TestMoon_PotentialEnergy(t *testing.T) {
	type fields struct {
		position Vector
		velocity Vector
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				position: Vector{2, 1, 3},
				velocity: Vector{3, 2, 1},
			},
			want: 36,
		},
		{
			fields: fields{
				position: Vector{1, 8, 0},
				velocity: Vector{1, 1, 3},
			},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moon{
				position: tt.fields.position,
				velocity: tt.fields.velocity,
			}
			if got := m.PotentialEnergy(); got != tt.want {
				t.Errorf("PotentialEnergy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoon_Step(t *testing.T) {
	type fields struct {
		position Vector
		velocity Vector
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			fields: fields{
				position: Vector{5, 5, 5},
				velocity: Vector{-1, 0, 1},
			},
			want: fields{
				position: Vector{4, 5, 6},
				velocity: Vector{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moon := &Moon{
				position: tt.fields.position,
				velocity: tt.fields.velocity,
			}
			want := &Moon{
				position: tt.want.position,
				velocity: tt.want.velocity,
			}
			moon.ApplyVelocity()

			if !reflect.DeepEqual(moon, want) {
				t.Errorf("moon, got = %v, want %v", moon, want)
			}
		})
	}
}

func TestSteps(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		moons []Moon
		steps int
		want  []Moon
	}{
		{
			moons: []Moon{
				Moon{
					position: Vector{-1, 0, 2},
				},
				Moon{
					position: Vector{2, -10, -7},
				},
				Moon{
					position: Vector{4, -8, 8},
				},
				Moon{
					position: Vector{3, 5, -1},
				},
			},
			steps: 10,
			want: []Moon{
				Moon{
					position: Vector{2, 1, -3},
					velocity: Vector{-3, -2, 1},
				},
				Moon{
					position: Vector{1, -8, 0},
					velocity: Vector{-1, 1, 3},
				},
				Moon{
					position: Vector{3, -6, 1},
					velocity: Vector{3, 2, -3},
				},
				Moon{
					position: Vector{2, 0, 4},
					velocity: Vector{1, -1, -1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Cycle(tt.moons, tt.steps)
			if !reflect.DeepEqual(tt.moons, tt.want) {
				t.Errorf("moon, got = %v, want %v", tt.moons, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		moons []Moon
	}
	tests := []struct {
		name    string
		moons   []Moon
		wantSum int
	}{
		{
			moons: []Moon{
				Moon{
					position: Vector{2, 1, -3},
					velocity: Vector{-3, -2, 1},
				},
				Moon{
					position: Vector{1, -8, 0},
					velocity: Vector{-1, 1, 3},
				},
				Moon{
					position: Vector{3, -6, 1},
					velocity: Vector{3, 2, -3},
				},
				Moon{
					position: Vector{2, 0, 4},
					velocity: Vector{1, -1, -1},
				},
			},
			wantSum: 179,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Sum(tt.moons); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestMoon_axisIdentical(t *testing.T) {
	type fields struct {
		position Vector
		velocity Vector
	}
	type args struct {
		other Moon
		axis  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				position: Vector{1, 2, 3},
				velocity: Vector{4, 5, 6},
			},
			args: args{
				other: Moon{
					position: Vector{1, 2, 3},
					velocity: Vector{4, 5, 6},
				},
				axis: X,
			},
			want: true,
		},
		{
			fields: fields{
				position: Vector{1, 2, 3},
				velocity: Vector{4, 5, 6},
			},
			args: args{
				other: Moon{
					position: Vector{2, 2, 3},
					velocity: Vector{4, 5, 6},
				},
				axis: X,
			},
			want: false,
		},
		{
			fields: fields{
				position: Vector{1, 2, 3},
				velocity: Vector{4, 5, 6},
			},
			args: args{
				other: Moon{
					position: Vector{3, 2, 3},
					velocity: Vector{4, 5, 6},
				},
				axis: Y,
			},
			want: true,
		},
		{
			fields: fields{
				position: Vector{1, 2, 3},
				velocity: Vector{4, 5, 6},
			},
			args: args{
				other: Moon{
					position: Vector{1, 5, 3},
					velocity: Vector{4, 5, 6},
				},
				axis: Y,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moon{
				position: tt.fields.position,
				velocity: tt.fields.velocity,
			}
			if got := m.HasIdenticalAxis(tt.args.other, tt.args.axis); got != tt.want {
				t.Errorf("HasIdenticalAxis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_HasNullValue(t *testing.T) {
	type fields struct {
		X int
		Y int
		Z int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{1, 1, 1},
			want:   false,
		},
		{
			fields: fields{1, 0, 1},
			want:   true,
		},
		{
			fields: fields{0, 0, 0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.HasNullValue(); got != tt.want {
				t.Errorf("HasNullValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateAxisInformation(t *testing.T) {
	type args struct {
		moon               Moon
		other              Moon
		axisAtInitialState []bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			args: args{
				moon: Moon{
					position: Vector{0, 0, 0},
					velocity: Vector{0, 0, 0},
				},
				other: Moon{
					position: Vector{1, 1, 1},
					velocity: Vector{0, 0, 0},
				},
				axisAtInitialState: []bool{
					true, true, true,
				},
			},
			want: []bool{false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateAxisInformation(tt.args.moon, tt.args.other, tt.args.axisAtInitialState)
			if !reflect.DeepEqual(tt.args.moon, tt.want) {
				t.Errorf("moon, got = %v, want %v", tt.args.axisAtInitialState, tt.want)
			}

		})
	}
}

func TestParseVectors(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []Vector
	}{
		{
			args: `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
			want: []Vector{
	{-1,0,2},
	{2,-10,-7},
	{4,-8,8},
	{3,5,-1},
			},
		},
		{
			args: "",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseVectors(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseVectors() = %v, want %v", got, tt.want)
			}
		})
	}
}
