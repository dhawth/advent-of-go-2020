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
	cards := convertToInts(lines)

	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			for k := j + 1; k < len(cards); k++ {
				if (cards[i] + cards[j] + cards[k]) == 2020 {
					fmt.Printf("%d * %d * %d = %d\n", cards[i], cards[j], cards[k],
						cards[i] * cards[j] * cards[k])
				}
			}
		}
	}
}

func convertToInts(lines []string) []int {
	var results []int
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("error converting %s to int: %v", line, err)
		}
		results = append(results, n)
	}

	return results
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
