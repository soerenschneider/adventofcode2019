package day16

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"regexp"
	"strconv"
	"strings"
)

const (
	phases = 100
	repeat = 10000
	offsetStart = 7
)

var (
	pattern          = []int{0, 1, 0, -1}
	numbersOnlyRegex = regexp.MustCompile("^[0-9]+$")
)

func Convert(input string) []int {
	if ! numbersOnlyRegex.MatchString(input) {
		return []int{}
	}

	ret := make([]int, len(input))
	for i, r := range input {
		ret[i] = int(r - '0')
	}
	return ret
}

func CleanSignal(signal []int, phases int) string {
	Phases(signal, phases)

	digits := 8
	var ret strings.Builder
	for i := 0; i < util.MinInt(digits, len(signal)); i++ {
		val := strconv.Itoa(signal[i])
		ret.WriteString(val)
	}

	return ret.String()
}

func Phases(signal []int, phases int) {
	for phase := 1; phase <= phases; phase++ {
		for iteration := 1; iteration <= len(signal); iteration++ {
			signal[iteration-1] = Apply(signal, iteration)
		}
	}
}

func Apply(signal []int, iteration int) int {
	ret := 0

	patternIndex := 1
	for j, digit := range signal {
		patternVal := pattern[(j+1)/iteration%4]
		ret += patternVal * digit
		patternIndex += 1
	}

	return util.Abs(ret) % 10
}

func ExtractEmbeddedSignal(content string) []int {
	// extract message offset
	offset, _ := strconv.Atoi(content[:offsetStart])

	// build signal from content without the offset
	signal := Convert(content)[offset:]

	ret := make([]int, len(signal))
	for i, c := range signal {
		ret[i] = c
	}

	for phase := 1; phase <= phases; phase++ {
		sum := 0
		for i := len(ret) - 1; i >= 0; i-- {
			sum += ret[i]
			ret[i] = sum % 10
		}
	}

	return ret
}

func Answer16() {
	content := util.ReadString("resources/day16/input.txt")
	signal := Convert(content)
	fmt.Println(CleanSignal(signal, phases))
}

func Answer16b() {
	content := util.ReadString("resources/day16/input.txt")
	content = strings.Repeat(content, repeat)

	output := ExtractEmbeddedSignal(content)
	var ret strings.Builder
	for _, c := range output[:offsetStart + 1] {
		ret.WriteString(strconv.Itoa(c))
	}
	fmt.Println(ret.String())
}
