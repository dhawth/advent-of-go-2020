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
	password []rune
}

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// 2-15 w: wlwwwlgmhwwgwwkwz means w must appear at either 2 or 15 but not at both
	passwords := parsePasswords(lines)
	var correct int

	for _, p := range passwords {
		if p.password[p.min] == p.c || p.password[p.max] == p.c {
			if p.password[p.min] != p.password[p.max] {
				correct++
			}
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
