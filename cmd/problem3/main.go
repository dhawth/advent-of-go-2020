package main

import (
	"fmt"
	"log"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile = "input.txt"
)

type passwordEntry struct{
	min, max int
	c rune
	password string
}

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// 2-15 w: wlwwwlgmhwwgwwkwz means w must appear 2-15 times in the following string
	passwords := parsePasswords(lines)
	var correct int

	for _, password := range passwords {
		var count int
		for _, r := range password.password {
			if r == password.c {
				count++
			}
		}

		if count >= password.min && count <= password.max {
			correct++
		}
	}

	fmt.Printf("There were %d correct passwords\n", correct)
}

func parsePasswords(lines []string) []passwordEntry {
	var results []passwordEntry

	for _, line := range lines {
		// 2-15 w: wlwwwlgmhwwgwwkwz
		var p passwordEntry
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &p.min, &p.max, &p.c, &p.password)
		if err != nil {
			log.Fatalf("invalid input: %s: %v", line, err)
		}

		results = append(results, p)
	}

	return results
}
