package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

func main() {
	lines := getTheStuff()
	seats := map[int]struct{}{}

	for _, line := range lines {
		s := strings.ReplaceAll(line, "F", "0")
		s = strings.ReplaceAll(s, "B", "1")
		s = strings.ReplaceAll(s, "L", "0")
		s = strings.ReplaceAll(s, "R", "1")

		v, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			log.Fatalf("error converting %s to number: %v", s, err)
		}

		seats[int(v)] = struct{}{}
	}

	for i := int(1); i < 1023; i++ {
		_, hasSeat := seats[i]
		if hasSeat {
			continue
		}

		_, hasBeforeSeat := seats[i-1]
		_, hasAfterSeat := seats[i+1]

		if !hasBeforeSeat || !hasAfterSeat {
			continue
		}

		log.Printf("Found empty seat %d surrounded by filled seats", i)
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
