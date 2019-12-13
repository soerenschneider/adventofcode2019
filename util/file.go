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

func ReadInt64Array(path string) (mem []int64) {
	input, _ := ioutil.ReadFile(path)
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	mem = make([]int64, len(split))

	for i, s := range split {
		x, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			log.Fatal("error reading input")
		}
		mem[i] = int64(x)
	}

	return
}
