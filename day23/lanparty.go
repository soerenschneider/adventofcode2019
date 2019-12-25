package day23

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

const (
	computers      = 50
	idlePacket     = -1
	desiredAddress = 255
)

func copyInput(input []int64) []int64 {
	copy := make([]int64, len(input))
	for i, v := range input {
		copy[i] = v
	}
	return copy
}

func bootComputers(input []int64, in, out []chan int64) {
	for i := 0; i < computers; i++ {
		copy := copyInput(input)

		in[i] = make(chan int64)
		out[i] = make(chan int64)

		interpreter := util.NewInterpreter(copy, in[i], out[i])
		go interpreter.Execute()

		in[i] <- int64(i)
		in[i] <- int64(-1)
	}
}

func Answer23() {
	input := util.ReadInt64Array("resources/day23/input.txt")

	in := make([]chan int64, computers)
	out := make([]chan int64, computers)

	bootComputers(input, in, out)

	receivePackets(in, out)
}

func receivePackets(in, out []chan int64) {
	idleComputers := 0
	var old, nat [2]int64

	for index := 0; ; index = (index + 1) % computers {
		select {
		case in[index] <- idlePacket:
			idleComputers++

		case destination := <-out[index]:
			if destination == desiredAddress {
				next := [2]int64{
					<-out[index],
					<-out[index],
				}

				if nat == [2]int64{} {
					fmt.Printf("Part a: %d\n", next[1])
				}

				nat = next
			} else {
				in[destination] <- <-out[index]
				in[destination] <- <-out[index]
			}
			idleComputers = 0
		}

		if idleComputers >= computers {
			if nat[1] == old[1] {
				fmt.Printf("Part b: %d\n", nat[1])
				return
			}

			in[0] <- nat[0]
			in[0] <- nat[1]
			old = nat
			idleComputers = 0
		}
	}
}
