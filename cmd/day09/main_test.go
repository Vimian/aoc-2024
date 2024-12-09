package main

import "testing"

var input string = `2333133121414131402`

var parsedInput []*int

func init() {
	parsedInput = parseInput(input, "\n")
}

func TestPart1(t *testing.T) {
	var expected int = 1928

	var result int = part1(parsedInput)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var expected int = 2858

	var result int = part2(parsedInput)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}