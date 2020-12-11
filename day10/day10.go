package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) []int {
	dat, _ := ioutil.ReadFile(filename)
	txt := strings.Trim(string(dat), "\n")
	var numbers []int
	for _, number := range strings.Split(txt, "\n") {
		convertedNum, _ := strconv.Atoi(number)
		numbers = append(numbers, convertedNum)
	}
	return numbers
}

func part1(jolts []int) {
	jolt1Diff := 0
	jolt3Diff := 0

	for index, jolt := range jolts {
		if index != len(jolts)-1 {
			switch diff := jolts[index+1] - jolt; diff {
			case 1:
				jolt1Diff++
			case 2:
			case 3:
				jolt3Diff++
			default:
				fmt.Println("wtf")
			}
		}
	}
	fmt.Printf("Part 1: %v\n", jolt1Diff*jolt3Diff)
}

func main() {
	jolts := readFile("inputday10.txt")
	sort.Ints(jolts)
	// Add 0 (the port) to start of slice
	jolts = append([]int{0}, jolts...)
	// Add device adapter (highest jolt + 3)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	part1(jolts)

}
