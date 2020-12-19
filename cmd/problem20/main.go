package main

import (
	"log"
	"sort"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile  = "input.txt"
)

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	values, err := lib.ConvertStringsToInts(lines)
	if err != nil {
		log.Fatal(err)
	}

	deviceJoltage := values[len(values) - 1] + 3

	// add the socket and device to the values list as the head and tail
	values = append(values, 0)
	values = append(values, deviceJoltage)

	sort.Ints(values)

	cache := map[int]int{}

	successes := doTheThingRight(values, len(values), 0, cache)

	log.Printf("successes: %d", successes)
}

func doTheThingRight(joltages []int, arrayLen, index int, visited map[int]int) int {
	if index >= arrayLen-3 {
		return 1
	}

	if res, ok := visited[index]; ok {
		return res
	}

	var recursiveSuccesses int
	for i := index + 1; i < index+4; i++ {
		if joltages[i] - joltages[index] <= 3 {
			recursiveSuccesses += doTheThingRight(joltages, arrayLen, i, visited)
		}
	}

	visited[index] = recursiveSuccesses
	return recursiveSuccesses
}

// brute-force method that obviously won't work, foundation for the actual answer
func recursivelyTest(values []int, arrayLen, index, previousJoltage int) int {
	if index == arrayLen {
		return 1 // success
	}

	if values[index] - 3 > previousJoltage {
		return 0 // failure
	}

	// return the number of successes from NOT using the adaptor at $index (because we don't add its joltage to the
	// previousJoltage, plus the number of successes from using the adapter at $index (because we do add its joltage
	// to the previousJoltage).
	return recursivelyTest(values, arrayLen, index + 1, previousJoltage) +
		   recursivelyTest(values, arrayLen, index + 1, previousJoltage + values[index])
}
