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
	password string
}

func main() {
	// 2-15 w: wlwwwlgmhwwgwwkwz means w must appear 2-15 times in the following string
	passwords := getTheStuff()
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
		// 2-15 w: wlwwwlgmhwwgwwkwz
		var p passwordEntry
		_, err := fmt.Sscanf(line, "%d-%d %c: %s", &p.min, &p.max, &p.c, &p.password)
		if err != nil {
			log.Fatalf("invalid input: %s: %v", line, err)
		}

		results = append(results, p)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return results
}