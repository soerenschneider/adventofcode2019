package day18

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"log"
)

const (
	maxRobots    = 4
	maxKeys 	 = 'Z' - 'A' + 1
	dispPosition = '@'
	dispWall     = '#'
	dispPath     = '.'
)

type state struct {
	Pos    [maxRobots]util.Coordinate
	Keys   [maxKeys]bool
	Active int
}

type maze struct {
	start     state
	goal      [maxKeys]bool
	field     map[util.Coordinate]rune
	numRobots int
}

func NewMaze(input []string) *maze {
	maze := &maze{
		field: map[util.Coordinate]rune{},
	}
	maze.discoverField(input)
	return maze
}

func (m *maze) discoverField(input []string) {
	for y, s := range input {
		for x, r := range s {
			if r == dispPosition {
				if m.numRobots >= maxRobots {
					log.Fatalf("can only handle %d robots", maxRobots)
				}

				m.start.Pos[m.numRobots] = util.Coordinate{X: x, Y: y}
				r = dispPath
				m.numRobots++
			}
			if isKey(r) {
				index := getArrayKeyIndex(r)
				m.goal[index] = true
			}
			m.field[util.Coordinate{X: x, Y: y}] = r
		}
	}
}

func (m *maze) solve() int {
	dist := map[state]int{}
	var queue []state

	for i := 0; i < m.numRobots; i++ {
		m.start.Active = i
		dist[m.start] = 0
		queue = append(queue, m.start)
	}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.Keys == m.goal {
			return dist[state]
		}

		for _, d := range util.Adjacent {
			lookaheadState := state
			lookaheadState.Pos[state.Active] = state.Pos[state.Active].Move(d)
			lookaheadContent := m.field[lookaheadState.Pos[state.Active]]

			if isMoveBlocked(lookaheadContent, state) {
				continue
			}

			if isKey(lookaheadContent) {
				index := getArrayKeyIndex(lookaheadContent)
				lookaheadState.Keys[index] = true
			}

			for i := 0; i < m.numRobots; i++ {
				if i == state.Active || lookaheadState.Keys != state.Keys {
					lookaheadState.Active = i

					if _, found := dist[lookaheadState]; !found {
						dist[lookaheadState] = dist[state] + 1
						queue = append(queue, lookaheadState)
					}
				}
			}
		}
	}

	return -1
}

func isMoveBlocked(fieldContent rune, state state) bool {
	return fieldContent == dispWall || isDoor(fieldContent) && !state.Keys[getArrayDoorIndex(fieldContent)]
}

func getArrayKeyIndex(r rune) int {
	return int(r - 'a')
}

func getArrayDoorIndex(r rune) int {
	return int(r - 'A')
}

func isDoor(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isKey(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func Answer18() {
	files := []string{"resources/day18/input.txt", "resources/day18/input2.txt"}

	for index, file := range files {
		input := util.ReadStringLinesFromFile(file)
		maze := NewMaze(input)
		dist := maze.solve()
		fmt.Printf("Part %d: %d\n", index+1, dist)
	}
}
