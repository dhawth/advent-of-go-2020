package main

import (
	"log"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile = "input.txt"
)

type group struct {
	numMembers int
	yesses     map[rune]int
}

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var groups []group

	v := map[rune]int{}
	numMembers := 0

	for _, line := range lines {
		if line == "" {
			groups = append(groups, group{
				numMembers: numMembers,
				yesses:     v,
			})

			v = map[rune]int{}
			numMembers = 0

			continue
		}

		for _, r := range line {
			n, found := v[r]
			if found {
				v[r] = n + 1
			} else {
				v[r] = 1
			}
		}

		numMembers++
	}

	groups = append(groups, group{
		numMembers: numMembers,
		yesses:     v,
	})

	total := 0

	for _, g := range groups {
		for _, n := range g.yesses {
			if n == g.numMembers {
				total++
			}
		}
	}

	log.Printf("Total: %d", total)
}
