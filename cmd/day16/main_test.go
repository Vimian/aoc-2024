package main

import "testing"

func TestPart1(t *testing.T) {
	var input string = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`
	layout, start := parseInput(input, "\n")
	_, end := calculateCostsToEnd(layout, start)

	var expected int = 7036

	var result int = part1(end)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart1_2(t *testing.T) {
	var input string = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`
	layout, start := parseInput(input, "\n")
	_, end := calculateCostsToEnd(layout, start)

	var expected int = 11048

	var result int = part1(end)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	var input string = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`
	layout, start := parseInput(input, "\n")
	costs, end := calculateCostsToEnd(layout, start)

	var expected int = 45

	var result int = part2(costs, end)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2_1(t *testing.T) {
	var input string = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`
	layout, start := parseInput(input, "\n")
	costs, end := calculateCostsToEnd(layout, start)

	var expected int = 64

	var result int = part2(costs, end)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}