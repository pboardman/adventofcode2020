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
	return strings.Split(txt, "\n\n")
}

func isValidByr(value string) bool {
	intValue, _ := strconv.Atoi(value)
	if intValue < 1920 || intValue > 2002 {
		return false
	}
	return true
}

func isValidIyr(value string) bool {
	intValue, _ := strconv.Atoi(value)
	if intValue < 2010 || intValue > 2020 {
		return false
	}
	return true
}

func isValidEyr(value string) bool {
	intValue, _ := strconv.Atoi(value)
	if intValue < 2020 || intValue > 2030 {
		return false
	}
	return true
}

func isValidHgt(value string) bool {
	height, _ := strconv.Atoi(value[:len(value)-2])

	switch value[len(value)-2:] {
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	case "in":
		if height < 59 || height > 76 {
			return false
		}
	default:
		return false
	}
	return true
}

func isValidHcl(value string) bool {
	matched, _ := regexp.MatchString(`^#[A-Za-z0-9]{6}$`, value)
	if !matched {
		return false
	}
	return true
}

func isValidEcl(value string) bool {
	for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if value == color {
			return true
		}
	}
	return false
}

func isValidPid(value string) bool {
	matched, _ := regexp.MatchString(`^\d{9}$`, value)
	if !matched {
		return false
	}
	return true
}

func isValidField(field, value string) bool {

	type fn func(string) bool

	validations := map[string]fn{
		"byr": isValidByr,
		"iyr": isValidIyr,
		"eyr": isValidEyr,
		"hgt": isValidHgt,
		"hcl": isValidHcl,
		"ecl": isValidEcl,
		"pid": isValidPid,
	}

	return validations[field](value)
}

func splitField(s string) (string, string) {
	x := strings.Split(s, ":")
	return x[0], x[1]
}

func isValidPassportPart1(passport string) bool {
	for _, field := range []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"} {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func isValidPassportPart2(passport string) bool {
	// Split password on spaces and newlines
	fields := strings.FieldsFunc(passport, func(r rune) bool {
		return r == ' ' || r == '\n'
	})

	fieldMap := make(map[string]string)
	for _, field := range fields {
		fieldName, fieldValue := splitField(field)
		fieldMap[fieldName] = fieldValue
	}

	for _, field := range []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"} {
		if !strings.Contains(passport, field) || !isValidField(field[:len(field)-1], fieldMap[field[:len(field)-1]]) {
			return false
		}
	}
	return true
}

func part1() {
	values := readFile("inputday4.txt")

	validPassports := 0
	for _, passport := range values {
		if isValidPassportPart1(passport) {
			validPassports++
		}
	}

	fmt.Printf("Part 1: %v\n", validPassports)
}

func part2() {
	values := readFile("inputday4.txt")

	validPassports := 0
	for _, passport := range values {
		if isValidPassportPart2(passport) {
			validPassports++
		}
	}

	fmt.Printf("Part 2: %v\n", validPassports)
}

func main() {
	part1()
	part2()
}
