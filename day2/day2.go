package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	dat, _ := ioutil.ReadFile(filename)
	txt := strings.Trim(string(dat), "\n")
	return strings.Split(txt, "\n")
}

func letterInWord(letter string, word string) int {
	numberLetter := 0
	for _, l := range word {
		if letter == string(l) {
			numberLetter++
		}
	}
	return numberLetter
}

func letterAtPosition(letter string, position int, word string) bool {
	return string(word[position-1]) == letter
}

func part1() {
	values := readFile("inputday2.txt")

	goodPasswords := 0
	for _, line := range values {
		r := regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): (.*)`)
		regResult := r.FindStringSubmatch(line)

		lowerBound, _ := strconv.Atoi(regResult[1])
		upperBound, _ := strconv.Atoi(regResult[2])
		letter := regResult[3]
		password := regResult[4]
		nbLetter := letterInWord(letter, password)

		if nbLetter >= lowerBound && nbLetter <= upperBound {
			goodPasswords++
		}
	}

	fmt.Printf("Part 1: %v\n", goodPasswords)
}

func part2() {
	values := readFile("inputday2.txt")

	goodPasswords := 0
	for _, line := range values {
		r := regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): (.*)`)
		regResult := r.FindStringSubmatch(line)

		firstPos, _ := strconv.Atoi(regResult[1])
		secondPos, _ := strconv.Atoi(regResult[2])
		letter := regResult[3]
		password := regResult[4]

		if letterAtPosition(letter, firstPos, password) != letterAtPosition(letter, secondPos, password) {
			goodPasswords++
		}
	}

	fmt.Printf("Part 2: %v\n", goodPasswords)
}

func main() {
	part1()
	part2()
}
