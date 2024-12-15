package main

import (
	"fmt"
	"os"
	"strings"
)

type Tiles int
const (
	WALL Tiles = iota
	BOX
	EMPTY
)

type Direction int
const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func parseInput(input string, lineEnding string) ([][]Tiles, []Direction, [2]int) {
	parts := strings.Split(input, lineEnding + lineEnding)
	layout := [][]Tiles{}
	robotPosition := [2]int{0, 0}

	for y, line := range strings.Split(parts[0], lineEnding) {
		row := []Tiles{}
		for x, symbol := range strings.Split(line, "") {
			switch symbol {
			case "#":
				row = append(row, WALL)
			case "O":
				row = append(row, BOX)
			case ".":
				row = append(row, EMPTY)
			case "@":
				row = append(row, EMPTY)
				robotPosition = [2]int{y, x}
			}
		}
		layout = append(layout, row)
	}

	directions := []Direction{}
	for _, symbol := range strings.Split(
		strings.ReplaceAll(parts[1], lineEnding, ""),
		"") {
		switch symbol {
		case "^":
			directions = append(directions, UP)
		case ">":
			directions = append(directions, RIGHT)
		case "v":
			directions = append(directions, DOWN)
		case "<":
			directions = append(directions, LEFT)
		}
	}

	return layout, directions, robotPosition
}

func part1(inputLayout [][]Tiles, directions []Direction, robotPosition [2]int) int {
	layout := make([][]Tiles, len(inputLayout))
	copy(layout, inputLayout)

	directionOffsets := map[Direction][2]int{
		UP: {-1, 0},
		RIGHT: {0, 1},
		DOWN: {1, 0},
		LEFT: {0, -1},
	}

	out:
	for _, direction := range directions {
		positionToCheck := [2]int{
			robotPosition[0] + directionOffsets[direction][0],
			robotPosition[1] + directionOffsets[direction][1],
		}
		for i := 0; positionToCheck[0] > 0 && positionToCheck[0] < len(layout) && positionToCheck[1] > 0 && positionToCheck[1] < len(layout[0]); i++ {
			switch layout[positionToCheck[0]][positionToCheck[1]] {
			case EMPTY:
				robotPosition = [2]int{robotPosition[0] + directionOffsets[direction][0], robotPosition[1] + directionOffsets[direction][1]}
				layout[positionToCheck[0]][positionToCheck[1]] = layout[robotPosition[0]][robotPosition[1]]
				layout[robotPosition[0]][robotPosition[1]] = EMPTY
				continue out
			case WALL:
				continue out
			}
			positionToCheck = [2]int{
				positionToCheck[0] + directionOffsets[direction][0],
				positionToCheck[1] + directionOffsets[direction][1],
			}
		}
	}

	sum := 0

	for y, row := range layout {
		for x, tile := range row {
			if tile == BOX {
				sum += (100 * y) + x
			}
		}
	}

	return sum
}

func part2(layout [][]Tiles, directions []Direction, robotPosition [2]int) int {
	return 0
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	layout, directions, robotPosition := parseInput(input, "\r\n")

	result := part1(layout, directions, robotPosition)
	fmt.Println("sum of all boxes' GPS coordinates:", result)

	result = part2(layout, directions, robotPosition)
	fmt.Println("checksum is:", result)
}