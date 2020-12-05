package main

import (
	"log"
	"strconv"
	"strings"

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

	seats := [1024]int{}

	for _, line := range lines {
		s := strings.ReplaceAll(line, "F", "0")
		s = strings.ReplaceAll(s, "B", "1")
		s = strings.ReplaceAll(s, "L", "0")
		s = strings.ReplaceAll(s, "R", "1")

		v, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			log.Fatalf("error converting %s to number: %v", s, err)
		}

		seats[int(v)] = 1
	}

	for i := int(1); i < 1023; i++ {
		if seats[i] == 1 {
			// seat is filled
			continue
		}

		if seats[i-1] == 0 || seats[i+1] == 0 {
			// one of the surrounding seats is not filled, and we know we're sitting between two filled seats
			// so this unfilled seat can't be the right one.
			continue
		}

		log.Printf("Found empty seat %d surrounded by filled seats", i)
	}
}

