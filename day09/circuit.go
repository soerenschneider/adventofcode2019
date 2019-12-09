package day09

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// holds information about the number of parameters per opcode
var opcodeParameterCount = map[int]int{
	1: 4,
	2: 4,
	3: 2,
	4: 2,
	5: 3,
	6: 3,
	7: 4,
	8: 4,
	9: 2,
	99: 1,
}

var operations = map[int]func(*Interpreter, *Opcode) int {
	1: c1,
	2: c2,
	3: c3,
	4: c4,
	5: c5,
	6: c6,
	7: c7,
	8: c8,
	9: c9,
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
	opcode 		int
	modes  		[]mode
	parameters 	int
}

type Interpreter struct {
	alphabet    []int64
	pointer     int
	shutdown    bool
	relPosition int64
}

func (i *Interpreter) halt() bool {
	return i.shutdown || i.pointer >= len(i.alphabet) - 1
}

func (i *Interpreter) get(n int, opcode Opcode) (addr int64) {
	switch opcode.modes[n-1] {
	case POS:
		addr = i.alphabet[n + i.pointer]
	case IMM:
		addr = int64(n + i.pointer)
	case REL:
		addr = i.relPosition + i.alphabet[n + i.pointer]
	}

	// Grow slice before accessing non-negative, out-of-range address space
	if int64(len(i.alphabet)) <= addr {
		size := addr - int64(len(i.alphabet))+1
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

func Answer09() {
	fmt.Println("--------------")
	fmt.Println("Day 09")
	input, _ := ioutil.ReadFile("resources/day09/input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	
	alphabet := make([]int64, len(split))
	for i, s := range split {
		x, _ := strconv.Atoi(s)
		alphabet[i] = int64(x)
	}
	
	i := Interpreter{
		alphabet: alphabet,
	}
	
	i.execute()
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
	fmt.Println("Enter input")
	index := i.get(1, *cmd)
	fmt.Scan(&i.alphabet[index])
	return opcodeParameterCount[cmd.opcode]
}

func c4(i *Interpreter, cmd *Opcode) int {
	index := i.get(1, *cmd)
	fmt.Println(i.alphabet[index])
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
	i.shutdown = true
	return opcodeParameterCount[cmd.opcode]
}