package main

import (
	"io/ioutil"
	"strconv"
	"testing"
)

func TestPartOne(t *testing.T) {
	lines, err := LinesFromFile("test-input.txt")
	if err != nil {
		t.Fatalf("Error reading test-input.txt. %s", err)
	}
	output, err := ioutil.ReadFile("test-output.txt")
	if err != nil {
		t.Fatalf("Could not read test-output.txt. %s", err)
	}

	num, err := strconv.Atoi(string(output))
	if err != nil {
		t.Fatalf("Could not parse file as int %s", err)
	}

	result, err := Process(lines)

	if err != nil {
		t.Fatalf("Error during process. %s", err)
	}

	if num != result {
		t.Fatalf("%d != %d", num, result)
	}
}
