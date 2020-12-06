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

	var groups []map[rune]struct{}

	v := map[rune]struct{}{}

	for _, line := range lines {
		if line == "" {
			groups = append(groups, v)
			v = map[rune]struct{}{}
			continue
		}

		for _, r := range line {
			v[r] = struct{}{}
		}
	}

	groups = append(groups, v)
	total := 0

	for _, m := range groups {
		total += len(m)

	}

	log.Printf("Total: %d", total)
}

