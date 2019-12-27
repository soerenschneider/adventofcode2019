package day10

import (
	"github.com/soerenschneider/adventofcode2019/util"
	"reflect"
	"testing"
)

func TestParseCoordinates(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []Coordinate
	}{
		{
			input: []string{
				".#..#",
				".....",
				"#####",
				"....#",
				"...##",
			},
			want: []Coordinate{
				{0, 1},
				{0, 4},
				{2, 0},
				{2, 1},
				{2, 2},
				{2, 3},
				{2, 4},
				{3, 4},
				{4, 3},
				{4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCoordinates(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadStringLinesFromFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []string
	}{
		{
			filename: "../resources/day10/test.txt",
			want:     []string{".#..#", ".....", "#####", "....#", "...##"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.ReadStringLinesFromFile(tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadStringLinesFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetermineMaxObersableAsteroids(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		input   []Coordinate
		want    Coordinate
		wantCnt int
	}{
		{
			input: []Coordinate{
				{0, 1},
				{0, 4},
				{2, 0},
				{2, 1},
				{2, 2},
				{2, 3},
				{2, 4},
				{3, 4},
				{4, 3},
				{4, 4},
			},
			want:    Coordinate{X: 4, Y: 3},
			wantCnt: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DetermineMaxObservableAsteroids(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetermineMaxObservableAsteroids() coord = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantCnt {
				t.Errorf("DetermineMaxObservableAsteroids() cnt = %v, want %v", got1, tt.wantCnt)
			}
		})
	}
}

func Test_getTargetsWithDistance(t *testing.T) {
	type args struct {
		coordinates []Coordinate
		source      Coordinate
	}
	tests := []struct {
		name string
		args args
		want map[float64]map[int]Coordinate
	}{
		{
			args: args{
				coordinates: []Coordinate{
					{0, 1},
					{0, 3},
					{0, 4},
					{2, 0},
					{2, 1},
					{2, 2},
					{2, 3},
					{2, 4},
				},
				source: Coordinate{0, 1},
			},
			want: map[float64]map[int]Coordinate{
				90:                 map[int]Coordinate{2: {0, 3}, 3: {0, 4}},
				135:                map[int]Coordinate{4: {2, 3}},
				123.69006752597979: map[int]Coordinate{5: {2, 4}},
				153.43494882292202: map[int]Coordinate{3: {2, 2}},
				180:                map[int]Coordinate{2: {2, 1}},
				206.56505117707798: map[int]Coordinate{3: {2, 0}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTargetsWithDistance(tt.args.coordinates, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTargetsWithDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNthVaporizedAsteroid(t *testing.T) {
	type args struct {
		targets   map[float64]map[int]Coordinate
		iteration int
	}
	tests := []struct {
		name    string
		args    args
		want    Coordinate
		wantErr bool
	}{
		{
			name: "Do a loop",
			args: args{
				targets: map[float64]map[int]Coordinate{
					90:                 map[int]Coordinate{2: {0, 3}, 3: {0, 4}},
					135:                map[int]Coordinate{4: {2, 3}},
					123.69006752597979: map[int]Coordinate{5: {2, 4}},
					153.43494882292202: map[int]Coordinate{3: {2, 2}},
					180:                map[int]Coordinate{2: {2, 1}},
					206.56505117707798: map[int]Coordinate{3: {2, 0}},
				},
				iteration: 7,
			},
			want: Coordinate{0, 4},
		},
		{
			name: "First vaporized",
			args: args{
				targets: map[float64]map[int]Coordinate{
					90:                 map[int]Coordinate{2: {0, 3}, 3: {0, 4}},
					135:                map[int]Coordinate{4: {2, 3}},
					123.69006752597979: map[int]Coordinate{5: {2, 4}},
					153.43494882292202: map[int]Coordinate{3: {2, 2}},
					180:                map[int]Coordinate{2: {2, 1}},
					206.56505117707798: map[int]Coordinate{3: {2, 0}},
				},
				iteration: 1,
			},
			want: Coordinate{0, 3},
		},
		{
			name: "Iteration too high",
			args: args{
				targets: map[float64]map[int]Coordinate{
					90:                 map[int]Coordinate{2: {0, 3}, 3: {0, 4}},
					135:                map[int]Coordinate{4: {2, 3}},
					123.69006752597979: map[int]Coordinate{5: {2, 4}},
					153.43494882292202: map[int]Coordinate{3: {2, 2}},
					180:                map[int]Coordinate{2: {2, 1}},
					206.56505117707798: map[int]Coordinate{3: {2, 0}},
				},
				iteration: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNthVaporizedAsteroid(tt.args.targets, tt.args.iteration)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNthVaporizedAsteroid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNthVaporizedAsteroid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAngles(t *testing.T) {
	tests := []struct {
		name    string
		targets map[float64]map[int]Coordinate
		want    []float64
	}{
		{
			targets: map[float64]map[int]Coordinate{
				90:                 map[int]Coordinate{2: {0, 3}, 3: {0, 4}},
				135:                map[int]Coordinate{4: {2, 3}},
				123.69006752597979: map[int]Coordinate{5: {2, 4}},
				153.43494882292202: map[int]Coordinate{3: {2, 2}},
				180:                map[int]Coordinate{2: {2, 1}},
				206.56505117707798: map[int]Coordinate{3: {2, 0}},
			},
			want: []float64{
				90,
				123.69006752597979,
				135,
				153.43494882292202,
				180,
				206.56505117707798,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSortedAngles(tt.targets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSortedAngles() = %v, want %v", got, tt.want)
			}
		})
	}
}
