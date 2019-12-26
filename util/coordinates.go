package util

import "fmt"

var (
	North = Coordinate{X: 0, Y: -1}
	South = Coordinate{X: 0, Y: 1}
	West  = Coordinate{X: -1, Y: 0}
	East  = Coordinate{X: 1, Y: 0}
	Adjacent = []Coordinate{North, South, West, East}
	TurnLeft = map[Coordinate]Coordinate{
		North: West,
		West:  South,
		South: East,
		East:  North}

	TurnRight = map[Coordinate]Coordinate{
		North: East,
		East:  South,
		South: West,
		West:  North}
)

type Coordinate64 struct {
	X int64
	Y int64
}

func (c Coordinate64) Move(t Coordinate64) Coordinate64 {
	return Coordinate64{
		X: c.X + t.X,
		Y: c.Y + t.Y,
	}
}

func (c *Coordinate64) String() string {
	return fmt.Sprintf("%d, %d", c.X, c.Y)
}

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Move(t Coordinate) Coordinate {
	return Coordinate{
		X: c.X + t.X,
		Y: c.Y + t.Y,
	}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%d, %d", c.X, c.Y)
}
