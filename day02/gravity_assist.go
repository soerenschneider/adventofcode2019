package day02

import (
	"errors"
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

const (
	WordLength = 4
	InstrAdd = 1
	InstrMult = 2
	InstrHalt = 99
)

type Operation struct {
	instruction int
	param1      int
	param2      int
	param3      int
}

func (o *Operation) halt() bool {
	return o.instruction == InstrHalt
}

func (o *Operation) apply(input []int) error {
	if o.param1 >= len(input) || o.param2 >= len(input) {
		return errors.New("incorrect sequence for input")
	}
	
	if o.instruction == InstrAdd {
		input[o.param3] = input[o.param1] + input[o.param2]
		return nil
	} else if o.instruction == InstrMult {
		input[o.param3] = input[o.param1] * input[o.param2]
		return nil
	}

	return fmt.Errorf("invalid instruction encountered: %d", o.instruction)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func extractOperation(input []int, index int) Operation {
	end := Min(index + WordLength, len(input))

	op := Operation{}
	for i := 0; i < WordLength; i++ {
		position := i + index
		if position < end {
			switch i {
			case 0:
				op.instruction = input[position]
			case 1:
				op.param1 = input[position]
			case 2:
				op.param2 = input[position]
			default:
				op.param3 = input[position]
			}
		}
	}
	return op
}

func ProcessInput(input []int, noun int, verb int) ([]int, error) {
	if input == nil || len(input) < 3 {
		return nil, fmt.Errorf("input either nil or |input| < 3")
	}
	input[1] = noun
	input[2] = verb

	for index := 0; index < len(input); index += WordLength {
		operation := extractOperation(input, index)
		if operation.halt() {
			break
		}
		operation.apply(input)
	}

	return input, nil
}

func Answer() {
	origInput := util.ReadIntArray("resources/day02/input.txt")
	input := append([]int(nil), origInput...)
	output, _ := ProcessInput(input, 12, 2)
	answer := output[0]

	fmt.Printf("Answer 02: %d\n", answer)
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			input := append([]int(nil), origInput...)
			output, _ := ProcessInput(input, noun, verb)
			if output[0] == 19690720 {
				fmt.Printf("Answer 02b: 100 * %d + %d = %d\n", noun, verb, 100 * noun + verb)
			}
		}
	}
}
