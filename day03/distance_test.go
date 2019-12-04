package day03

import (
	"reflect"
	"testing"
)

func TestPoint_movement(t *testing.T) {
	tests := []struct {
		name    	string
		point   	Point
		movement    Movement
		want    	Point
		trace 		[]Point
		wantErr 	bool
	}{
		{
			name: "No movement",
			point: Point{},
			movement: Movement{up, 0},
			want: Point{},
			trace: nil,
			wantErr: false,
		},
		{
			name: "One up",
			point: Point{},
			movement: Movement{up, 1},
			want: Point{x:0, y:1},
			trace: []Point{{0, 1}},
			wantErr: false,
		},
		{
			name: "One down",
			point: Point{},
			movement: Movement{down, 1},
			want: Point{x:0, y:-1},
			trace: []Point{{0, -1}},
			wantErr: false,
		},
		{
			name: "Two left",
			point: Point{},
			movement: Movement{left, 2},
			want: Point{x:-2, y:0},
			trace: []Point{{-1, 0}, {-2, 0}},
			wantErr: false,
		},
		{
			name: "Three right",
			point: Point{},
			movement: Movement{right, 3},
			want: Point{x:3, y:0},
			trace: []Point{{1, 0}, {2, 0}, {3,0}},
			wantErr: false,
		},
		{
			name: "10 right, starting at (-5, 0)",
			point: Point{x:-5, y:0},
			movement: Movement{right, 10},
			want: Point{x:5, y:0},
			trace: []Point{{-4, 0},{-3, 0},{-2, 0},{-1, 0},{0, 0},{1, 0},{2, 0},{3, 0},{4, 0},{5, 0}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trace := tt.point.move(tt.movement)
			if !reflect.DeepEqual(tt.point, tt.want) {
				t.Errorf("move() got = %v, want %v", tt.point, tt.want)
			}
			if !reflect.DeepEqual(trace, tt.trace) {
				t.Errorf("move() got = %v, want %v", trace, tt.trace)
			}
		})
	}
}

func TestPoint_ManhattanDistance(t *testing.T) {
	tests := []struct {
		name   string
		point  Point
		want   int
	}{
		{
			name: "Example 1",
			point: Point{3, 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{0,0}
			if got := p.ManhattanDistance(tt.point); got != tt.want {
				t.Errorf("ManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	tests := []struct {
		arg int
		want int
	}{
		{
			arg: 1,
			want: 1,
		},
		{
			arg: 0,
			want: 0,
		},
		{
			arg: -1,
			want: 1,
		},
		{
			arg: -10,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := abs(tt.arg); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIntersectionEmpty(t *testing.T) {
	a := make(map[Point]int)
	b := make(map[Point]int)
	want := make(map[Point][2]int)
	if got := findIntersection(a, b); !reflect.DeepEqual(got, want) {
		t.Errorf("findIntersection() = %v, want %v", got, want)
	}
}

func Test_findIntersectionNegative(t *testing.T) {
	a := make(map[Point]int)
	p := Point{0, 1}
	a[p] = 1
	p = Point{0, 2}
	a[p] = 2
	b := make(map[Point]int)
	want := make(map[Point][2]int)
	if got := findIntersection(a, b); !reflect.DeepEqual(got, want) {
		t.Errorf("findIntersection() = %v, want %v", got, want)
	}
}

func Test_findIntersectionPositive(t *testing.T) {
	a := make(map[Point]int)
	b := make(map[Point]int)
	p := Point{0, 1}
	a[p] = 1
	p = Point{0, 2}
	a[p] = 2
	b[p] = 3
	want := make(map[Point][2]int)
	want[p] = [2]int{2, 3}
	if got := findIntersection(a, b); !reflect.DeepEqual(got, want) {
		t.Errorf("findIntersection() = %v, want %v", got, want)
	}
}

func Test_getMovement(t *testing.T) {
	tests := []struct {
		name    string
		instruction string
		want    Movement
		wantErr bool
	}{
		{
			name: "U4",
			instruction: "U4",
			want: Movement{dir: up, arg: 4},
			wantErr: false,
		},
		{
			name: "D10",
			instruction: "D10",
			want: Movement{dir: down, arg: 10},
			wantErr: false,
		},
		{
			name: "L3",
			instruction: "L3",
			want: Movement{dir: left, arg: 3},
			wantErr: false,
		},
		{
			name: "R353",
			instruction: "R353",
			want: Movement{dir: right, arg: 353},
			wantErr: false,
		},
		{
			name: "RZ",
			instruction: "RZ",
			want: Movement{dir: right, arg: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMovement(tt.instruction)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMovement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMovement() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMovements(t *testing.T) {
	tests := []struct {
		name string
		instructions string
		want []Movement
		wantErr bool
	}{
		{
			name: "Positive",
			instructions: "L5,L5,L5",
			want: []Movement{{left, 5}, {left, 5}, {left, 5}},
			wantErr: false,
		},
		{
			name: "Positive",
			instructions: "L5,R12",
			want: []Movement{{left, 5}, {right, 12}},
			wantErr: false,
		},
		{
			name: "Extreme",
			instructions: "",
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMovements(tt.instructions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMovements() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("getMovement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getTrail(t *testing.T) {
	movements := []Movement{}
	movements = append(movements, Movement{left, 1})
	movements = append(movements, Movement{up, 2})
	movements = append(movements, Movement{left, 1})
	movements = append(movements, Movement{up, 1})

	want := make(map[Point]int)

	want[Point{-1, 0}] = 1
	want[Point{-1, 1}] = 2
	want[Point{-1, 2}] = 3
	want[Point{-2, 2}] = 4
	want[Point{-2, 3}] = 5


	if got := getTrail(movements); !reflect.DeepEqual(got, want) {
		t.Errorf("getTrail() = %v, want %v", got, want)
	}
}

func Test_getTrailEmpty(t *testing.T) {
	movements := []Movement{}
	want := make(map[Point]int)

	if got := getTrail(movements); !reflect.DeepEqual(got, want) {
		t.Errorf("getTrail() = %v, want %v", got, want)
	}
}

func Test_findMinCost(t *testing.T) {
		points :=  make(map[Point][2]int)

		min := Point{666, 666}

		points[min] = [2]int{2, 5}
		points[Point{1, 2}] = [2]int{ 5, 50 }

		want :=  min
		want1 := 7
		t.Run("", func(t *testing.T) {
			got, got1 := findMinCost(points)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("findMinCost() got = %v, want %v", got, want)
			}
			if got1 != want1 {
				t.Errorf("findMinCost() got1 = %v, want %v", got1, want1)
			}
		})
}

func Test_findMinDist(t *testing.T) {
	points :=  make(map[Point][2]int)

	min := Point{10, -25}

	points[min] = [2]int{100, 500}
	points[Point{11, 26}] = [2]int{ 5, 50 }

	want :=  min
	want1 := 35
	t.Run("", func(t *testing.T) {
		got, got1 := findMinDist(points)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("findMinCost() got = %v, want %v", got, want)
		}
		if got1 != want1 {
			t.Errorf("findMinCost() got1 = %v, want %v", got1, want1)
		}
	})
}

func Test_solveDist(t *testing.T) {
	input := []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	routes := make([]map[Point]int, 2, 2)
	for i, _ := range input {
		movements, _ := getMovements(input[i])
		routes[i] = getTrail(movements)
	}

	intersections := findIntersection(routes[0], routes[1])
	_, minDist := findMinDist(intersections)
	if minDist != 159 {
		t.Errorf("Expected minDist to be %d", minDist)
	}
}

func Test_solveDist2(t *testing.T) {
	input := []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	routes := make([]map[Point]int, 2, 2)
	for i, _ := range input {
		movements, _ := getMovements(input[i])
		routes[i] = getTrail(movements)
	}

	intersections := findIntersection(routes[0], routes[1])
	_, minDist := findMinDist(intersections)
	if minDist != 135 {
		t.Errorf("Expected minDist to be %d", minDist)
	}
}

func Test_solveCost1(t *testing.T) {
	input := []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	routes := make([]map[Point]int, 2, 2)
	for i, _ := range input {
		movements, _ := getMovements(input[i])
		routes[i] = getTrail(movements)
	}

	intersections := findIntersection(routes[0], routes[1])
	_, minCost := findMinCost(intersections)
	if minCost != 610 {
		t.Errorf("Expected minDist to be %d", minCost)
	}
}

func Test_solveCost2(t *testing.T) {
	input := []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	routes := make([]map[Point]int, 2, 2)
	for i, _ := range input {
		movements, _ := getMovements(input[i])
		routes[i] = getTrail(movements)
	}

	intersections := findIntersection(routes[0], routes[1])
	_, minCost := findMinCost(intersections)
	if minCost != 410 {
		t.Errorf("Expected minDist to be %d", minCost)
	}
}