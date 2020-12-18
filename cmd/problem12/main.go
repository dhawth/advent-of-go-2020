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
		log.Fatalf("error reading file: %s", err)
	}

	total := 0
	yesCounts := map[rune]int{}
	numMembers := 0

	for _, line := range lines {
		if line == "" {
			for _, yesses := range yesCounts {
				if yesses == numMembers {
					total++
				}
			}

			yesCounts = map[rune]int{}
			numMembers = 0

			continue
		}

		for _, question := range line {
			yesses, _ := yesCounts[question]
			yesCounts[question] = yesses + 1
		}

		numMembers++
	}

	for _, yesses := range yesCounts {
		if yesses == numMembers {
			total++
		}
	}

	log.Printf("Total: %d", total)
}
