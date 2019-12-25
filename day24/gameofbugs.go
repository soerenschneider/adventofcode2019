package day24

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

const (
	gridSize = 5
	maxIterations = 200
	bug = '#'
)

type Level [gridSize][gridSize]bool

func part1(input []string) int {
	state := getState(input)
	visited := make(map[int]bool)

	// Loop until state is reached twice
	for !visited[state] {
		visited[state] = true

		var nextStage int

		for y := 0; y < gridSize; y++ {
			for x := 0; x < gridSize; x++ {
				neighbours := getNeighbours(x, y, state)
				bug := 1 << int(gridSize*y+x)

				// A bug dies (becoming an empty space) unless there is exactly one bug adjacent to it.
				// An empty space becomes infested with a bug if exactly one or two bugs are adjacent to it.
				if ((state&bug != 0) && neighbours == 1) || ((state&bug == 0) && neighbours >= 1 && neighbours <= 2) {
					nextStage |= bug
				}
			}
		}

		state = nextStage
	}

	return state
}

func getState(input []string) (state int) {
	for y, line := range input {
		for x, char := range line {
			if char == bug {
				state |= 1 << int(gridSize*y+x)
			}
		}
	}
	return 
}

func getNeighbours(x, y, state int) (neighbours int) {
	if x > 0 && state&(1<<int(gridSize*y+x-1)) != 0 {
		neighbours++
	}
	if y > 0 && state&(1<<int(gridSize*(y-1)+x)) != 0 {
		neighbours++
	}
	if x < gridSize- 1 && state&(1<<int(gridSize*y+x+1)) != 0 {
		neighbours++
	}
	if y < gridSize- 1 && state&(1<<int(gridSize*(y+1)+x)) != 0 {
		neighbours++
	}
	return
}

func buildLevel(lines []string) (layer Level) {
	for y, line := range lines {
		for x, char := range line {
			if char == bug {
				layer[y][x] = true
			}
		}
	}
	return layer
}

func part2(lines []string) int {
	startLevel := buildLevel(lines)

	levels := make(map[int]Level)
	levels[0] = startLevel
	
	min, max := 0, 0

	for minute := 0; minute < maxIterations; minute++ {
		next := make(map[int]Level)

		for curLevel := min - 1; curLevel <= max+1; curLevel++ {
			var nextLevel Level

			for y := 0; y < gridSize; y++ {
				for x := 0; x < gridSize; x++ {
					// Ignore center
					if x == gridSize/2 && y == gridSize/2 {
						continue
					}

					neighbors := findRecursiveNeighbours(x, y, curLevel, levels)
					nextLevel[y][x] = (levels[curLevel][y][x] && neighbors == 1) || (!levels[curLevel][y][x] && neighbors >= 1 && neighbors <= 2)
				}
			}

			next[curLevel] = nextLevel
		}

		levels = next
		min, max = min-1, max+1
	}

	bugCount := countBugs(levels)
	return bugCount
}

func countBugs(state map[int]Level) (bugs int) {
	for _, level := range state {
		for y := 0; y < gridSize; y++ {
			for x := 0; x < gridSize; x++ {
				if level[y][x] {
					bugs++
				}
			}
		}
	}
	
	return 
}

func findRecursiveNeighbours(x, y, curLevel int, levels map[int]Level) (neighbours int) {
	if x > 0 && levels[curLevel][y][x-1] {
		neighbours++
	}
	if y > 0 && levels[curLevel][y-1][x] {
		neighbours++
	}
	if x < gridSize-1 && levels[curLevel][y][x+1] {
		neighbours++
	}
	if y < gridSize-1 && levels[curLevel][y+1][x] {
		neighbours++
	}

	if x == 0 && levels[curLevel-1][2][1] {
		neighbours++
	}
	if y == 0 && levels[curLevel-1][1][2] {
		neighbours++
	}
	if x == gridSize-1 && levels[curLevel-1][2][3] {
		neighbours++
	}
	if y == gridSize-1 && levels[curLevel-1][3][2] {
		neighbours++
	}

	if x == 1 && y == 2 {
		for i := 0; i < gridSize; i++ {
			if levels[curLevel+1][i][0] {
				neighbours++
			}
		}
	}

	if x == 3 && y == 2 {
		for i := 0; i < gridSize; i++ {
			if levels[curLevel+1][i][gridSize-1] {
				neighbours++
			}
		}
	}

	if y == 1 && x == 2 {
		for i := 0; i < gridSize; i++ {
			if levels[curLevel+1][0][i] {
				neighbours++
			}
		}
	}

	if y == 3 && x == 2 {
		for i := 0; i < gridSize; i++ {
			if levels[curLevel+1][gridSize-1][i] {
				neighbours++
			}
		}
	}

	return 
}

func Answer24() {
	lines := util.ReadStringLinesFromFile("resources/day24/input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
