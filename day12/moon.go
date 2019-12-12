package day12

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"regexp"
	"strconv"
)

var parseVectorPattern = regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)

const (
	X = 0
	Y = 1
	Z = 2
)

type Vector struct {
	X int
	Y int
	Z int
}

func (v *Vector) String() string {
	return fmt.Sprintf("(%d, %d, %d)\n", v.X, v.Y, v.Z)
}

func (v *Vector) Sum() int {
	return util.Abs(v.X) + util.Abs(v.Y) + util.Abs(v.Z)
}

func (v *Vector) HasNullValue() bool {
	return v.X == 0 || v.Y == 0 || v.Z == 0
}

type Moon struct {
	position Vector
	velocity Vector
}

func (m *Moon) ApplyGravity(other Moon) {
	if m.position.X < other.position.X {
		m.velocity.X++
	} else if m.position.X > other.position.X {
		m.velocity.X--
	}

	if m.position.Y < other.position.Y {
		m.velocity.Y++
	} else if m.position.Y > other.position.Y {
		m.velocity.Y--
	}

	if m.position.Z < other.position.Z {
		m.velocity.Z++
	} else if m.position.Z > other.position.Z {
		m.velocity.Z--
	}
}

func (m *Moon) ApplyVelocity() {
	m.position.X += m.velocity.X
	m.position.Y += m.velocity.Y
	m.position.Z += m.velocity.Z
}

func (m *Moon) PotentialEnergy() int {
	return m.velocity.Sum() * m.position.Sum()
}

func Cycle(moons []Moon, steps int) {
	for step := 0; step < steps; step++ {
		for a := range moons {
			for b := range moons {
				if a != b {
					moons[a].ApplyGravity(moons[b])
				}
			}
		}

		for moon := range moons {
			moons[moon].ApplyVelocity()
		}
	}
}

func Sum(moons []Moon) (sum int) {
	sum = 0
	for moon := range moons {
		sum += moons[moon].PotentialEnergy()
	}
	return
}

// setAxisValue sets the given value to all the axes whose index evaluate to true. 
func (v *Vector) setAxisValue(axis []bool, val int) error {
	if len(axis) < 3 {
		return fmt.Errorf("|axis| != 3")
	}
	if axis[X] {
		v.X = val
	}

	if axis[Y] {
		v.Y = val
	}

	if axis[Z] {
		v.Z = val
	}
	
	return nil
}

func (m *Moon) HasIdenticalAxis(other Moon, axis int) bool {
	switch axis {
	case X:
		return m.position.X == other.position.X && m.velocity.X == other.velocity.X
	case Y:
		return m.position.Y == other.position.Y && m.velocity.Y == other.velocity.Y
	case Z:
		return m.position.Z == other.position.Z && m.velocity.Z == other.velocity.Z
	}
	return false
}

func updateAxisInformation(moon Moon, other Moon, identicalAxis []bool) {
	// we only account negative cases – once we found the steps for an axis we
	// don't want to update the steps vector on the appropriate axis
	if ! moon.HasIdenticalAxis(other, X) {
		identicalAxis[X] = false
	}

	if ! moon.HasIdenticalAxis(other, Y) {
		identicalAxis[Y] = false
	}

	if ! moon.HasIdenticalAxis(other, Z) {
		identicalAxis[Z] = false
	}
}

func Answer12() {
	moons := ParseMoons("resources/day12/test.txt")
	steps := 1000
	Cycle(moons, steps)
	sum := Sum(moons)
	fmt.Println(sum)
}

func Answer12Part2() {
	moons := ParseMoons("resources/day12/input.txt")

	initialState := make([]Moon, len(moons))
	copy(initialState, moons)

	// Store steps for each axis to reach initial state again
	var stepsTaken Vector
	
	// repeat until each axis is different from 0
	for step := 1; stepsTaken.HasNullValue() ; step++ {
		for a := range moons {
			for b := range moons {
				if a != b {
					moons[a].ApplyGravity(moons[b])
				}
			}
		}
		
		// the index of this var evaluates to true if the given axis is back its initital state
		axisAtInitialState := getAbsAxis(stepsTaken)

		// Perform a complete cycle with all moons
		for moon := range moons {
			moons[moon].ApplyVelocity()
			updateAxisInformation(moons[moon], initialState[moon], axisAtInitialState)
		}
		
		// Check whether at least one of the axis is back at square one
		stepsTaken.setAxisValue(axisAtInitialState, step)
	}

	steps := util.LCM(stepsTaken.X, stepsTaken.Y, stepsTaken.Z)
	fmt.Println(steps)
}

// getAbsAxis returns a slice of bools whose index are true if
// the appropriate axis is not equals zero.
func getAbsAxis(v Vector) []bool {
	axes := make([]bool, 3)
	
	axes[X] = v.X == 0
	axes[Y] = v.Y == 0
	axes[Z] = v.Z == 0
	
	return axes
}

func ParseMoons(filepath string) (moons []Moon) {
	input := util.ReadString(filepath)
	vectors := ParseVectors(input)
	return BuildMoons(vectors)
}

func BuildMoons(vectors []Vector) (moons []Moon) {
	for _, v := range vectors {
		moons = append(moons, Moon{position: v})
	}

	return
}

func ParseVectors(input string) []Vector {
	var vectors []Vector

	for _, match := range parseVectorPattern.FindAllStringSubmatch(input, -1) {
		var v Vector
		v.X, _ = strconv.Atoi(match[1])
		v.Y, _ = strconv.Atoi(match[2])
		v.Z, _ = strconv.Atoi(match[3])

		vectors = append(vectors, v)
	}

	return vectors
}
