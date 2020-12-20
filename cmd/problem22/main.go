package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dhawth/advent-of-go-2020/lib"
)

const (
	inputFile  = "input.txt"
	chair      = "L"
	notASeat   = 0x0
	emptySeat  = 0x1
	filledSeat = 0x2
)

type seat struct {
	row, position int
}

func (s seat) Row() int {
	return s.row
}

func (s seat) Position() int {
	return s.position
}

func newSeat(row, position int) *seat {
	return &seat{
		row:      row,
		position: position,
	}
}

func (s seat) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("(%d, %d)", s.row, s.position))
}

func main() {
	lines, err := lib.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	grid := convertToGrid(lines)
	rowLength := len(grid[0])

	// list of adjacencies - this never changes and is a list of edges between nodes, indexed by (row, position)
	var adjacencies = make([][][]*seat, len(grid))
	var seatStatus = make([]byte, len(grid)*rowLength)

	for rowIndex, row := range grid {
		adjacencies[rowIndex] = make([][]*seat, len(row))

		for position := 0; position < len(row); position++ {
			if row[position] != chair {
				continue
			}

			seatStatus[rowIndex*rowLength+position] = emptySeat

			for p := position - 1; p >= 0; p-- {										// to the left
				if buildAdjacency(grid, rowIndex, position, rowIndex, p, adjacencies) {
					break
				}
			}

			for p := position + 1; p < len(row); p++ {									// to the right
				if buildAdjacency(grid, rowIndex, position, rowIndex, p, adjacencies) {
					break
				}
			}

			for r := rowIndex + 1; r < len(grid); r++ {								// same seat, straight down
				if buildAdjacency(grid, rowIndex, position, r, position, adjacencies) {
					break
				}
			}

			for r, p := rowIndex + 1, position - 1; r < len(grid) && p >= 0; {		// down and left diagonal
				if buildAdjacency(grid, rowIndex, position, r, p, adjacencies) {
					break
				}
				r++
				p--
			}

			for r, p := rowIndex + 1, position + 1; r < len(grid) && p < len(row); { // down and right diagonal
				if buildAdjacency(grid, rowIndex, position, r, p, adjacencies) {
					break
				}
				r++
				p++
			}

			for r := rowIndex - 1; r >= 0; r-- {								// same seat, straight up
				if buildAdjacency(grid, rowIndex, position, r, position, adjacencies) {
					break
				}
			}

			for r, p := rowIndex - 1, position - 1; r >= 0 && p >= 0; {		// up and left diagonal
				if buildAdjacency(grid, rowIndex, position, r, p, adjacencies) {
					break
				}
				r--
				p--
			}

			for r, p := rowIndex - 1, position + 1; r >= 0 && p < len(row); { // up and right diagonal
				if buildAdjacency(grid, rowIndex, position, r, p, adjacencies) {
					break
				}
				r--
				p++
			}
		}
	}

	// adjacencies is now filled with the _first_ seat (not floor) visible in each direction from the given position.
	// a seat with no adjacencies can not, by definition, see a filled seat because it can only see floor in all
	// directions.

	// a cycle fills all seats that are empty and have no adjacencies filled
	// and then removes people from seats that have at least 4 adjacent seats filled
	// repeat these cycles until the seat status array does not change either from filling or removing
	// if we use []byte for seatIsFilled, we can use bytes.Compare to quickly compare them

	var iterations int
	for {
		iterations++

		newSeatIsFilled := fill(seatStatus, rowLength, adjacencies)
		if bytes.Compare(seatStatus, newSeatIsFilled) == 0 {
			break
		}

		seatStatus = newSeatIsFilled

		newSeatIsFilled = relax(seatStatus, rowLength, adjacencies)
		if bytes.Compare(seatStatus, newSeatIsFilled) == 0 {
			break
		}

		seatStatus = newSeatIsFilled
	}

	var filledSeats int
	for _, v := range seatStatus {
		if v == filledSeat {
			filledSeats++
		}
	}

	log.Printf("finished after %d iterations with %d filled seats", iterations, filledSeats)
}

func printFilledSeatsArray(seats []byte, rowLength int) {
	for i, b := range seats {
		if i%rowLength == 0 {
			fmt.Println()
		}

		switch b {
		case emptySeat:
			fmt.Print("L")
		case filledSeat:
			fmt.Print("#")
		case notASeat:
			fmt.Print(".")
		}
	}
}

func fill(seatStatus []byte, rowLength int, adjacencies [][][]*seat) []byte {
	copyOfSeatIsFilled := make([]byte, len(seatStatus))
	copy(copyOfSeatIsFilled, seatStatus)

	for rowIndex, row := range adjacencies {
	NextSeat:
		for positionIndex, adjacencyList := range row {
			// adjacencyList will be nil if this is not a seat
			if adjacencyList == nil {
				continue
			}

			for _, adjacentSeat := range adjacencyList {
				if seatStatus[adjacentSeat.Row()*rowLength+adjacentSeat.Position()] == filledSeat {
					continue NextSeat
				}
			}

			// no adjacencies are filled, fill this adjacentSeat
			copyOfSeatIsFilled[rowIndex*rowLength+positionIndex] = filledSeat
		}
	}
	return copyOfSeatIsFilled
}

func relax(seatStatus []byte, rowLength int, adjacencies [][][]*seat) []byte {
	copyOfSeatIsFilled := make([]byte, len(seatStatus))
	copy(copyOfSeatIsFilled, seatStatus)

	for rowIndex, row := range adjacencies {
		for positionIndex, adjacencyList := range row {
			seatIndex := rowIndex*rowLength + positionIndex
			if seatStatus[seatIndex] == emptySeat {
				continue
			}

			var filledAdjacencies int
			for _, adjacentSeat := range adjacencyList {
				if seatStatus[adjacentSeat.Row()*rowLength+adjacentSeat.Position()] == filledSeat {
					filledAdjacencies++
				}
			}

			if filledAdjacencies > 4 {
				copyOfSeatIsFilled[seatIndex] = emptySeat
			}
		}
	}

	return copyOfSeatIsFilled
}

func buildAdjacency(grid [][]string, row, position, rowToTest, positionToTest int, adjacencies [][][]*seat) bool {
	if grid[rowToTest][positionToTest] == chair {
		adjacencies[row][position] = append(adjacencies[row][position], newSeat(rowToTest, positionToTest))
		return true
	}

	return false
}

func convertToGrid(lines []string) [][]string {
	var res [][]string
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}

	return res
}
