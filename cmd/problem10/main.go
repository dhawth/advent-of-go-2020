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

	emptySeats := map[int]struct{}{}

	for i := int(0); i < 1024; i++ {
		emptySeats[i] = struct{}{}
	}

	for _, s := range lines {
		s = strings.ReplaceAll(s, "F", "0")
		s = strings.ReplaceAll(s, "B", "1")
		s = strings.ReplaceAll(s, "L", "0")
		s = strings.ReplaceAll(s, "R", "1")

		v, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			log.Fatalf("error converting %s to number: %v", s, err)
		}

		delete(emptySeats, int(v))
	}

	for i := range emptySeats {
		_, beforeSeatIsEmpty := emptySeats[i-1]
		_, afterSeatIsEmpty := emptySeats[i+1]

		if beforeSeatIsEmpty || afterSeatIsEmpty {
			continue
		}

		log.Printf("Found empty seat %d surrounded by filled seats", i)
	}
}

