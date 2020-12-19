package main

import (
	"log"
	"os"
	"sort"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile  = "cmd/problem18/input.txt"
	windowSize = 25
)

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	intQueue := lib.NewUniqueIntQueue()
	values, err := lib.ConvertStringsToInts(lines)
	if err != nil {
		log.Fatal(err)
	}

	// prepopulate the sliding window
	for i := 0; i < windowSize; i++ {
		line := values[i]
		intQueue.Push(line)
	}

	for i := windowSize; i < len(lines); i++ {
		v := values[i]
		var found bool

		// v is sum of 2 values of the sliding window... or not.
		for item := intQueue.Head(); item != nil; item = item.Next() {
			candidate := v - item.Value.(int)
			if intQueue.Contains(candidate) {
				found = true
				break
			}
		}

		if !found {
			log.Printf("could not find two numbers in the window that summed to %d at line %d", v, i)

			for start := 0; start < i; start++ {
				runningSum := 0

				for c := start; c < i; c++ {
					runningSum += values[c]

					if runningSum == v {
						answer := values[start:c]
						sort.Ints(answer)
						log.Printf("Answer: %d", answer[0] + answer[len(answer)-1])
						os.Exit(0)
					}

					if runningSum > v {
						break
					}
				}
			}
		}

		// advance the sliding window
		_, _ = intQueue.Pop()
		intQueue.Push(v)
	}
}
