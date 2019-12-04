package day03

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	up    = "U"
	down  = "D"
	left  = "L"
	right = "R"
)

type Point struct {
	x int
	y int
}

type Movement struct {
	dir string
	arg int
}

// move performs a Movement on a point.
func (p *Point) move(movement Movement) []Point {
	var trail []Point

	for n := 0; n < abs(movement.arg); n++ {
		switch movement.dir {
		case up:
			p.y += 1
		case down:
			p.y -= 1
		case left:
			p.x -= 1
		case right:
			p.x += 1
		}
		point := Point{p.x, p.y}
		trail = append(trail, point)
	}
	
	return trail
}

// ManhattanDistance calculates the Manhattan Distance to another point.
func (p *Point) ManhattanDistance(q Point) int {
	return abs(q.x - p.x) + abs(q.y - p.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMovement(instruction string) (Movement, error) {
	if len(instruction) < 2 {
		return Movement{}, fmt.Errorf("invalid instruction %s", instruction)
	}
	direction := instruction[0:1]
	steps, err := strconv.Atoi(instruction[1:])
	return Movement{dir: direction, arg: steps}, err
}

func getMovements(instructions string) ([]Movement, error) {
	var movements []Movement

	tokenized := strings.Split(instructions, ",")
	for _, instruction := range tokenized {
		movement, err := getMovement(instruction); if err != nil {
			return nil, err
		}
		movements = append(movements, movement)
	}
	
	return movements, nil
}

func getTrail(movements []Movement) map[Point]int {
	p := Point{}
	var cost int = 0
	route := make(map[Point]int)

	for _, movement := range movements {
		trail := p.move(movement)
		for _, point := range trail {
			cost++
			if _, ok := route[point]; ok {
				continue
			} else {
				route[point] = cost
			}
		}
	}

	return route
}

func findIntersection(a, b map[Point]int) map[Point][2]int {
	joint := make(map[Point][2]int)

	for intersection, costA := range a {
		if costB, ok := b[intersection]; ok {
			joint[intersection] = [2]int{costA, costB}
		}
	}
	return joint
}

func findMinCost(points map[Point][2]int) (Point, int) {
	var minPoint Point
	minCost := math.MaxInt32

	for point, costs := range points {
		costs := costs[0] + costs[1]
		if costs < minCost {
			minPoint = point
			minCost = costs
		}
	}

	return minPoint, minCost
}

func findMinDist(points map[Point][2]int) (Point, int) {
	minDist := math.MaxInt32
	var minPoint Point

	center := Point{}
	for point, _ := range points {
		dist := center.ManhattanDistance(point)
		if dist < minDist {
			minPoint = point
			minDist = dist
		}
	}

	return minPoint, minDist
}

func Solve(input []string) {
	routes := make([]map[Point]int, 2, 2)
	for i, _ := range input {
		movements, _ := getMovements(input[i])
		routes[i] = getTrail(movements)
	}

	intersections := findIntersection(routes[0], routes[1])

	intersection, minDist := findMinDist(intersections)
	fmt.Printf("(%d,%d): %d cost\n", intersection.x, intersection.y, minDist)

	intersection, minCost := findMinCost(intersections)
	fmt.Printf("(%d,%d): %d cost\n", intersection.x, intersection.y, minCost)
}

func Answer03() {
	p, err :=  filepath.Abs("resources/day03/input.txt"); if err != nil {
		log.Fatal("Couldnt open file")
	}

	content, _ := ioutil.ReadFile(p)
	input := strings.Split(string(content), "\n")
	Solve(input)
}