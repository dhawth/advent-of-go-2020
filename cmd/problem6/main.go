package main

import (
	"fmt"
	"github.com/dhawth/advent-of-go-2020/lib"
	"log"
)

const (
	inputFile = "input.txt"
	runeTypeTree = '#'
)

type slope struct {
	over, down int
}

func main() {
	slopes := []slope{
		{
			over: 1,
			down: 1,
		},
		{
			over: 3,
			down: 1,
		},
		{
			over: 5,
			down: 1,
		},
		{
			over: 7,
			down: 1,
		},
		{
			over: 1,
			down: 2,
		},
	}

	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var treeCounts []int

	for _, s := range slopes {
		var trees int
		var x, y int

		for {
			if x >= len(lines) {
				break
			}

			newYIndex := recalculateIndex(lines[x], y)

			if []rune(lines[x])[newYIndex] == runeTypeTree {
				trees++
			}

			x += s.down
			y += s.over
		}

		fmt.Printf("We hit %d trees for slope %+v\n", trees, s)
		treeCounts = append(treeCounts, trees)
	}

	foo := 1

	for _, n := range treeCounts {
		foo *= n
	}
	fmt.Printf("We hit %d trees\n", foo)
}

func recalculateIndex(s string, n int) int {
	return n % len(s)
}
