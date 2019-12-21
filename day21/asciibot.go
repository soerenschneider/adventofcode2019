package day21

import (
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
	"log"
	"strings"
)

const (
	Walk              mode = "WALK"
	Run               mode = "RUN"
	BotInputPrompt         = "Input instructions:"
	botFinishedSignal      = '\n'
	maxScriptLength        = 15
)

type mode string

type asciiBot struct {
	mode         mode
	script       []string
	outputBuffer strings.Builder
}

func NewAsciiBot(mode mode, script []string) (*asciiBot, error) {
	if len(script) > maxScriptLength {
		return nil, fmt.Errorf("max script length succeeded %d", maxScriptLength)
	}

	return &asciiBot{
		mode:         mode,
		script:       script,
		outputBuffer: strings.Builder{},
	}, nil
}

func (b *asciiBot) processOutput(botOutput int64, botInput chan int64) int64 {
	if !isAscii(botOutput) {
		return botOutput
	} else if botOutput == botFinishedSignal {
		message := b.outputBuffer.String()
		switch message {
		case BotInputPrompt:
			formatted := b.format()
			send(formatted, botInput)
		}
		b.outputBuffer.Reset()
	} else {
		mb := byte(botOutput)
		b.outputBuffer.WriteByte(mb)
	}

	return -1
}

func (b *asciiBot) format() string {
	msg := strings.Join(b.script, "\n")
	return  msg + "\n" + string(b.mode) + "\n"
}

func send(input string, in chan<- int64) {
	for _, c := range input {
		in <- int64(c)
	}
}

func isAscii(i int64) bool {
	return 0 <= i && i <= 128
}

func Answer21() {
	input := util.ReadInt64Array("resources/day21/input.txt")
	in := make(chan int64)
	out := make(chan int64)

	interpreter := util.NewInterpreter(input, in, out)

	script := []string{
		"OR A J",
		"AND B J",
		"AND C J",
		"NOT J J",
		"AND D J",
	}

	bot, err := NewAsciiBot(Walk, script); if err != nil {
		log.Fatal(err)
	}

	go interpreter.Execute()
	var result int64
	for output := range out {
		result = bot.processOutput(output, in)
	}

	fmt.Println(result)
}

func Answer21b() {
	input := util.ReadInt64Array("resources/day21/input.txt")
	in := make(chan int64)
	out := make(chan int64)

	interpreter := util.NewInterpreter(input, in, out)

	script := []string{
		"OR A J",
		"AND B J",
		"AND C J",
		"NOT J J",
		"AND D J",
		"OR H T",
		"OR E T",
		"AND T J",
	}

	bot, err := NewAsciiBot(Run, script); if err != nil {
		log.Fatal(err)
	}

	go interpreter.Execute()
	var result int64

	for output := range out {
		result = bot.processOutput(output, in)
	}

	fmt.Println(result)
}
