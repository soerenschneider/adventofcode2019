package day07

import (
	"fmt"
	"log"
	"sync"
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

var wg sync.WaitGroup

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
	input		chan int
	output		chan int
}

func Phases(input []int, alphabet []int) int {
	permutations := Permutate(input)
	max := 0
	for _, p := range permutations {
		cost := TryPermutation(alphabet, p)
		if cost > max {
			max = cost
		}
	}
	return max
}

func TryPermutation(alphabet []int, permutation []int) int {
	var chans []chan int
	for range permutation {
		chans = append(chans, make(chan int, 1))
	}

	for phase, perm := range permutation {
		i := Interpreter{
			alphabet: append([]int(nil), alphabet...),
			input: chans[phase],
			output: chans[(phase+1)%len(permutation)],
		}
		wg.Add(1)
		go i.execute()
		chans[phase] <- perm
	}
	chans[0] <- 0

	wg.Wait()
	return <- chans[0]
}

func Permutate(arr []int)[][]int{
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, depth int){
		if depth == 1{
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < depth; i++{
				helper(arr, depth- 1)
				if depth% 2 == 1{
					tmp := arr[i]
					arr[i] = arr[depth- 1]
					arr[depth- 1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[depth- 1]
					arr[depth- 1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
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
			wg.Done()
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
	read := <- i.input
	i.alphabet[a] = read
	i.pointer += 1 + cmd.parameters
}

func (i *Interpreter) c4(cmd Opcode) {
	a, _ := i.get(1, cmd.modes[0])
	i.output <- a
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