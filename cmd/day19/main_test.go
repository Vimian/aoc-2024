package main

import "testing"

var input string = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

var towels []string
var designs []string

func init() {
	towels, designs = parseInput(input, "\n")
}

func TestPart1(t *testing.T) {
	var expected int = 6

	var result, _ = part1(towels, designs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var expected int = 16
	
	var result int = part2(towels, designs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}