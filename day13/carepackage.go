package day13

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

// holds information about the number of parameters per opcode
var opcodeParameterCount = map[int]int{1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2, 99: 1}

var operations = map[int]func(*Interpreter, *Opcode) int{
	1:  c1,
	2:  c2,
	3:  c3,
	4:  c4,
	5:  c5,
	6:  c6,
	7:  c7,
	8:  c8,
	9:  c9,
	99: c99,
}

// mode defines the mode of an opcode which can be POS and IMM
type mode int

const (
	POS = 0
	IMM = 1
	REL = 2
)

type Opcode struct {
	opcode     int
	modes      []mode
	parameters int
}

type Interpreter struct {
	alphabet    []int64
	pointer     int
	shutdown    bool
	relPosition int64
	in          chan int64
	out         chan int64
}

func Answer13() {
	mem := util.ReadInt64Array("resources/day13/input.txt")

	in := make(chan int64, 1)
	out := make(chan int64)

	ins := Interpreter{
		alphabet: append([]int64{}, mem...),
		in:       in,
		out:      out,
	}

	go ins.execute()
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

	ins := Interpreter{
		alphabet: append([]int64{}, mem...),
		in:       in,
		out:      out,
	}
	
	playForFree(ins.alphabet)

	go ins.execute()

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

func (i *Interpreter) halt() bool {
	return i.shutdown || i.pointer >= len(i.alphabet)-1
}

func (i *Interpreter) get(n int, opcode Opcode) (addr int64) {
	pos := n + i.pointer
	switch opcode.modes[n-1] {
	case POS:
		addr = i.alphabet[pos]
	case IMM:
		addr = int64(pos)
	case REL:
		addr = i.relPosition + i.alphabet[pos]
	}

	// Grow slice before accessing non-negative, out-of-range address space
	if int64(len(i.alphabet)) <= addr {
		size := addr - int64(len(i.alphabet)) + 1
		newSpace := make([]int64, size)
		i.alphabet = append(i.alphabet, newSpace...)
	}

	return
}

func parseOpcode(instruction int64) Opcode {
	ret := Opcode{}
	ret.opcode = int(instruction % 100)
	ret.parameters = opcodeParameterCount[ret.opcode]

	for instruction /= 100; len(ret.modes) < ret.parameters; instruction /= 10 {
		ret.modes = append(ret.modes, mode(instruction%10))
	}

	return ret
}

func (i *Interpreter) execute() {
	for ! i.halt() {
		cmd := parseOpcode(i.alphabet[i.pointer])
		funky := operations[cmd.opcode]
		inc := funky(i, &cmd)
		i.pointer += inc
	}
}

func c1(i *Interpreter, cmd *Opcode) int {
	op1 := i.alphabet[i.get(1, *cmd)]
	op2 := i.alphabet[i.get(2, *cmd)]
	i.alphabet[i.get(3, *cmd)] = op1 + op2
	return opcodeParameterCount[cmd.opcode]
}

func c2(i *Interpreter, cmd *Opcode) int {
	op1 := i.alphabet[i.get(1, *cmd)]
	op2 := i.alphabet[i.get(2, *cmd)]
	i.alphabet[i.get(3, *cmd)] = op1 * op2
	return opcodeParameterCount[cmd.opcode]
}

func c3(i *Interpreter, cmd *Opcode) int {
	index := i.get(1, *cmd)
	read := <-i.in
	i.alphabet[index] = read
	return opcodeParameterCount[cmd.opcode]
}

func c4(i *Interpreter, cmd *Opcode) int {
	index := i.get(1, *cmd)
	i.out <- i.alphabet[index]
	return opcodeParameterCount[cmd.opcode]
}

func c5(i *Interpreter, cmd *Opcode) int {
	index := i.get(1, *cmd)
	if i.alphabet[index] != 0 {
		val := i.alphabet[i.get(2, *cmd)]
		i.pointer = int(val)
		return 0
	}
	return opcodeParameterCount[cmd.opcode]
}

func c6(i *Interpreter, cmd *Opcode) int {
	index := i.get(1, *cmd)
	if i.alphabet[index] == 0 {
		i.pointer = int(i.alphabet[i.get(2, *cmd)])
		return 0
	}
	return opcodeParameterCount[cmd.opcode]
}

func c7(i *Interpreter, cmd *Opcode) int {
	index1 := i.get(1, *cmd)
	index2 := i.get(2, *cmd)
	if i.alphabet[index1] < i.alphabet[index2] {
		i.alphabet[i.get(3, *cmd)] = 1
	} else {
		i.alphabet[i.get(3, *cmd)] = 0
	}
	return opcodeParameterCount[cmd.opcode]
}

func c8(i *Interpreter, cmd *Opcode) int {
	index1 := i.get(1, *cmd)
	index2 := i.get(2, *cmd)
	dest := i.get(3, *cmd)
	if i.alphabet[index1] == i.alphabet[index2] {
		i.alphabet[dest] = 1
	} else {
		i.alphabet[dest] = 0
	}
	return opcodeParameterCount[cmd.opcode]
}

func c9(i *Interpreter, cmd *Opcode) int {
	i.relPosition += i.alphabet[i.get(1, *cmd)]
	return opcodeParameterCount[cmd.opcode]
}

func c99(i *Interpreter, cmd *Opcode) int {
	close(i.out)
	close(i.in)
	i.shutdown = true
	return opcodeParameterCount[cmd.opcode]
}
