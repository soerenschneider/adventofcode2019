package day17

import (
	"errors"
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"strconv"
	"strings"
)

const (
	scaffold      = "#"
	empty         = "."
	newline       = "\n"
	posNorth      = "^"
	posWest       = "<"
	posSouth      = "v"
	posEast       = ">"
)

var (
	North = util.Coordinate{X: 0, Y: -1}
	South = util.Coordinate{X: 0, Y: 1}
	West  = util.Coordinate{X: -1, Y: 0}
	East  = util.Coordinate{X: 1, Y: 0}

	turnLeft = map[util.Coordinate]util.Coordinate{
		North: West,
		West:  South,
		South: East,
		East:  North}

	turnRight = map[util.Coordinate]util.Coordinate{
		North: East,
		East:  South,
		South: West,
		West:  North}
)

func interpretInput(input int) string {
	return string(rune(input))
}

type bot struct {
	field      [][]string
	currentRow int
	direction  util.Coordinate
	position   util.Coordinate
}

func newBot() *bot {
	return &bot{
		field: [][]string{[]string{}},
	}
}

func wakeUpRobot(input []int64) {
	input[0] = 2
}

func (b *bot) discoverField(response int64) {
	character := interpretInput(int(response))
	if character == newline {
		b.field = append(b.field, []string{})
		b.currentRow++
	} else {
		b.field[b.currentRow] = append(b.field[b.currentRow], character)
	}
}

func (b *bot) SumOfIntersections() int {
	sum := 0
	for y, row := range b.field {
		if y < 1 || y > len(b.field)-2 {
			continue
		}
		for x, char := range row {
			if x < 1 || x > len(row)-2 {
				continue
			}

			if char != scaffold {
				continue
			}

			if len(b.field[y+1]) <= x {
				continue
			}

			if isIntersection(x, y, b.field) {
				sum += y * x
			}
		}
	}

	return sum
}

func isIntersection(x int, y int, field [][]string) bool {
	if y < 1 || y > len(field)-2 {
		return false
	}
	if x < 1 || x > len(field[y])-2 {
		return false
	}

	return field[y-1][x] == scaffold &&
		field[y+1][x] == scaffold &&
		field[y][x-1] == scaffold &&
		field[y][x+1] == scaffold
}

func isScaffold(position util.Coordinate, field [][]string) bool {
	return position.X >= 0 &&
		position.Y >= 0 &&
		position.Y < len(field) &&
		position.X < len(field[position.Y]) &&
		field[position.Y][position.X] == scaffold
}

func Answer17() {
	input := util.ReadInt64Array("resources/day17/input.txt")
	in := make(chan int64)
	out := make(chan int64, 1)

	interpreter := util.NewInterpreter(input, in, out)
	go interpreter.Execute()
	bot := newBot()

	for output := range out {
		bot.discoverField(output)
	}

	fmt.Println(bot.SumOfIntersections())
}

func Answer17b() {
	input := util.ReadInt64Array("resources/day17/input.txt")
	in := make(chan int64)
	out := make(chan int64, 1)

	interpreter := util.NewInterpreter(input, in, out)
	go interpreter.Execute()
	bot := newBot()

	for output := range out {
		bot.discoverField(output)
	}

	ret, _ := determinePositionAndDirection(bot.field)
	bot.position = ret[0]
	bot.direction = ret[1]

	path := bot.discoverPath()
	compacted := asAsciiStream(path)
	bot.PrintField()

	in = make(chan int64, 100)
	out = make(chan int64)

	wakeUpRobot(input)

	interpreter = util.NewInterpreter(input, in, out)
	go interpreter.Execute()

	for _, c := range compacted {
		in <- c
	}

	for output := range out {
		if !util.IsAscii(output) {
			fmt.Println(output)
		}
	}
}

// determinePositionAndDirection finds out the position and the direction of the robot. 
// The first value in the array is the position, the second coordinate is the direction. 
// If no robot is found, an error and an empty coordinate slice is returned. 
func determinePositionAndDirection(field [][]string) ([]util.Coordinate, error) {
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			switch field[y][x] {
			case posNorth:
				return []util.Coordinate{
					{x, y},
					North,
				}, nil
			case posSouth:
				return []util.Coordinate{
					{x, y},
					South,
				}, nil
			case posWest:
				return []util.Coordinate{
					{x, y},
					West,
				}, nil
			case posEast:
				return []util.Coordinate{
					{x, y},
					East,
				}, nil
			}
		}
	}

	return nil, errors.New("no position found")
}

func (b *bot) discoverPath() []string {
	var path []string

	quit := false
	for !quit {
		steps := 0

		// move in the current direction as far as it goes
		for isScaffold(b.position.Move(b.direction), b.field) {
			b.position = b.position.Move(b.direction)
			steps++
		}
		// append the amount of steps to the current path
		if steps > 0 {
			path = append(path, strconv.Itoa(steps))
		}

		if isScaffold(b.position.Move(turnLeft[b.direction]), b.field) {
			b.direction = turnLeft[b.direction]
			path = append(path, "L")
		} else if isScaffold(b.position.Move(turnRight[b.direction]), b.field) {
			b.direction = turnRight[b.direction]
			path = append(path, "R")
		} else {
			quit = true
		}
	}

	return path
}

func (b *bot) PrintField() {
	for y := 0; y < len(b.field); y++ {
		for x := 0; x < len(b.field[y]); x++ {
			fmt.Printf("%s ", b.field[y][x])
		}
		fmt.Println()
	}
}

func asAsciiStream(path []string) []int64 {
	result := compressPath(path, [][]string{path}, nil)
	functions := result[0]

	main := strings.Join(functions[0], ",")
	a := strings.Join(functions[1], ",")
	b := strings.Join(functions[2], ",")
	c := strings.Join(functions[3], ",")

	var ret []int64
	for _, rune := range fmt.Sprintf("%s\n%s\n%s\n%s\nn\n", main, a, b, c) {
		ret = append(ret, int64(rune))
	}

	return ret
}

// Source: https://github.com/GreenLightning/aoc19/blob/40ba3bc2d9e874dd1044748682a5488f208a4f47/day17/main.go
func compressPath(path []string, fragments [][]string, functions [][]string) (result [][4][]string) {
	if len(functions) == 3 {
		// The main function cannot call movement commands, so there must not be any commands left.
		if len(fragments) != 0 {
			return nil
		}

		// Replace path with function calls to compute main function.
		var mainFunction []string
		for len(path) != 0 {
			for i, function := range functions {
				if hasPrefix(path, function) {
					mainFunction = append(mainFunction, string('A'+i))
					path = path[len(function):]
				}
			}
		}

		// Check memory limit for main function.
		if len(strings.Join(mainFunction, ",")) > 20 {
			return nil
		}

		var program [4][]string
		program[0] = mainFunction
		program[1] = functions[0]
		program[2] = functions[1]
		program[3] = functions[2]

		result = append(result, program)
		return
	}

	if len(fragments) == 0 {
		// Add empty candidate to functions.
		newFunctions := make([][]string, 0, 3)
		newFunctions = append(newFunctions, functions...)
		newFunctions = append(newFunctions, []string{})

		subresult := compressPath(path, fragments, newFunctions)
		result = append(result, subresult...)
		return
	}

	// Checking the first fragment is enough.
	fragment := fragments[0]

	// Collect candidates.
	var candidates [][]string
	for length := 1; length <= len(fragment); length++ {
		candidate := fragment[:length]
		text := strings.Join(candidate, ",")
		if len(text) <= 20 {
			candidates = append(candidates, candidate)
		}
	}

	// Try each candidate.
	for _, candidate := range candidates {
		// Split fragments by candidate.
		var newFragments [][]string
		for _, fragment := range fragments {
			for {
				i := indexOf(fragment, candidate)
				if i == -1 {
					break
				}
				if i != 0 {
					newFragments = append(newFragments, fragment[:i])
				}
				fragment = fragment[i+len(candidate):]
			}
			if len(fragment) != 0 {
				newFragments = append(newFragments, fragment)
			}
		}

		// Add candidate to functions.
		newFunctions := make([][]string, 0, 3)
		newFunctions = append(newFunctions, functions...)
		newFunctions = append(newFunctions, candidate)

		subresult := compressPath(path, newFragments, newFunctions)
		result = append(result, subresult...)
	}

	return
}

func hasPrefix(list []string, prefix []string) bool {
	if len(list) < len(prefix) {
		return false
	}

	for i, val := range prefix {
		if list[i] != val {
			return false
		}
	}

	return true
}

func indexOf(list []string, sublist []string) int {
	for i := 0; i <= len(list)-len(sublist); i++ {
		if hasPrefix(list[i:], sublist) {
			return i
		}
	}

	return -1
}
