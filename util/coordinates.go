package util

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
