package day17

import (
	"github.com/soerenschneider/adventofcode2019/util"
	"reflect"
	"testing"
)

func Test_interpretInput(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{
			input: 35,
			want:  scaffold,
		},
		{
			input: 46,
			want:  empty,
		},
		{
			input: 10,
			want:  newline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interpretInput(tt.input); got != tt.want {
				t.Errorf("interpretInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bot_sum(t *testing.T) {
	type fields struct {
		field      [][]string
		currentRow int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				field: [][]string{
					{".", ".", "#",},
					{".", "#", "#",},
					{".", ".", "#",},
				},
			},
			want: 0,
		},
		{
			fields: fields{
				field: [][]string{
					{".", ".", "#", "."},
					{".", ".", "#", "."},
					{".", "#", "#", "#"},
					{".", ".", "#", "."},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bot{
				field:      tt.fields.field,
				currentRow: tt.fields.currentRow,
			}
			if got := b.SumOfIntersections(); got != tt.want {
				t.Errorf("SumOfIntersections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIntersection(t *testing.T) {
	type args struct {
		y     int
		x     int
		field [][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
				x: 1,
				y: 2,
			},
			want: true,
		},
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
				x: 2,
				y: 2,
			},
			want: false,
		},
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
				x: 0,
				y: 0,
			},
			want: false,
		},
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
				x: 3,
				y: 3,
			},
			want: false,
		},
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
				x: 4,
				y: 4,
			},
			want: false,
		},
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
				x: 4,
				y: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIntersection(tt.args.x, tt.args.y, tt.args.field); got != tt.want {
				t.Errorf("isIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determinePosition(t *testing.T) {
	type args struct {
		field [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    []util.Coordinate
		wantErr bool
	}{
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "^"},
					{".", "#", ".", "."},
				},
			},
			want: []util.Coordinate{
				{3, 2},
				North,
			},
			wantErr: false,
		},
		{
			args: args{
				field: [][]string{
					{".", "v", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
			},
			want: []util.Coordinate{
				{1, 0},
				South,
			},
			wantErr: false,
		},
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := determinePositionAndDirection(tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("determinePositionAndDirection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("determinePositionAndDirection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isScaffold(t *testing.T) {
	type args struct {
		pos   util.Coordinate
		field [][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
				pos: util.Coordinate{0, 0},
			},
			want: false,
		},
		{
			args: args{
				field: [][]string{
					{".", "#", ".", "."},
					{".", "#", ".", "."},
					{"#", "#", "#", "#"},
					{".", "#", ".", "."},
				},
				pos: util.Coordinate{1, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScaffold(tt.args.pos, tt.args.field); got != tt.want {
				t.Errorf("isScaffold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bot_discoverPath(t *testing.T) {
	type fields struct {
		field      [][]string
		direction  util.Coordinate
		position   util.Coordinate
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			fields: fields{
				field: [][]string{
					{"#", "#", "#", ".", "v"},
					{"#", ".", "#", ".", "#"},
					{"#", "#", "#", "#", "#"},
					{".", ".", "#", ".", "."},
				},
				position: util.Coordinate{4,0},
				direction: South,
			},
			want: []string{"2", "R", "4", "R", "2", "R","2", "R", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bot{
				field:      tt.fields.field,
				direction:  tt.fields.direction,
				position:   tt.fields.position,
			}
			if got := b.discoverPath(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("discoverPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compressPath(t *testing.T) {
	type args struct {
		path      []string
		fragments [][]string
		functions [][]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][4][]string
	}{
		{
			args: args {
				path: []string{"R","8","R","8","R","4","R","4","R","8","L","6","L","2","R","4","R","4","R","8","R","8","R","8","L","6","L","2"},
				fragments: [][]string{[]string{"R","8","R","8","R","4","R","4","R","8","L","6","L","2","R","4","R","4","R","8","R","8","R","8","L","6","L","2"}},
				functions: nil,
			},
			wantResult: [][4][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := compressPath(tt.args.path, tt.args.fragments, tt.args.functions); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("compressPath() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}