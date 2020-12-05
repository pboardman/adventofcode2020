package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func readFile(filename string) []string {
	dat, _ := ioutil.ReadFile(filename)
	txt := strings.Trim(string(dat), "\n")
	return strings.Split(txt, "\n")
}

func computeID(boardingPass string) int {
	var finalRow, finalCol int
	lowestRow, highestRow, lowestCol, highestCol := 0, 127, 0, 7
	for index, char := range boardingPass {
		// Compute row
		if index < 7 {
			switch char {
			case 'B':
				if index != 6 {
					lowestRow += (highestRow - lowestRow + 1) / 2
				} else {
					finalRow = highestRow
				}
			case 'F':
				if index != 6 {
					highestRow -= (highestRow - lowestRow + 1) / 2
				} else {
					finalRow = lowestRow
				}
			}
		} else {
			// Compute column
			switch char {
			case 'R':
				if index != 9 {
					lowestCol += (highestCol - lowestCol + 1) / 2
				} else {
					finalCol = highestCol
				}
			case 'L':
				if index != 9 {
					highestCol -= (highestCol - lowestCol + 1) / 2
				} else {
					finalCol = lowestCol
				}
			}
		}
	}

	return finalRow*8 + finalCol
}

func main() {
	boardingPasses := readFile("inputDay5.txt")

	var seatIDs []int

	highestID := 0
	lowestID := computeID(boardingPasses[0])
	for _, boardingPass := range boardingPasses {
		seatID := computeID(boardingPass)
		if seatID > highestID {
			highestID = seatID
		} else if seatID < lowestID {
			lowestID = seatID
		}

		seatIDs = append(seatIDs, seatID)
	}

	fmt.Printf("Part 1: %v\n", highestID)

	sort.Ints(seatIDs)
	for i := lowestID; i < highestID; i++ {
		if i != seatIDs[i-lowestID] {
			fmt.Printf("Part 2: %v\n", i)
			break
		}
	}
}
