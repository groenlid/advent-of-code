package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partOne(numbers []int, expectedSum int) {
	for index, value := range numbers {
		for index2, value2 := range numbers {
			if index2 != index && (value+value2 == expectedSum) {
				log.Printf("Found the numbers for part1! %v + %v = %v. %v * %v = %v", value, value2, expectedSum, value, value2, value*value2)
				return
			}
		}
	}

	log.Fatalf("Could not find the correct numbers that sums up to %v", expectedSum)
}

func partTwo(numbers []int, expectedSum int) {
	for index, value := range numbers {
		for index2, value2 := range numbers {
			for index3, value3 := range numbers {
				if index2 != index && index2 != index3 && index != index3 && (value+value2+value3 == expectedSum) {
					log.Printf("Found the numbers for part3! %v + %v + %v = %v. %v * %v * %v = %v", value, value2, value3, expectedSum, value, value2, value3, value*value2*value3)
					return
				}
			}
		}
	}

	log.Fatalf("Could not find the correct numbers that sums up to %v", expectedSum)
}

func main() {
	log.Println("Starting day01")
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	csv := strings.Split(string(content), "\n")
	numbers := make([]int, len(csv))

	for index, value := range csv {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		numbers[index] = parsed
	}

	expectedSum := 2020
	partOne(numbers, expectedSum)
	partTwo(numbers, expectedSum)
}
