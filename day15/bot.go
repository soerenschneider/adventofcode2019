package day15

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

const (
	ResponseWall  = 0
	ResponseMoved = 1
	ResponseFound = 2
)

const (
	MapWall       = 1
	MapEmptySpace = 2
	MapOxygen     = 3
	MapDeadEnd    = 4
	MapUnknown    = 0
)

var movements = map[int64]util.Coordinate64 {
	North: {X: 0, Y: -1},
	South: {X: 0, Y: 1},
	West: {X: -1, Y: 0},
	East: {X: 1, Y: 0},
}

const (
	North = 1
	South = 2
	West = 3
	East = 4
)

type Bot struct {
	position  util.Coordinate64
	oxygen    util.Coordinate64
	grid      Map
	direction int64
}

type Map map[util.Coordinate64]int64

func NewBot() (bot Bot) {
	bot = Bot{
		grid: make(Map),
		direction: North,
	}

	bot.grid[util.Coordinate64{}] = MapEmptySpace
	return
}

func ExploreSpaceMap() Bot {
	input := util.ReadInt64Array("resources/day15/input.txt")
	in := make(chan int64)
	out := make(chan int64, 1)

	interpreter := util.NewInterpreter(input, in, out)
	bot := NewBot()

	go interpreter.Execute()

	in <- bot.direction
	for response := range out {
		moveSuccessful := bot.move(response)
		if moveSuccessful {
			in <- bot.direction
		} else {
			in <- 0
		}
	}

	bot.grid[bot.oxygen] = MapOxygen
	return bot
}

// checkDeadEnd checks the adjacent fields for blocker elements. If a blocker is found,
// the current position is marked as as dead end.
func (b *Bot) checkDeadEnd() {
	wallCount := 0
	
	for direction := range movements {
		move := b.position.Move(movements[direction])
		spaceElement := b.grid[move]
		if spaceElement == MapWall || spaceElement == MapDeadEnd {
			wallCount++
		}
	}

	if wallCount >= 3 {
		b.grid[b.position] = MapDeadEnd
	}
}

func (b *Bot) setGridValue(response int64, target util.Coordinate64) {
	if _, ok := b.grid[target]; !ok {
		if response == ResponseWall {
			b.grid[target] = MapWall
		} else {
			b.grid[target] = MapEmptySpace
		}
	}
}

func (b *Bot) move(response int64) bool {
	target := b.position.Move(movements[b.direction])

	b.setGridValue(response, target)
	b.checkDeadEnd()
	b.updatePosition(response, target)

	// try to move to an adjacent location that is unknown
	movementSuccessful := b.tryMoveTo(MapUnknown)
	if !movementSuccessful {
		// try to move to an adjacent empty space
		movementSuccessful = b.tryMoveTo(MapEmptySpace)
	}

	return movementSuccessful
}

// tryMoveTo accepts a desired element of the space map and tries a move to
// all adjacent positions. The bots turn is set to reach the first position that matches the desired
// element and true is returned. If no adjacent move to a desired element is possible, no changes are
// made and false is returned.
func (b *Bot) tryMoveTo(desiredPositionType int64) bool {
	for direction := range movements {
		move := movements[direction]
		newPos := b.position.Move(move)

		if b.grid[newPos] == desiredPositionType {
			b.direction = direction
			return true
		}
	}
	return false
}

// updatePosition marks the map with the answer of the intcode computer.
func (b *Bot) updatePosition(response int64, target util.Coordinate64) {
	switch response {
	case ResponseFound:
		b.oxygen = target
		b.position = target
	case ResponseMoved:
		b.position = target
	}
}

func (m Map) traverse(grid map[util.Coordinate64]int64, pos, to util.Coordinate64, steps int64) {
	_, visited := grid[pos]
	if visited || m[pos] == MapWall {
		return
	}
	grid[pos] = steps

	if pos == to {
		return
	}

	for direction := range movements {
		move := movements[direction]
		m.traverse(grid, pos.Move(move), to, steps + 1)
	}
}

func (b *Bot) Answer15b() int64 {
	distances := make(map[util.Coordinate64]int64)
	b.grid.traverse(distances, b.oxygen, util.Coordinate64{}, 0)

	var maxDistance int64
	for _, distance := range distances {
		maxDistance = util.MaxInt64(distance, maxDistance)
	}

	return maxDistance
}

func Answer15() {
	bot := ExploreSpaceMap()
	grid := make(map[util.Coordinate64]int64)
	bot.grid.traverse(grid, util.Coordinate64{}, bot.oxygen, 0)
	fmt.Printf("Answer 15: %d\n", grid[bot.oxygen])

	t := bot.Answer15b()
	fmt.Printf("Answer15b: %d m\n", t)
}