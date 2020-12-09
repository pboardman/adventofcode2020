package main

import (
	"fmt"
	"io/ioutil"
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

func isValid(number int, last25 []int) bool {
	for i, firstNum := range last25 {
		for y, secondNum := range last25 {
			if firstNum+secondNum == number && i != y {
				return true
			}
		}
	}
	return false
}

func part1(numbers []int) int {
	var invalidNumber int
	for index, number := range numbers {
		if index >= 25 {
			if !isValid(number, numbers[index-25:index]) {
				fmt.Printf("Part 1: %v\n", number)
				invalidNumber = number
				break
			}
		}
	}

	return invalidNumber
}

func addBiggestSmallestNum(sequenceList []int) int {
	smallestNumber := sequenceList[0]
	biggestNumber := 0

	for _, number := range sequenceList {
		if number < smallestNumber {
			smallestNumber = number
		}

		if number > biggestNumber {
			biggestNumber = number
		}
	}

	return smallestNumber + biggestNumber
}

func part2(numbers []int, invalidNumber int) {
	for index, number := range numbers {
		totalSum := 0
		sequenceList := []int{number}
		for _, nextNumber := range numbers[index:] {
			totalSum += nextNumber
			sequenceList = append(sequenceList, nextNumber)
			if totalSum > invalidNumber {
				break
			} else if totalSum == invalidNumber {
				fmt.Printf("Part 2: %v\n", addBiggestSmallestNum(sequenceList))
				return
			}
		}
	}
}

func main() {
	numbers := readFile("inputday9.txt")
	invalidNumber := part1(numbers)
	part2(numbers, invalidNumber)
}
