package day04

import (
	"fmt"
	"regexp"
	"strconv"
)

var numbersOnlyRegex = regexp.MustCompile("^[0-9]+$")

func isAdjacent(candidate string) bool {
	if ! numbersOnlyRegex.MatchString(candidate) {
		return false
	}

	var prev rune
	for pos, char := range candidate {
		if pos > 0 && char == prev {
			return true
		}
		prev = char
	}
	return false
}

// 2nd part
func isAdjacent2nd(candidate string) bool {
	if ! numbersOnlyRegex.MatchString(candidate) {
		return false
	}

	var prev rune
	streak := 1
	for pos, char := range candidate {
		if pos > 0 && char == prev {
			streak++
		} else {
			// The streak has been broken – if our streak lasted only for 2
			// iterations, the filter is already passed.
			if streak == 2 {
				return true
			}
			// Otherwise reset counting
			streak = 1
		}
		prev = char
	}
	// If we're at the end and the streak is 2 – we've done it again!
	return streak == 2
}

func isIncreasing(candidate string) bool {
	var prev int
	for pos, char := range candidate {
		current, err := strconv.Atoi(string(char)); if err != nil {
			return false
		}
		if pos > 0 && current < prev {
			return false
		}
		prev = current
	}
	return true
}

type meetsRules func(candidate string) bool

func meetsRules1st(candidate string) bool {
	return isAdjacent(candidate) && isIncreasing(candidate)
}

func meetsRules2nd(candidate string) bool {
	return isAdjacent2nd(candidate) && isIncreasing(candidate)
}

func testRange(from, to int, fn meetsRules) int {
	hits := 0
	for candidate := from; candidate <= to; candidate++ {
		s := fmt.Sprintf("%07d", candidate)
		if fn(s) {
			hits++
		}
	}
	
	return hits
}

func Answer04() {
	fmt.Println("---------------------")

	a4 := testRange(357253,892942, meetsRules1st)
	fmt.Printf("Answer 04: %d\n", a4)
	
	a4b := testRange(357253,892942, meetsRules2nd)
	fmt.Printf("Answer 04b: %d\n", a4b)
}