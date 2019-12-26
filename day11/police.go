package day11

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

var directions = []util.Coordinate{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
var output = map[int64]string{1: "##", 0: "  "}

type rob struct {
	position  util.Coordinate
	direction int
}

func (r *rob) move(turn int64) {
	r.direction = (r.direction + 2 * int(turn) + 1) % 4
	r.position = r.position.Move(directions[r.direction])
}

func PaintHull(in chan int64, out chan int64, initHull bool) map[util.Coordinate]int64 {
	rob := rob{}
	hull := map[util.Coordinate]int64{}

	if initHull {
		hull[rob.position] = 1
	}

	for {
		in <- hull[rob.position]
		hull[rob.position] = <-out
		turn, ok := <-out
		if !ok {
			break
		}
		rob.move(turn)
	}
	
	return hull
}

func PrintMessage(hull map[util.Coordinate]int64) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			tile := output[hull[util.Coordinate{X: x, Y: y}]]
			fmt.Print(tile)
		}
		fmt.Println()
	}
}

func Answer11() {
	mem := util.ReadInt64Array("resources/day11/input.txt")

	in := make(chan int64, 1)
	out := make(chan int64)

	ins := util.NewInterpreter(append([]int64{}, mem...), in, out)
	
	go ins.Execute()
	hull := PaintHull(in, out, false)
	fmt.Println(len(hull))
}

func Answer11b() {
	mem := util.ReadInt64Array("resources/day11/input.txt")

	in := make(chan int64, 1)
	out := make(chan int64)

	ins := util.NewInterpreter(append([]int64{}, mem...), in, out)

	go ins.Execute()
	hull := PaintHull(in, out, true)
	PrintMessage(hull)
}