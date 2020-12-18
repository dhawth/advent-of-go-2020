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
		log.Fatalf("error reading file: %s", err)
	}

	// bags is a map of any bag type to its list of children and their quantities
	bags := map[string]map[string]int{}

	for _, line := range lines {
		line = strings.TrimSuffix(line, ".")
		// split on contain, test for "no other bags" first, then split on comma
		fields := strings.Split(line, " contain ")
		parent := fields[0]
		remainder := fields[1]

		if remainder == "no other bags" {
			bags[parent] = map[string]int{}
			continue
		}

		others := strings.Split(remainder, ",")

		for _, other := range others {
			other = strings.TrimSpace(other)

			// normalize around plurals
			if strings.HasSuffix(other, "bag") {
				other = other + "s"
			}

			fields = strings.SplitN(other, " ", 2)

			if len(fields) != 2 {
				log.Fatalf("fields does not contain 2 fields: (%d) %v", len(fields), fields)
			}

			amount, err := strconv.Atoi(fields[0])
			if err != nil {
				log.Fatalf("could not convert %s to int on line '%s': %v", fields[0], line, err)
			}

			child := fields[1]

			// add child as a child of the parent
			children, ok := bags[parent]

			if !ok {
				children = map[string]int{}
			}
			children[child] = amount
			bags[parent] = children

			// make sure child has an entry even if it's an empty set since it may not have children of its own
			if _, ok = bags[child]; !ok {
				bags[child] = map[string]int{}
			}
		}
	}

	// note the lack of loop detection: we're really trusting our input here.
	queue := lib.NewStringQueue()
	numBags := 0

	queue.Push("shiny gold bags")

	for !queue.IsEmpty() {
		n := queue.Pop()

		children := bags[n]

		if len(children) == 0 {
			continue
		}

		for k, v := range children {
			numBags += v
			for i := 0; i < v; i++ {
				queue.Push(k)
			}
		}
	}

	log.Printf("There are %d total bags in the shiny gold bag", numBags)
}
