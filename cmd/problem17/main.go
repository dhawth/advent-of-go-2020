package main

import (
	"log"
	"strconv"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile  = "input.txt"
	windowSize = 25
)

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	intQueue := lib.NewUniqueIntQueue()

	// prepopulate the sliding window
	for i := 0; i < windowSize; i++ {
		line := lines[i]
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not convert line %s to int: %v", line, err)
		}
		intQueue.Push(v)
	}

	for i := windowSize; i < len(lines); i++ {
		v, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatalf("could not convert line %s to int: %v", lines[i], err)
		}

		var found bool

		// v is sum of 2 values of the sliding window... or not.
		for item := intQueue.Head(); item != nil; item = item.Next() {
			candidate := v - item.Value.(int)
			if intQueue.Contains(candidate) {
				log.Printf("a sum for %d was found with %d and %d", v, candidate, item.Value.(int))
				found = true
				break
			}
		}

		if !found {
			log.Fatalf("could not find two numbers in the window that summed to %d", v)
		}

		// advance the sliding window
		_, _ = intQueue.Pop()
		intQueue.Push(v)
	}
}
