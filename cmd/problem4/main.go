package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	// 2-15 w: wlwwwlgmhwwgwwkwz means w must appear at either 2 or 15 but not at both
	passwords := getTheStuff()
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

func getTheStuff() []passwordEntry {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	var results []passwordEntry

	for ; scanner.Scan(); {
		line := scanner.Text()

		var p passwordEntry
		var passwdString string

		// 2-15 w: wlwwwlgmhwwgwwkwz
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &p.min, &p.max, &p.c, &passwdString)
		if err != nil {
			log.Fatalf("invalid input: %s: %v", line, err)
		}

		// convert string to array of runes to make indexing trivial later
		p.password = []rune(passwdString)

		// fix index offsets
		p.min -= 1
		p.max -= 1

		results = append(results, p)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return results
}