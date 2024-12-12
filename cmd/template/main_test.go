package main

import "testing"

var input string = ``

var parsedInput 

func init() {
	parsedInput = parseInput(input, "\n")
}

func TestPart1(t *testing.T) {
	var expected int = -1

	var result int = part1(parsedInput)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var expected int = -1

	var result int = part2(parsedInput)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}