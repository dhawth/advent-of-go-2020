package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile = "input.txt"
	runeTypeTree = '#'
)

type slope struct {
	over, down int
}

func main() {
	s := slope{
		over: 3,
		down: 1,
	}

	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var trees int
	var x, y int

	for {
		if x == len(lines) {
			break
		}

		line := growString(lines[x], y)

		if line[y] == runeTypeTree {
			trees++
		}

		x += s.down
		y += s.over
	}

	fmt.Printf("We hit %d trees\n", trees)
}

func growString(s string, minLength int) []rune {
	// if we need the 2000th index of a string that is only 15 char long, we need to repeat it
	// 2000 / 15 = 133.33 =~ 134 times.
	n := minLength / len(s) + 1
	return []rune(strings.Repeat(s, n))
}
