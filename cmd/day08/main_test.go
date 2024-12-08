package main

import "testing"

var input string = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

var parsedInput antennas
var dimension dimensions

func init() {
	parsedInput, dimension = parseInput(input, "\n")
}

func TestPart1(t *testing.T) {
	var expected int = 14

	var result int = part1(parsedInput, dimension)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var expected int = 34

	var result int = part2(parsedInput, dimension)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}