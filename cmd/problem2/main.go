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
	cards := getTheStuff()

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

func getTheStuff() []int {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	var cards []int

	for ; scanner.Scan(); {
		i := scanner.Text()
		n, err := strconv.Atoi(i)
		if err != nil {
			log.Fatalf("invalid number in input: %s: %v", i, err)
		}

		cards = append(cards, n)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return cards
}