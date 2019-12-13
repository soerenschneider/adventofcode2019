package day13

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

func Answer13() {
	mem := util.ReadInt64Array("resources/day13/input.txt")

	in := make(chan int64, 1)
	out := make(chan int64)

	ins := util.NewInterpreter(mem, in, out)

	go ins.Execute()
	count := 0

	for range out {
		<-out
		if <-out == 2 {
			count++
		}
	}

	fmt.Println(count)
}

func playForFree(mem []int64) {
	mem[0] = 2
}

func Answer13b() {
	mem := util.ReadInt64Array("resources/day13/input.txt")
	
	in := make(chan int64, 1)
	out := make(chan int64)

	playForFree(mem)
	ins := util.NewInterpreter(mem, in, out)

	go ins.Execute()

	score := 0
	paddle := 0

	for x := range out {
		y := <-out
		id := <-out

		if showScore(x, y) {
			score = int(id)
		} else if id == 3 {
			paddle = int(x)
		} else if id == 4 {
			input := getJoystickInput(paddle, int(x))
			in <- input
		}
	}

	fmt.Println(score)
}

func showScore(x, y int64) bool {
	return x == -1 && y == 0
}

func getJoystickInput(paddle, x int) int64 {
	if paddle < x {
		return  1
	}
	if paddle > x {
		return  -1
	}

	return  0
}