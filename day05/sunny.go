package day05

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

func Answer05() {
	input := util.ReadInt64Array("resources/day05/input.txt")

	in := make(chan int64, 1)
	out := make(chan int64)
	i := util.NewInterpreter(input, in, out)
	go i.Execute()
	
	in <- 1
	for o := range out {
		fmt.Println(o)
	}
}

func Answer05b() {
	input := util.ReadInt64Array("resources/day05/input.txt")

	in := make(chan int64, 1)
	out := make(chan int64)
	i := util.NewInterpreter(input, in, out)
	go i.Execute()

	in <- 5
	for o := range out {
		fmt.Println(o)
	}
}