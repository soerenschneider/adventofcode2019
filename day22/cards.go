package day22

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"strconv"
	"strings"
)

type cards struct {
	deck []int
}

const (
	instIncre = "deal with increment "
	instCut = "cut "
	instDeal = "deal into new stack"
)

func Answer22() {
	cards := ReadInstructionsAndShuffle("resources/day22/input.txt", 10007)
	for i, card := range cards.deck {
		if card == 2019 {
			fmt.Println(i)
		}
	}
}

func ReadInstructionsAndShuffle(file string, n int) *cards {
	input := util.ReadStringLinesFromFile(file)
	cards := NewCardDeck(n)
	cards.Shuffle(input)
	return cards
}

func NewCardDeck(n int) *cards {
	ret := &cards{
		deck: make([]int, n),
	}

	for i := 0; i < n; i++ {
		ret.deck[i] = i
	}
	
	return ret
}

func parseArg(instr string) int {
	cut := -1
	if strings.HasPrefix(instr, instIncre) {
		cut = len(instIncre)
	} else if strings.HasPrefix(instr, instCut) {
		cut = len(instCut)
	}
	
	if cut == -1 {
		return -1
	}

	paramIndex := instr[cut:]
	i, _ := strconv.Atoi(paramIndex)
	return i
}

func (c *cards) Shuffle(instructions []string) {
	for _, instr := range instructions {
		if strings.HasPrefix(instr, instIncre) {
			n := parseArg(instr)
			c.Increment(n)
		} else if strings.HasPrefix(instr, instCut) {
			n := parseArg(instr)
			c.Cut(n)
		} else if strings.HasPrefix(instr, instDeal){
			c.Deal()
		} else {
			fmt.Printf("Didn't understand cmd: %s\n", instr)
		}
	}
}

func (c *cards) Deal() {
	new := make([]int, len(c.deck))
	negIndex := len(c.deck) - 1
	for index, _ := range c.deck {
		new[index] = c.deck[negIndex]
		negIndex--
	}
	c.deck = new
}

func (c* cards) Cut(n int) {
	if n < 0 {
		n = len(c.deck) - n * -1
	}

	cut := c.deck[:n]
	c.deck = c.deck[n:]
	c.deck = append(c.deck, cut...)
}

func (c *cards) Increment(n int) {
	new := make([]int, len(c.deck))
	for index, card := range c.deck {
		new[(index * n)  % len(c.deck)] = card
	}

	c.deck = new
}