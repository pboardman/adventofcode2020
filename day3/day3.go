package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filename string) []string {
	dat, _ := ioutil.ReadFile(filename)
	txt := strings.Trim(string(dat), "\n")
	return strings.Split(txt, "\n")
}

func treeAtPos(line string, position int) int {
	if string(line[position%len(line)]) == "#" {
		return 1
	}

	return 0
}

func part1() {
	values := readFile("inputday3.txt")

	trees := 0
	for index, line := range values {
		trees += treeAtPos(line, index*3)
	}

	fmt.Printf("Part 1: %v\n", trees)
}

func part2() {
	values := readFile("inputday3.txt")

	slope1 := 0
	slope2 := 0
	slope3 := 0
	slope4 := 0
	slope5 := 0
	for index, line := range values {
		slope1 += treeAtPos(line, index*1)
		slope2 += treeAtPos(line, index*3)
		slope3 += treeAtPos(line, index*5)
		slope4 += treeAtPos(line, index*7)

		if index%2 == 0 {
			slope5 += treeAtPos(line, (index/2)*1)
		}
	}

	fmt.Printf("Part 2: %v\n", slope1*slope2*slope3*slope4*slope5)
}

func main() {
	part1()
	part2()
}
