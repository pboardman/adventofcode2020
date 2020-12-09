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

func instructionInSlice(instructionsRan []int, instruction int) bool {
	for _, listInstruction := range instructionsRan {
		if instruction == listInstruction {
			return true
		}
	}
	return false
}

func part1(instructions []string) {
	acc := 0
	insPtr := 0

	var instructionsRan []int

	for {
		r := regexp.MustCompile(`([a-z]{3}) ([+-])(\d*)`)
		parsedIns := r.FindStringSubmatch(instructions[insPtr])
		ins := parsedIns[1]
		argSign := parsedIns[2]
		argValue, _ := strconv.Atoi(parsedIns[3])

		if instructionInSlice(instructionsRan, insPtr) {
			fmt.Printf("Part 1: %v\n", acc)
			return
		}
		instructionsRan = append(instructionsRan, insPtr)

		switch ins {
		case "jmp":
			if argSign == "+" {
				insPtr += argValue
			} else {
				insPtr -= argValue
			}
		case "acc":
			if argSign == "+" {
				acc += argValue
			} else {
				acc -= argValue
			}
			insPtr++
		case "nop":
			insPtr++
		}
	}
}

func part2(instructions []string) {

	var instructionsChanged []int

	for {
		acc := 0
		insPtr := 0

		var instructionsRan []int
		var instructionChanged bool
		for {
			r := regexp.MustCompile(`([a-z]{3}) ([+-])(\d*)`)
			parsedIns := r.FindStringSubmatch(instructions[insPtr])
			ins := parsedIns[1]
			argSign := parsedIns[2]
			argValue, _ := strconv.Atoi(parsedIns[3])

			// Change nop to jmp
			if !instructionChanged && ins != "acc" && !instructionInSlice(instructionsChanged, insPtr) {
				if ins == "jmp" {
					ins = "nop"
				} else {
					ins = "jmp"
				}
				instructionsChanged = append(instructionsChanged, insPtr)
				instructionChanged = true
			}

			// break if we loop
			if instructionInSlice(instructionsRan, insPtr) {
				break
			}
			instructionsRan = append(instructionsRan, insPtr)

			switch ins {
			case "jmp":
				if argSign == "+" {
					insPtr += argValue
				} else {
					insPtr -= argValue
				}
			case "acc":
				if argSign == "+" {
					acc += argValue
				} else {
					acc -= argValue
				}
				insPtr++
			case "nop":
				insPtr++
			}

			if insPtr == len(instructions) {
				fmt.Printf("Part 2: %v\n", acc)
				return
			}
		}
	}
}

func main() {
	instructions := readFile("inputday8.txt")

	part1(instructions)
	part2(instructions)
}
