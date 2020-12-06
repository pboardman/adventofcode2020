package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filename string) []string {
	dat, _ := ioutil.ReadFile(filename)
	txt := strings.Trim(string(dat), "\n")
	return strings.Split(txt, "\n\n")
}

func nbAnswerGroupPart1(groupAnswers string) int {
	answeredYes := make(map[rune]int)

	for _, answerLine := range strings.Split(groupAnswers, "\n") {
		for _, answer := range answerLine {
			answeredYes[answer]++
		}
	}

	return len(answeredYes)
}

func nbAnswerGroupPart2(groupAnswers string) int {
	answeredYes := make(map[rune]int)
	groupSize := len(strings.Split(groupAnswers, "\n"))

	for _, answerLine := range strings.Split(groupAnswers, "\n") {
		for _, answer := range answerLine {
			answeredYes[answer]++
		}
	}

	allAnsweredYes := 0
	for _, nbAnswer := range answeredYes {
		if nbAnswer == groupSize {
			allAnsweredYes++
		}
	}

	return allAnsweredYes
}

func part1(groups []string) {
	totalAnswer := 0

	for _, groupAnswers := range groups {
		totalAnswer += nbAnswerGroupPart1(groupAnswers)
	}

	fmt.Printf("Part 1: %v\n", totalAnswer)
}

func part2(groups []string) {
	totalAnswer := 0

	for _, groupAnswers := range groups {
		totalAnswer += nbAnswerGroupPart2(groupAnswers)
	}

	fmt.Printf("Part 2: %v\n", totalAnswer)
}

func main() {
	groups := readFile("inputday6.txt")

	part1(groups)
	part2(groups)
}
