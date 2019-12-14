package day14

import (
	"errors"
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var formularRegex = regexp.MustCompile(`(\d+) (\w+)`)

type Formulars map[string]Formular

const (
	fuel         = "FUEL"
	ore          = "ORE"
	oreInventory = 1000000000000
)

type Formular struct {
	producesUnits int
	needs         map[string]int
}

func Answer14() {
	input, _ := Parse()
	fmt.Println(input.CalculateFuelRequirements(1))
	fmt.Println(MaxFuel())
}

func Parse() (Formulars, error) {
	input := util.ReadStringLinesFromFile("resources/day14/input.txt")
	return BuildEquation(input)
}

func BuildEquation(input []string) (Formulars, error) {
	formulars := make(Formulars)

	for _, line := range input {
		equation := strings.Split(line, " => ")
		if len(equation) != 2 {
			return nil, errors.New("invalid input")
		}

		ingredients := formularRegex.FindAllStringSubmatch(equation[0], -1)
		produces := formularRegex.FindAllStringSubmatch(equation[1], -1)

		ret := make(map[string]int)
		for _, ingredient := range ingredients {
			ret[ingredient[2]], _ = strconv.Atoi(ingredient[1])
		}
		
		producesUnits, _ := strconv.Atoi(produces[0][1])
		producesChemical := produces[0][2]

		formulars[producesChemical] = Formular{
			needs:         ret,
			producesUnits: producesUnits,
		}
	}

	return formulars, nil
}

func MaxFuel() int {
	chemicals, _ := Parse()
	
	lx := sort.Search(oreInventory, func(n int) bool {
		return chemicals.CalculateFuelRequirements(n) > oreInventory
	}) - 1

	return lx
}

func (formular Formulars) CalculateFuelRequirements(amount int) int {
	if _, found := formular[fuel]; !found {
		log.Fatal("no idea how to refine fuel")
	}

	want := make(map[string]int, len(formular))
	want[fuel] = amount

	for !isProductionFinished(want) {
		formular.updateRequirements(want)
	}

	return want[ore]
}

func isProductionFinished(want map[string]int) bool {
	for chemical, amount := range want {
		// Ignore ORE as it's the source of every refined chemical
		if chemical != ore && amount > 0 {
			return false
		}
	}

	return true
}

func (formular Formulars) updateRequirements(want map[string]int) {
	for what := range want {
		if what == ore || want[what] <= 0 {
			continue
		}

		unitsProduced := want[what] / formular[what].producesUnits
		spill := want[what] % formular[what].producesUnits

		// We're going to express the need for `what` by the Formulars
		// needed to refine 'what' later, so reset the amount here.
		want[what] = 0
		
		if spill != 0 {
			want[what] -= formular[what].producesUnits - spill
			unitsProduced += 1
		}
		
		for chemical, amount := range formular[what].needs {
			want[chemical] += unitsProduced * amount
		}
	}
}

