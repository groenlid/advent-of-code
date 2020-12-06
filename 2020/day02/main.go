package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func convertStringsToNumbers(strings []string) []int {
	numbers := make([]int, len(strings))

	for index, value := range strings {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		numbers[index] = parsed
	}
	return numbers
}

type Row struct {
	password       string
	charToEvaluate rune
	ruleNumbers    []int
}

func parseRow(rowContent string) (*Row, error) {

	delimited := strings.Split(strings.ReplaceAll(rowContent, ":", ""), " ")
	if len(delimited) != 3 {
		return nil, errors.New("Row did not have the correct format: " + rowContent)
	}

	rules := delimited[0]
	chars := []rune(delimited[1])
	password := delimited[2]

	if len(chars) != 1 {
		return nil, errors.New("Could not fetch a single char from the row: " + rowContent)
	}

	ruleNumbers := convertStringsToNumbers(strings.Split(rules, "-"))

	if len(ruleNumbers) != 2 {
		return nil, errors.New("Could not pase the rule numbers from the row: " + rowContent)
	}

	return &Row{
		password:       password,
		charToEvaluate: chars[0],
		ruleNumbers:    ruleNumbers,
	}, nil
}

func passwordIsValidForDayOne(password string, charToEvaluate rune, minInstances int, maxInstances int) bool {
	instances := 0
	for _, char := range password {
		if char == charToEvaluate {
			instances++
			if instances > maxInstances {
				return false
			}
		}
	}
	return instances >= minInstances
}

func passwordIsValidForDayTwo(password string, charToEvaluate rune, positions []int) bool {
	numberOfHits := 0
	passwordLength := len(password)

	for _, position := range positions {
		if passwordLength < position {
			continue
		}
		if password[position-1] == byte(charToEvaluate) {
			numberOfHits++
		}
	}
	return numberOfHits == 1
}

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(content), "\n")

	validPasswordsForDayOne := 0
	validPasswordsForDayTwo := 0
	for _, row := range rows {
		parsedRow, error := parseRow(row)
		if error != nil {
			log.Fatal(error)
		}

		if passwordIsValidForDayOne(parsedRow.password, parsedRow.charToEvaluate, parsedRow.ruleNumbers[0], parsedRow.ruleNumbers[1]) {
			validPasswordsForDayOne++
		}

		if passwordIsValidForDayTwo(parsedRow.password, parsedRow.charToEvaluate, parsedRow.ruleNumbers) {
			validPasswordsForDayTwo++
		}
	}

	log.Printf("Number of valid passwords for day one: %v", validPasswordsForDayOne)
	log.Printf("Number of valid passwords for day two: %v", validPasswordsForDayTwo)

}
