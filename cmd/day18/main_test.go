package main

import "testing"

var input string = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

var parsedInput []location

func init() {
	parsedInput = parseInput(input, "\n")
}

func TestPart1(t *testing.T) {
	var expected int = 22

	result := part1(parsedInput, 6, 12, location{x: 0, y: 0}, location{x: 6, y: 6})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var expected location = location{x: 6, y: 1}

	result := part2(parsedInput, 6, 12, location{x: 0, y: 0}, location{x: 6, y: 6})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}