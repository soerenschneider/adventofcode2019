package day10

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"log"
	"math"
	"sort"
)

const (
	asteroid = '#'
)

type void struct{}

func ParseCoordinates(input []string) []Coordinate {
	var asteroids []Coordinate

	for x, line := range input {
		for y, value := range line {
			if value == asteroid {
				asteroids = append(asteroids, Coordinate{Y: y, X: x})
			}
		}
	}

	return asteroids
}

func DetermineMaxObservableAsteroids(asteroids []Coordinate) (Coordinate, int) {
	maxObservableAsteroids := 0
	var maxCoord Coordinate

	for _, source := range asteroids {
		targets := make(map[float64]void)
		oberservableAsteroids := 0

		for _, target := range asteroids {
			if source == target {
				continue
			}

			angle := source.Angle(target)
			_, found := targets[angle]
			if !found {
				targets[angle] = void{}
				oberservableAsteroids++
			}
		}

		if oberservableAsteroids > maxObservableAsteroids {
			maxObservableAsteroids = oberservableAsteroids
			maxCoord = source
		}
	}

	return maxCoord, maxObservableAsteroids
}

func getSortedAngles(targets map[float64]map[int]Coordinate) []float64 {
	angles := make([]float64, len(targets))
	i := 0
	for k := range targets {
		angles[i] = k
		i++
	}
	sort.Float64s(angles)
	return angles
}

func getTargetsWithDistance(coordinates []Coordinate, source Coordinate) map[float64]map[int]Coordinate {
	targets := make(map[float64]map[int]Coordinate)

	for _, target := range coordinates {
		if target != source {
			angle := source.Angle(target)
			distance := source.Distance(target)

			_, found := targets[angle]
			if !found {
				// Build list to append other asteroids to
				targets[angle] = make(map[int]Coordinate)
			}
			targets[angle][distance] = target
		}
	}

	return targets
}

func GetNthVaporizedAsteroid(targets map[float64]map[int]Coordinate, iteration int) (Coordinate, error) {
	angles := getSortedAngles(targets)

	hits := 0
	for len(angles) > 0 {
		for _, angle := range angles {
			closest := -1
			hits++

			for dist := range targets[angle] {
				if dist < closest || closest == -1 {
					closest = dist
				}
			}

			if hits == iteration {
				return targets[angle][closest], nil
			}
			delete(targets[angle], closest)
		}

		var clean []float64
		for _, angle := range angles {
			if len(targets[angle]) > 0 {
				clean = append(clean, angle)
			}
		}
		angles = clean
	}

	return Coordinate{}, fmt.Errorf("there were only %d hits", hits)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) flip() Coordinate {
	return Coordinate{X: c.Y, Y: c.X}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%d, %d", c.Y, c.X)
}

func (c *Coordinate) Distance(other Coordinate) int {
	return AbsInt(c.X-other.X) + AbsInt(c.Y-other.Y)
}

func (c *Coordinate) Angle(t Coordinate) float64 {
	arctan := math.Atan2(float64(t.Y-c.Y), float64(c.X-t.X))
	angle := arctan * 180 / math.Pi
	if angle < 0 {
		angle = angle + 360
	}
	return angle
}

func Answer10() {
	input := util.ReadStringLinesFromFile("resources/day10/input.txt")
	asteroids := ParseCoordinates(input)
	DetermineMaxObservableAsteroids(asteroids)
	coord, observable := DetermineMaxObservableAsteroids(asteroids)
	fmt.Printf("Coordinates: %v, overseeing %d asteroids\n", coord.flip(), observable)

	targets := getTargetsWithDistance(asteroids, coord)
	iteration := 200
	jackpot, err := GetNthVaporizedAsteroid(targets, iteration)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Iteration %d: %d", iteration, jackpot.Y*100+jackpot.X)
}