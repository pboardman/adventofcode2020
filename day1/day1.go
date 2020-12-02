package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	dat, _ := ioutil.ReadFile(filename)
	txt := string(dat)
	return strings.Split(txt, "\n")
}

func part1() {
	values := readFile("inputday1.txt")

	for i, firstNum := range values {
		firstNum, _ := strconv.Atoi(firstNum)
		for y, secondNum := range values {
			secondNum, _ := strconv.Atoi(secondNum)
			if firstNum+secondNum == 2020 && i != y {
				fmt.Printf("Part 1: %v\n", firstNum*secondNum)
				return
			}
		}
	}
}

func part2() {
	values := readFile("inputday1.txt")

	for i := 0; i < len(values); i++ {
		firstNum, _ := strconv.Atoi(values[i])
		for j := 1; j < len(values); j++ {
			secondNum, _ := strconv.Atoi(values[j])
			for k := 0; k < len(values); k++ {
				thirdNum, _ := strconv.Atoi(values[k])
				if firstNum+secondNum+thirdNum == 2020 {
					fmt.Printf("part 2: %v\n", firstNum*secondNum*thirdNum)
					return
				}
			}
		}
	}
}

func main() {
	part1()
	part2()
}
