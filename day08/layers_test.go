package day08

import (
	"reflect"
	"testing"
)

func TestGetLayers(t *testing.T) {
	type args struct {
		x    int
		y    int
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    *Image
		wantErr bool
	}{
		{
			name:    "Invalid dimensions",
			args:    args{5, 4, "1"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid dimensions",
			args: args{3, 2, "123456789012"},
			want: &Image{
				layers: []ImageLayer{
					{data: [][]int{
						{1, 2, 3},
						{4, 5, 6}},
					},
					{data: [][]int{
						{7, 8, 9},
						{0, 1, 2},
					}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseImage(tt.args.data, tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseImage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImageLayer_Count(t *testing.T) {
	type fields struct {
		data [][]int
	}
	tests := []struct {
		name   string
		fields fields
		digit  int
		want   int
	}{
		{
			name: "Nada",
			fields: fields{
				data: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			digit: 0,
			want:  0,
		},
		{
			name: "Found one",
			fields: fields{
				data: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			digit: 6,
			want:  1,
		},
		{
			name: "Found two",
			fields: fields{
				data: [][]int{
					{6, 2, 3},
					{4, 5, 6},
				},
			},
			digit: 6,
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ImageLayer{
				data: tt.fields.data,
			}
			if got := i.Count(tt.digit); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getArray(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    []int
		wantErr bool
	}{
		{
			name:    "Simple",
			args:    "1234",
			want:    []int{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "Empty",
			args:    "abcd",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Empty",
			args:    "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getArray(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("getArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_GetLayerWithFewestDigits(t *testing.T) {
	type fields struct {
		layers []ImageLayer
	}
	tests := []struct {
		name   string
		fields fields
		args   int
		want   *ImageLayer
	}{
		{
			fields: fields{
				layers: []ImageLayer{
					{data: [][]int{
						{1, 2, 3},
						{4, 5, 6},
					},
					},
					{data: [][]int{
						{7, 8, 9},
						{0, 1, 2},
					},
					},
				},
			},
			args: 0,
			want: &ImageLayer{
				data: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				layers: tt.fields.layers,
			}
			if got := i.GetLayerWithFewestDigits(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLayerWithFewestDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_render(t *testing.T) {
	type fields struct {
		layers []ImageLayer
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]string
	}{
		{
			fields: fields{
				layers: []ImageLayer{
					{data: [][]int{
						{0, 2},
						{2, 2}},
					},
					{data: [][]int{
						{1, 1},
						{2, 2}},
					},
					{data: [][]int{
						{2, 2},
						{1, 2}},
					},
					{data: [][]int{
						{0, 0},
						{0, 0}},
					},
				},
			},
			want: [][]string{
				{renderRepresentation[0], renderRepresentation[1]},
				{renderRepresentation[1], renderRepresentation[0]},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				layers: tt.fields.layers,
			}
			if got := i.Render(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderRepresentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_getDominantPixel(t *testing.T) {
	type fields struct {
		layers []ImageLayer
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			args:    args{0, 0},
			want:    0,
			wantErr: false,
		},
		{
			args:    args{0, 1},
			want:    1,
			wantErr: false,
		},
		{
			args:    args{1, 0},
			want:    1,
			wantErr: false,
		},
		{
			args:    args{1, 1},
			want:    0,
			wantErr: false,
		},
	}
	layers := []ImageLayer{
		{data: [][]int{
			{0, 2},
			{2, 2}},
		},
		{data: [][]int{
			{1, 1},
			{2, 2}},
		},
		{data: [][]int{
			{2, 2},
			{1, 2}},
		},
		{data: [][]int{
			{0, 0},
			{0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				layers: layers,
			}
			got, err := i.getDominantPixel(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDominantPixel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getDominantPixel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_getDimensions(t *testing.T) {
	type fields struct {
		layers []ImageLayer
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			fields: fields{
				layers: []ImageLayer{
					{data: [][]int{
						{0, 2},
						{2, 2}},
					},
				},
			},
			want: []int{2, 2},
		},
		{
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				layers: tt.fields.layers,
			}
			if got := i.getDimensions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}
