package day19

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

const (
	stationary = 0
	pulled     = 1
)

func getDroneReport(input []int64, x, y int) int64 {
	in := make(chan int64, 2)
	out := make(chan int64)

	interpreter := util.NewInterpreter(input, in, out)
	go interpreter.Execute()

	in <- int64(x)
	in <- int64(y)

	return <-out
}

func Answer19() {
	input := util.ReadInt64Array("resources/day19/input.txt")

	gridSize := 50
	var pointsAffected int

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if getDroneReport(input, x, y) == pulled {
				pointsAffected++
			}
		}
	}

	fmt.Printf("%d points affected\n", pointsAffected)
}

func Answer19b() {
	input := util.ReadInt64Array("resources/day19/input.txt")

	squareSize := 100
	actual := squareSize - 1

	x := 0
	y := actual

	var point *util.Coordinate

	// Slowly move down diagonally from the origin and check whether
	// a points spans a squareSize by checking its edges.
	for nil == point {
		for getDroneReport(input, x, y) == stationary {
			x++
		}

		if checkEdges(input, x, y, squareSize) {
			point = &util.Coordinate{X: x, Y: y}
		}
		y++
	}

	res := point.X*10000 + (point.Y - actual)
	fmt.Println(res)
}

func checkEdges(input []int64, x, y int, size int) bool {
	actual := size - 1
	return getDroneReport(input, x, y) == pulled &&
		getDroneReport(input, x+(actual), y) == pulled &&
		getDroneReport(input, x, y-actual) == pulled &&
		getDroneReport(input, x+(actual), y-(actual)) == pulled
}
