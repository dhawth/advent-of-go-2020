package main

import (
	"fmt"
	"log"

	"github.com/dhawth/advent-of-go-2020/lib"
)

type commandType int

const (
	inputFile = "input.txt"
	commandTypeAcc commandType = iota
	commandTypeNop commandType = iota
	commandTypeJmp commandType = iota
)

var (
	ErrLoopDetected = fmt.Errorf("loop detected")
)

type command struct {
	cmd commandType
	value int
}

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	commands, err := linesToCommands(lines)
	if err != nil {
		log.Fatal(err)
	}

	queueOfCandidatesToChange := detectLoopCandidates(commands)

	for !queueOfCandidatesToChange.IsEmpty() {
		copyOfCommands := copyCommands(commands)
		lineToChange, err := queueOfCandidatesToChange.Shift()
		if err != nil {
			log.Fatal(err)
		}

		switch c := copyOfCommands[lineToChange]; c.cmd {
		case commandTypeNop:
			copyOfCommands[lineToChange] = command{
				cmd: commandTypeJmp,
				value: c.value,
			}
		case commandTypeJmp:
			copyOfCommands[lineToChange] = command{
				cmd: commandTypeNop,
				value: c.value,
			}
		}

		acc, err := testLineSet(copyOfCommands)
		if err == ErrLoopDetected {
			continue
		}

		// success!
		log.Printf("No loop after changing line %d, acc is %d", lineToChange, acc)
		break
	}
}

func copyCommands(commands []command) []command {
	cpy := make([]command, len(commands))
	for i := 0; i < len(commands); i++ {
		cpy[i] = commands[i]
	}

	return cpy
}

func testLineSet(commands []command) (int, error) {
	lineIndex := 0
	acc := 0
	visitedLineNumbers := lib.NewIntSet()

	for {
		if visitedLineNumbers.Contains(lineIndex) {
			return 0, ErrLoopDetected
		}

		visitedLineNumbers.Add(lineIndex)

		switch c := commands[lineIndex]; c.cmd {
		case commandTypeAcc:
			acc += c.value
			lineIndex++
		case commandTypeNop:
			lineIndex++
		case commandTypeJmp:
			lineIndex += c.value
		}

		// if we've reached the end, exit the loop
		if lineIndex == len(commands) {
			break
		}
	}

	return acc, nil
}

func linesToCommands(lines []string) ([]command, error) {
	commands := make([]command, len(lines))

	for i := 0; i < len(lines); i++ {
		var cmd string
		var value int
		_, err := fmt.Sscanf(lines[i], "%s %d", &cmd, &value)

		if err != nil {
			return nil, fmt.Errorf("error scanning line %s: %w", lines[i], err)
		}

		switch cmd {
		case "acc":
			commands[i] = command{
				cmd:   commandTypeAcc,
				value: value,
			}
		case "nop":
			commands[i] = command{
				cmd:   commandTypeNop,
				value: value,
			}
		case "jmp":
			commands[i] = command{
				cmd:   commandTypeJmp,
				value: value,
			}
		default:
			return nil, fmt.Errorf("unknown command: %s on line %d: %s", cmd, i, lines[i])
		}
	}

	return commands, nil
}

func detectLoopCandidates(commands []command) lib.IntQueue{
	lineIndex := 0
	previousLineIndex := 0
	visitedLineNumbers := lib.NewIntSet()
	queueOfCandidatesToChange := lib.NewIntQueue()

	for{
		if visitedLineNumbers.Contains(lineIndex){
			log.Printf("Detected bug on line %d: %v", previousLineIndex, commands[previousLineIndex])
			break
		}

		visitedLineNumbers.Add(lineIndex)

		switch c := commands[lineIndex]; c.cmd{
		case commandTypeAcc:
			previousLineIndex = lineIndex
			lineIndex++
		case commandTypeNop:
			previousLineIndex = lineIndex
			queueOfCandidatesToChange.Push(lineIndex)
			lineIndex++
		case commandTypeJmp:
			previousLineIndex = lineIndex
			queueOfCandidatesToChange.Push(lineIndex)
			lineIndex += c.value
		}
	}

	return queueOfCandidatesToChange
}