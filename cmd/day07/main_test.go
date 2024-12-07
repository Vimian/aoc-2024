package main

import "testing"

var input string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

var equations []equation = parseInput(input, "\n")

func TestPart1(t *testing.T) {
	var expected int = 3749

	var result int = part1(equations)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var expected int = 11387

	var result int = part2(equations)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}