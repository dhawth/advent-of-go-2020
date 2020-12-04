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
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer func() {
		_ = f.Close()
	}()

	cards := map[int]int{}

	scanner := bufio.NewScanner(f)

	for ; scanner.Scan(); {
		i := scanner.Text()
		n, err := strconv.Atoi(i)
		if err != nil {
			log.Fatalf("invalid number in input: %s: %v", i, err)
		}

		if o, ok := cards[2020 - n]; ok {
			fmt.Printf("%d * %d = %d\n", o, n, o * n)
		}

		cards[n] = n
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("scanner error: %v", err)
	}
}

