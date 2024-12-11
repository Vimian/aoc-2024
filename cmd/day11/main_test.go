package main

import "testing"

var input string = `125 17`

var parsedInput []int

func init() {
	parsedInput = parseInput(input, "\n")
}

func TestPart1(t *testing.T) {
	var expected int = 55312

	var result int = part1(parsedInput, 25)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var expected int = 22

	var result int = part2(parsedInput, 6)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}