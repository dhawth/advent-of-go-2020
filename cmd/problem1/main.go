package main

import (
	"fmt"
	"log"

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

	numbers, err := lib.ConvertStringsToInts(lines)
	if err != nil {
		log.Fatal(err)
	}
	cards := map[int]int{}

	for _, n := range numbers {
		if o, ok := cards[2020 - n]; ok {
			fmt.Printf("%d * %d = %d\n", o, n, o * n)
		}

		cards[n] = n
	}
}

