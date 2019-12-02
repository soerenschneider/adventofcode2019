package day01

import "math"

type Calc01 struct {
}

// Fuel required to launch a given Module is based on its mass. Specifically, to find the fuel required for a Module,
// take its mass, divide by three, round down, and subtract 2.
func (m *Calc01) RequiredFuel(mass int) int {
	return int(math.Floor(float64(mass) / 3)) - 2
}