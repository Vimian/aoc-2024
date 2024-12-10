package main

import "testing"

var input string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

var parsedInput [][]int

func init() {
	parsedInput = parseInput(input, "\n")
}

func TestPart1(t *testing.T) {
	var expected int = 36

	var result int = part1(parsedInput)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var expected int = 81

	var result int = part2(parsedInput)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}