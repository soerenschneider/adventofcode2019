package util

import (
	"io/ioutil"
	"log"
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
