package main

import (
	"fmt"
	"log"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile = "input.txt"
)

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	acc := 0
	lineIndex := 0
	visitedLineNumbers := lib.NewIntSet()

	for {
		if visitedLineNumbers.Contains(lineIndex) {
			break
		}

		visitedLineNumbers.Add(lineIndex)

		line := lines[lineIndex]
		var command string
		var value int
		_, err := fmt.Sscanf(line, "%s %d", &command, &value)

		if err != nil {
			log.Fatalf("error scanning line %s: %v", line, err)
		}

		switch command {
		case "acc":
			acc += value
			lineIndex++
		case "nop":
			lineIndex++
		case "jmp":
			lineIndex += value
		default:
			log.Fatalf("unknown command: %s on line %d", command, lineIndex + 1)
		}
	}

	log.Printf("accumulator value pre-loop is %d", acc)
}
