package day11

import (
	"github.com/soerenschneider/adventofcode2019/util"
	"reflect"
	"testing"
)

func Test_hull_move(t *testing.T) {
	type fields struct {
		direction int
		position util.Coordinate
	}
	tests := []struct {
		name   	string
		fields 	fields
		turn 	int64
		wantPos util.Coordinate
		wantDir int
	}{
		{
			fields: fields{
				position: util.Coordinate{2,2},
			},
			wantPos: util.Coordinate{1,2},
			turn: 0,
		},
		{
			fields: fields{
				position: util.Coordinate{1,2},
				direction: 1,
			},
			wantPos: util.Coordinate{1,3},
			turn: 0,
		},
		{
			fields: fields{
				position: util.Coordinate{1,3},
				direction: 3,
			},
			wantPos: util.Coordinate{1,2},
			turn: 0,
		},
		{
			fields: fields{
				position: util.Coordinate{1,3},
				direction: 3,
			},
			wantPos: util.Coordinate{1,4},
			turn: 1,
		},
		{
			fields: fields{
				position: util.Coordinate{2,2},
			},
			wantPos: util.Coordinate{3,2},
			turn: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &rob{
				position:  tt.fields.position,
				direction: tt.fields.direction,
			}

			h.move(tt.turn)

			if !reflect.DeepEqual(h.position, tt.wantPos) {
				t.Errorf("move() got = %v, want %v", h.position, tt.wantPos)
			}
		})
	}
}