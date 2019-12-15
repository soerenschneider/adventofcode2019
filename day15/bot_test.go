package day15

import (
	"github.com/soerenschneider/adventofcode2019/util"
	"reflect"
	"testing"
)

func TestBot_checkDeadEnd(t *testing.T) {
	type fields struct {
		position  util.Coordinate64
		oxygen    util.Coordinate64
		grid      Map
		direction int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			fields: fields {
				position: util.Coordinate64{0,0},
				grid: Map{
					util.Coordinate64{0, -1}: MapWall,
					util.Coordinate64{0, 1}: MapWall,
					util.Coordinate64{-1, 0}: MapWall,
					util.Coordinate64{1, 0}: MapWall,
				},
			},
			want: MapDeadEnd,
		},
		{
			fields: fields {
				position: util.Coordinate64{0,0},
				grid: Map{
					util.Coordinate64{0, -1}: MapWall,
					util.Coordinate64{0, 1}: MapEmptySpace,
					util.Coordinate64{-1, 0}: MapOxygen,
					util.Coordinate64{1, 0}: MapWall,
				},
			},
			want: MapUnknown,
		},
		{
			fields: fields {
				position: util.Coordinate64{0,0},
				grid: Map{
					util.Coordinate64{0, -1}: MapUnknown,
					util.Coordinate64{0, 1}: MapUnknown,
					util.Coordinate64{-1, 0}: MapUnknown,
					util.Coordinate64{1, 0}: MapUnknown,
				},
			},
			want: MapUnknown,
		},
		{
			fields: fields {
				position: util.Coordinate64{0,0},
				grid: Map{
					util.Coordinate64{0, -1}: MapUnknown,
					util.Coordinate64{0, 1}: MapUnknown,
					util.Coordinate64{-1, 0}: MapUnknown,
					util.Coordinate64{1, 0}: MapUnknown,
					util.Coordinate64{0, 0}: MapOxygen,
				},
			},
			want: MapOxygen,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bot{
				position:  tt.fields.position,
				oxygen:    tt.fields.oxygen,
				grid:      tt.fields.grid,
				direction: tt.fields.direction,
			}
			b.checkDeadEnd()
			if b.grid[b.position] != tt.want {
				t.Errorf("Expected %d, got %d", tt.want, b.grid[b.position])
			}
		})
	}
}

func TestBot_interpretResponse(t *testing.T) {
	type fields struct {
		position  util.Coordinate64
		oxygen    util.Coordinate64
		grid      Map
		direction int64
	}
	type args struct {
		response int64
		target   util.Coordinate64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOxygen util.Coordinate64
		wantPosition util.Coordinate64
	}{
		{
			args: args {
				response: ResponseFound,
				target: util.Coordinate64{5, 5},
			},
			wantOxygen: util.Coordinate64{5, 5},
			wantPosition: util.Coordinate64{5, 5},
		},
		{
			args: args {
				response: ResponseMoved,
				target: util.Coordinate64{1, 5},
			},
			wantOxygen: util.Coordinate64{0, 0},
			wantPosition: util.Coordinate64{1, 5},
		},
		{
			args: args {
				response: ResponseWall,
				target: util.Coordinate64{1, 5},
			},
			wantOxygen: util.Coordinate64{0, 0},
			wantPosition: util.Coordinate64{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bot{
				position:  tt.fields.position,
				oxygen:    tt.fields.oxygen,
				grid:      tt.fields.grid,
				direction: tt.fields.direction,
			}
			b.updatePosition(tt.args.response, tt.args.target)
			if !reflect.DeepEqual(b.oxygen, tt.wantOxygen) {
				t.Errorf("FUCK, got %v, want %v", b.oxygen, tt.wantOxygen)
			}
			b.updatePosition(tt.args.response, tt.args.target)
			if !reflect.DeepEqual(b.position, tt.wantPosition) {
				t.Errorf("FUCK, got %v, want %v", b.position, tt.wantPosition)
			}
		})
	}
}

func TestBot_tryMoveTo(t *testing.T) {
	type fields struct {
		position  util.Coordinate64
		oxygen    util.Coordinate64
		grid      Map
		direction int64
	}
	tests := []struct {
		name   string
		fields fields
		desiredPositionType int64
		want   bool
	}{
		{
			fields: fields {
				position: util.Coordinate64{0,0},
				grid: Map{
					util.Coordinate64{0, -1}: MapWall,
					util.Coordinate64{0, 1}: MapWall,
					util.Coordinate64{-1, 0}: MapWall,
					util.Coordinate64{1, 0}: MapWall,
				},
			},
			desiredPositionType: MapUnknown,
			want: false,
		},
		{
			fields: fields {
				position: util.Coordinate64{0,0},
				grid: Map{
					util.Coordinate64{0, -1}: MapWall,
					util.Coordinate64{0, 1}: MapWall,
					util.Coordinate64{-1, 0}: MapWall,
					util.Coordinate64{1, 0}: MapUnknown,
				},
			},
			desiredPositionType: MapUnknown,
			want: true,
		},
		{
			fields: fields {
				position: util.Coordinate64{0,0},
				grid: Map{
					util.Coordinate64{0, -1}: MapWall,
					util.Coordinate64{0, 1}: MapWall,
					util.Coordinate64{-1, 0}: MapWall,
					util.Coordinate64{1, 0}: MapWall,
					util.Coordinate64{2, 0}: MapUnknown,
				},
			},
			desiredPositionType: MapUnknown,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bot{
				position:  tt.fields.position,
				oxygen:    tt.fields.oxygen,
				grid:      tt.fields.grid,
				direction: tt.fields.direction,
			}
			if got := b.tryMoveTo(tt.desiredPositionType); got != tt.want {
				t.Errorf("tryMoveTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBot_addResponseToMap(t *testing.T) {
	type fields struct {
		position  util.Coordinate64
		oxygen    util.Coordinate64
		grid      Map
		direction int64
	}
	type args struct {
		response int64
		target   util.Coordinate64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantValue int64
	}{
		{
			fields: fields{
				position:  util.Coordinate64{0, 0},
				grid:      Map{
				},
			},
			args: args{
				response: ResponseWall,
				target:   util.Coordinate64{0, 1},
			},
			wantValue: MapWall,
		},
		{
			fields: fields{
				position:  util.Coordinate64{0, 0},
				grid:      Map{
				},
			},
			args: args{
				response: ResponseMoved,
				target:   util.Coordinate64{0, 1},
			},
			wantValue: MapEmptySpace,
		},
		{
			fields: fields{
				position:  util.Coordinate64{0, 0},
				grid:      Map{
				},
			},
			args: args{
				response: ResponseFound,
				target:   util.Coordinate64{0, 1},
			},
			wantValue: MapEmptySpace,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bot{
				position:  tt.fields.position,
				oxygen:    tt.fields.oxygen,
				grid:      tt.fields.grid,
				direction: tt.fields.direction,
			}
			b.setGridValue(tt.args.response, tt.args.target)
			if b.grid[tt.args.target] != tt.wantValue {
				t.Errorf("Expected value of %d, got %d", tt.wantValue, b.grid[tt.args.target])
			}
		})
	}
}