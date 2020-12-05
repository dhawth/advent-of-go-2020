package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	inputFile = "input.txt"
)

func main() {
	lines := getTheStuff()
	cards := map[int]int{}

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("invalid number in input: %s: %v", i, err)
		}

		if o, ok := cards[2020 - n]; ok {
			fmt.Printf("%d * %d = %d\n", o, n, o * n)
		}

		cards[n] = n
	}
}

func getTheStuff() []string {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	var results []string

	for ; scanner.Scan(); {
		results = append(results, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return results
}
