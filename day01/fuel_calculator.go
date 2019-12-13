package day01

import (
	"errors"
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"math"
)

type FuelCalculator interface {
	RequiredFuel(mass int) int
}

type Calc01 struct {
}

// Fuel required to launch a given Module is based on its mass. Specifically, to find the fuel required for a Module,
// take its mass, divide by three, round down, and subtract 2.
func (m *Calc01) RequiredFuel(mass int) int {
	return int(math.Floor(float64(mass) / 3)) - 2
}

type Calc01b struct {
	naive Calc01
}

// So, for each module mass, calculate its fuel and add it to the total. Then, treat the fuel amount you just
// calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative.
func (m *Calc01b) RequiredFuel(mass int) int {
	x := m.naive.RequiredFuel(mass)
	if x <= 0 || mass <= 0 {
		return 0
	}

	return x + m.RequiredFuel(x)
}

func Answer() {
	modules := util.ReadIntLines("resources/day01/input.txt")
	naive := &Calc01{}
	x, _ := ProcessInput(naive, modules)
	fmt.Printf("Answer 01: %d\n", x)

	advanced := &Calc01b{}
	x, _ = ProcessInput(advanced, modules)
	fmt.Printf("Answer 01b: %d\n", x)
}

func ProcessInput(strategy FuelCalculator, input []int) (int, error) {
	if input == nil {
		return -1, errors.New("invalid input")
	}
	
	sum := 0
	for _, value := range input {
		sum += strategy.RequiredFuel(value)
	}
	
	return sum, nil
}

