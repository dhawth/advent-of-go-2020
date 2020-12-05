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

	cards, err := lib.ConvertStringsToInts(lines)
	if err != nil {
		log.Fatal(err)
	}

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

