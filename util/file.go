package util

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadString(path string) string {
	inputBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading file %s: %s", path, err.Error())
	}
	return string(inputBytes)
}

func ReadStringLinesFromFile(path string) []string {
	inputBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading file %s: %s", path, err.Error())
	}
	input := string(inputBytes)
	return strings.Split(input, "\n")
}

func ReadIntLines(path string) []int {
	lines := ReadStringLinesFromFile(path)
	var result []int

	for _, line := range lines {
		intValue, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, intValue)
	}

	return result
}


func ReadIntArray(path string) (mem []int) {
	input, _ := ioutil.ReadFile(path)
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	mem = make([]int, len(split))

	for i, s := range split {
		x, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			log.Fatal("error reading input")
		}
		mem[i] = int(x)
	}

	return
}

func ReadInt64Array(path string) (mem []int64) {
	input, _ := ioutil.ReadFile(path)
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	mem = make([]int64, len(split))

	for i, s := range split {
		x, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			log.Fatal("error reading input")
		}
		mem[i] = x
	}

	return
}
