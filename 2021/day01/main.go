package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func main() {
	lines, err := LinesFromFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input.txt. %s", err)
	}

	result, err := Process(lines)
	if err != nil {
		log.Fatalf("Error during process. %s", err)
	}

	log.Println(result)
}

func Process(numbers []string) (int, error) {
	if len(numbers) <= 1 {
		return 0, errors.New("not enough numbers given")
	}

	increments := 0

	for i := 1; i < len(numbers); i++ {
		prev, errPrev := strconv.Atoi(numbers[i-1])
		if errPrev != nil {
			return 0, errPrev
		}
		cur, errCur := strconv.Atoi(numbers[i])
		if errCur != nil {
			return 0, errPrev
		}
		if prev < cur {
			increments++
		}
	}

	return increments, nil
}

func LinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	result := []string{}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
