package main

import (
	"log"
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

	// bags is a map of any bag type to its list of parents
	bags := map[string]lib.StringSet{}
	bagsThatContainNoOtherBags := lib.NewStringSet()
	bagsThatCanContainOtherBags := lib.NewStringSet()

	for _, line := range lines {
		line = strings.TrimSuffix(line, ".")
		// split on contain, test for "no other bags" first, then split on comma
		fields := strings.Split(line, " contain")
		parent := fields[0]
		remainder := fields[1]

		if remainder == "no other bags" {
			if bagsThatCanContainOtherBags.Contains(parent) {
				log.Fatalf("consistency error: %s should not contain other bags but has been seen as a parent previously", parent)
			}
			bagsThatContainNoOtherBags.Add(parent)
			continue
		}

		if bagsThatContainNoOtherBags.Contains(parent) {
			log.Fatalf("bag %s should not have any children", parent)
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

			// fields[0] is just the quantity and we don't care about it
			child := fields[1]

			// add parent as a parent of child
			parents, ok := bags[child]

			if !ok {
				parents = lib.NewStringSet()
			}
			parents.Add(parent)
			bags[child] = parents

			// make sure parent has an entry even if it's an empty set since it may not have a parent of its own
			if _, ok = bags[parent]; !ok {
				bags[parent] = lib.NewStringSet()
			}
		}
	}

	visited := lib.NewStringSet()
	queue := lib.NewUniqueStringQueue()

	queue.Push("shiny gold bags")

	for !queue.IsEmpty() {
		n := queue.Pop()

		parentBags := bags[n]

		if parentBags.IsEmpty() {
			continue
		}

		for _, p := range parentBags.Members() {
			if !visited.Contains(p) {
				visited.Add(p)
				queue.Push(p)
			}
		}
	}

	log.Printf("There are %d parents of the shiny gold bag", visited.Len())
}
