package day09

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

func sunny(input int64) {
	alphabet := util.ReadInt64Array("resources/day09/input.txt")

	in := make(chan int64, 1)
	out := make(chan int64)
	i := util.NewInterpreter(alphabet, in, out)

	go i.Execute()
	in <- input

	for o := range out {
		fmt.Println(o)
	}
}

func Answer09() {
	sunny(1)
}

func Answer09b() {
	sunny(2)
}