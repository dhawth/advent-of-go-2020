package main

import (
	"log"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile = "input.txt"
)

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	total := 0
	v := map[rune]struct{}{}

	for _, line := range lines {
		if line == "" {
			total += len(v)
			v = map[rune]struct{}{}
			continue
		}

		for _, r := range line {
			v[r] = struct{}{}
		}
	}

	total += len(v)

	log.Printf("Total: %d", total)
}
