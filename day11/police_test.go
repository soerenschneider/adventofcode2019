package day11

import (
	"image"
	"reflect"
	"testing"
)

func Test_hull_move(t *testing.T) {
	type fields struct {
		direction int
		position image.Point
	}
	tests := []struct {
		name   	string
		fields 	fields
		turn 	int64
		wantPos image.Point
		wantDir int
	}{
		{
			fields: fields{
				position: image.Point{2,2},
			},
			wantPos: image.Point{1,2},
			turn: 0,
		},
		{
			fields: fields{
				position: image.Point{1,2},
				direction: 1,
			},
			wantPos: image.Point{1,3},
			turn: 0,
		},
		{
			fields: fields{
				position: image.Point{1,3},
				direction: 3,
			},
			wantPos: image.Point{1,2},
			turn: 0,
		},
		{
			fields: fields{
				position: image.Point{1,3},
				direction: 3,
			},
			wantPos: image.Point{1,4},
			turn: 1,
		},
		{
			fields: fields{
				position: image.Point{2,2},
			},
			wantPos: image.Point{3,2},
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