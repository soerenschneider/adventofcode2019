package util

import "fmt"

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
