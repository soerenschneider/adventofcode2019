package day05

import (
	"fmt"
	"log"
)

// holds information about the number of parameters per opcode
var opcodeParameterCount = map[int]int{
	1: 3,
	2: 3,
	3: 1,
	4: 1,
	5: 2,
	6: 2,
	7: 3,
	8: 3,
	99: 0,
}

// mode defines the mode of an opcode which can be POS and IMM
type mode int
const (
	POS = 0
	IMM = 1
)

type Opcode struct {
	opcode 		int
	modes  		[]mode
	parameters 	int
}

type Interpreter struct {
	alphabet 	[]int
	processed	int
	pointer		int
	shutdown	bool
}

func (i *Interpreter) halt() bool {
	return i.shutdown || i.pointer >= len(i.alphabet) - 1
}

func (i *Interpreter) get(n int, m mode) (int, error) {
	index := n + i.pointer
	if index > len(i.alphabet) - 1 {
		return -1, fmt.Errorf("invalid index %d", n+i.pointer)
	}

	v := i.alphabet[index]

	switch m {
	case POS:
		if v > len(i.alphabet) - 1 {
			return -1, fmt.Errorf("invalid index %d", i.alphabet[v])
		}
		return i.alphabet[v], nil
	case IMM:
		return v, nil
	}

	return 0, fmt.Errorf("invalid mode")
}

func parseOpcode(instruction int) Opcode {
	ret := Opcode{}
	ret.opcode = instruction - (instruction / 100 * 100)
	ret.parameters = opcodeParameterCount[ret.opcode]

	for instruction /= 100; len(ret.modes) < ret.parameters; instruction /= 10 {
		ret.modes = append(ret.modes, mode(instruction%10))
	}

	return ret
}

func (i *Interpreter) execute() {
	for ! i.halt() {
		cmd := parseOpcode(i.alphabet[i.pointer])
		switch cmd.opcode {
		case 1:
			i.c1(cmd)
		case 2:
			i.c2(cmd)
		case 3:
			i.c3(cmd)
		case 4:
			i.c4(cmd)
		case 5:
			i.c5(cmd)
		case 6:
			i.c6(cmd)
		case 7:
			i.c7(cmd)
		case 8:
			i.c8(cmd)
		case 99:
			i.c99(cmd)
		default:
			log.Fatalf("Encountered unknown code: %d", cmd.opcode)
		}

		i.processed++
	}
}

func (i *Interpreter) c1(cmd Opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	c := i.alphabet[i.pointer + cmd.parameters]
	i.alphabet[c] = a + b
	i.pointer += 1 + cmd.parameters
}

func (i *Interpreter) c2(cmd Opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	c := i.alphabet[i.pointer + cmd.parameters]
	i.alphabet[c] = a * b
	i.pointer += 1 + cmd.parameters
}

func (i *Interpreter) c3(cmd Opcode) {
	a := i.alphabet[i.pointer + cmd.parameters]
	var read int
	fmt.Println("Input the number")
	fmt.Scanf("%d", &read)
	i.alphabet[a] = read
	i.pointer += 1 + cmd.parameters
}

func (i *Interpreter) c4(cmd Opcode) {
	a, _ := i.get(1, cmd.modes[0])
	fmt.Println(a)
	i.pointer += 1 + cmd.parameters
}

func (i *Interpreter) c5(cmd Opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	if a != 0 {
		i.pointer = b
	} else {
		i.pointer += 1 + cmd.parameters
	}
}

func (i *Interpreter) c6(cmd Opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	if a == 0 {
		i.pointer = b
	} else {
		i.pointer += 1 + cmd.parameters
	}
}

func (i *Interpreter) c7(cmd Opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	c := i.alphabet[i.pointer + cmd.parameters]
	if a < b {
		i.alphabet[c] = 1
	} else {
		i.alphabet[c] = 0
	}
	i.pointer += 1 + cmd.parameters
}

func (i *Interpreter) c8(cmd Opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	c := i.alphabet[i.pointer + cmd.parameters]
	if a == b {
		i.alphabet[c] = 1
	} else {
		i.alphabet[c] = 0
	}
	i.pointer += 1 + cmd.parameters
}

func (i *Interpreter) c99(cmd Opcode) {
	i.shutdown = true
	i.pointer += 1 + cmd.parameters
}

func Answer05() {
	fmt.Println("--------------")
	fmt.Println("Day 05")
	i := Interpreter{
		alphabet: []int{3,225,1,225,6,6,1100,1,238,225,104,0,1102,31,68,225,1001,13,87,224,1001,224,-118,224,4,224,102,8,223,223,1001,224,7,224,1,223,224,223,1,174,110,224,1001,224,-46,224,4,224,102,8,223,223,101,2,224,224,1,223,224,223,1101,13,60,224,101,-73,224,224,4,224,102,8,223,223,101,6,224,224,1,224,223,223,1101,87,72,225,101,47,84,224,101,-119,224,224,4,224,1002,223,8,223,1001,224,6,224,1,223,224,223,1101,76,31,225,1102,60,43,225,1102,45,31,225,1102,63,9,225,2,170,122,224,1001,224,-486,224,4,224,102,8,223,223,101,2,224,224,1,223,224,223,1102,29,17,224,101,-493,224,224,4,224,102,8,223,223,101,1,224,224,1,223,224,223,1102,52,54,225,1102,27,15,225,102,26,113,224,1001,224,-1560,224,4,224,102,8,223,223,101,7,224,224,1,223,224,223,1002,117,81,224,101,-3645,224,224,4,224,1002,223,8,223,101,6,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,8,226,677,224,102,2,223,223,1005,224,329,1001,223,1,223,1108,677,226,224,102,2,223,223,1006,224,344,101,1,223,223,108,677,226,224,102,2,223,223,1006,224,359,101,1,223,223,7,677,226,224,102,2,223,223,1005,224,374,101,1,223,223,1007,226,677,224,102,2,223,223,1005,224,389,101,1,223,223,8,677,677,224,102,2,223,223,1006,224,404,1001,223,1,223,1007,677,677,224,1002,223,2,223,1006,224,419,101,1,223,223,1108,677,677,224,1002,223,2,223,1005,224,434,1001,223,1,223,1107,226,677,224,102,2,223,223,1005,224,449,101,1,223,223,107,226,226,224,102,2,223,223,1006,224,464,101,1,223,223,1108,226,677,224,1002,223,2,223,1005,224,479,1001,223,1,223,7,677,677,224,102,2,223,223,1006,224,494,1001,223,1,223,1107,677,226,224,102,2,223,223,1005,224,509,101,1,223,223,107,677,677,224,1002,223,2,223,1006,224,524,101,1,223,223,1008,677,677,224,1002,223,2,223,1006,224,539,101,1,223,223,7,226,677,224,1002,223,2,223,1005,224,554,101,1,223,223,108,226,226,224,1002,223,2,223,1006,224,569,101,1,223,223,1008,226,677,224,102,2,223,223,1005,224,584,101,1,223,223,8,677,226,224,1002,223,2,223,1005,224,599,101,1,223,223,1007,226,226,224,1002,223,2,223,1005,224,614,101,1,223,223,1107,226,226,224,1002,223,2,223,1006,224,629,101,1,223,223,107,677,226,224,1002,223,2,223,1005,224,644,1001,223,1,223,1008,226,226,224,1002,223,2,223,1006,224,659,101,1,223,223,108,677,677,224,1002,223,2,223,1005,224,674,1001,223,1,223,4,223,99, 226},
	}
	i.execute()
}