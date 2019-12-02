package day01

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type FuelCalculator interface {
	RequiredFuel(mass int) int
}

func Answer() {
	modules, err := ReadModules("resources/day01/input.txt")
	if err != nil {
		log.Fatalf("error parsing modules: %s", err.Error())
	}

	naive := &Calc01{}
	x, err := ProcessInput(naive, modules)
	fmt.Printf("Answer 01: %d\n", x)

	advanced := &Calc01b{}
	x, err = ProcessInput(advanced, modules)
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

func ReadModules(path string) ([]int, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, errors.New("could not open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []int

	for scanner.Scan() {
		line := scanner.Text()
		intValue, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("couldn't convert value: %s", err.Error())
		}
		result = append(result, intValue)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("couldn't read file: %s", err.Error())
	}

	return result, nil
}
