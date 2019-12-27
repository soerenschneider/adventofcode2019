package day07

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
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

type opcode struct {
	opcode 		int
	modes  		[]mode
	parameters 	int
}

type interpreter struct {
	alphabet 	[]int
	processed	int
	pointer		int
	shutdown	bool
	input		chan int
	output		chan int
	wg 			*sync.WaitGroup
}

func Answer07() {
	alphabet := util.ReadIntArray("resources/day07/input.txt")
	phase := []int{0,1,2,3,4}
	ret := Phases(alphabet, phase)
	fmt.Println(ret)
}

func Answer07b() {
	alphabet := util.ReadIntArray("resources/day07/input.txt")
	phase := []int{5,6,7,8,9}
	ret := Phases(alphabet, phase)
	fmt.Println(ret)
}

func Phases(alphabet []int, phase []int) int {
	permutations := GetPermutations(phase)
	
	max := 0
	for _, permutation := range permutations {
		cost := TryPermutation(alphabet, permutation)
		if cost > max {
			max = cost
		}
	}

	return max
}

func TryPermutation(alphabet []int, permutation []int) int {
	wg := &sync.WaitGroup{}
	var chans []chan int

	for range permutation {
		chans = append(chans, make(chan int, 1))
	}

	for phase, perm := range permutation {
		i := interpreter{
			alphabet: append([]int(nil), alphabet...),
			input: chans[phase],
			output: chans[(phase+1)%len(permutation)],
			wg: wg,
		}

		wg.Add(1)
		go i.execute()
		chans[phase] <- perm
	}
	chans[0] <- 0

	wg.Wait()
	return <- chans[0]
}

func GetPermutations(arr []int)[][]int{
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

func (i *interpreter) halt() bool {
	return i.shutdown || i.pointer >= len(i.alphabet) - 1
}

func (i *interpreter) get(n int, m mode) (int, error) {
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

func parseOpcode(instruction int) opcode {
	ret := opcode{}
	ret.opcode = instruction - (instruction / 100 * 100)
	ret.parameters = opcodeParameterCount[ret.opcode]

	for instruction /= 100; len(ret.modes) < ret.parameters; instruction /= 10 {
		ret.modes = append(ret.modes, mode(instruction%10))
	}

	return ret
}

func (i *interpreter) execute() {
	defer i.wg.Done()

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

func (i *interpreter) c1(cmd opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	c := i.alphabet[i.pointer + cmd.parameters]
	i.alphabet[c] = a + b
	i.pointer += 1 + cmd.parameters
}

func (i *interpreter) c2(cmd opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	c := i.alphabet[i.pointer + cmd.parameters]
	i.alphabet[c] = a * b
	i.pointer += 1 + cmd.parameters
}

func (i *interpreter) c3(cmd opcode) {
	a := i.alphabet[i.pointer + cmd.parameters]
	read := <- i.input
	i.alphabet[a] = read
	i.pointer += 1 + cmd.parameters
}

func (i *interpreter) c4(cmd opcode) {
	a, _ := i.get(1, cmd.modes[0])
	i.output <- a
	i.pointer += 1 + cmd.parameters
}

func (i *interpreter) c5(cmd opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	if a != 0 {
		i.pointer = b
	} else {
		i.pointer += 1 + cmd.parameters
	}
}

func (i *interpreter) c6(cmd opcode) {
	a, _ := i.get(1, cmd.modes[0])
	b, _ := i.get(2, cmd.modes[1])
	if a == 0 {
		i.pointer = b
	} else {
		i.pointer += 1 + cmd.parameters
	}
}

func (i *interpreter) c7(cmd opcode) {
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

func (i *interpreter) c8(cmd opcode) {
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

func (i *interpreter) c99(cmd opcode) {
	i.shutdown = true
	i.pointer += 1 + cmd.parameters
}