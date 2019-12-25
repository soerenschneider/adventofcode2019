package day25

import (
	"bufio"
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"os"
	"strings"
)

const done = '\n'

func readInput(in chan int64) {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		for _, r := range reader.Bytes() {
			in <- int64(r)
		}
		in <- int64(done)
	}
}

func Answer25() int64 {
	input := util.ReadInt64Array("resources/day25/input.txt")
	
	in := make(chan int64)
	out := make(chan int64)
	interpreter := util.NewInterpreter(input, in, out)
	
	go interpreter.Execute()
	go readInput(in)

	var buffer strings.Builder
	for o := range out {
		if o == done {
			fmt.Println(buffer.String())
			buffer.Reset()
		} else if util.IsAscii(o) {
			buffer.WriteByte(byte(o))
		} else {
			fmt.Println(buffer.String())
			return o
		}
	}
	return -1
}

