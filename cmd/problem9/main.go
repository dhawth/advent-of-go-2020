package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile = "input.txt"
)

/*
There are 128 rows on the plane, indexed 0 through 127 and identified by a coded binary number representation:
FBFBBFF translates to...

0101100 = 4 + 8 + 32 = 44 ! success

There are 8 seats per row indexed 0 through 7 and identified by a coded binary number representation:

RLR = 5
101 = 4 + 1 = 5

the Seat ID = 8 * $row + $seat.  8 * 44 + 5 = 357

Find the highest seatID.

Strings are FBFBBFFRLR

Convert all F, B, L, R to 1s and 0s accordingly, and then cast from bin to dec.  Find max.
*/

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	maxSeatID := int64(0)

	for _, line := range lines {
		s := strings.ReplaceAll(line, "F", "0")
		s = strings.ReplaceAll(s, "B", "1")
		s = strings.ReplaceAll(s, "L", "0")
		s = strings.ReplaceAll(s, "R", "1")

		v, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			log.Fatalf("error converting %s to number: %v", s, err)
		}

		if v > maxSeatID {
			maxSeatID = v
		}
	}

	fmt.Printf("The maximum seatID is %d\n", maxSeatID)
}

