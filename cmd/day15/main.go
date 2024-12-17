package main

import (
	"fmt"
	"os"
	"strings"
)

type tile int
const (
	WALL tile = iota
	BOX
	EMPTY
	LEFTBOX
	RIGHTBOX
)

type direction int
const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

func parseInput(input string, lineEnding string, thickness int) ([][]tile, []direction, [2]int) {
	parts := strings.Split(input, lineEnding + lineEnding)
	layout := [][]tile{}
	robotPosition := [2]int{0, 0}

	for y, line := range strings.Split(parts[0], lineEnding) {
		row := []tile{}
		for x, symbol := range strings.Split(line, "") {
			switch symbol {
			case "#":
				for i := 0; i < thickness; i++ {
					row = append(row, WALL)
				}
			case "O":
				if thickness == 1 {
					row = append(row, BOX)
				} else {
					row = append(row, LEFTBOX)
					row = append(row, RIGHTBOX)
				}
			case ".":
				for i := 0; i < thickness; i++ {
					row = append(row, EMPTY)
				}
			case "@":
				for i := 0; i < thickness; i++ {
					row = append(row, EMPTY)
				}
				robotPosition = [2]int{y, x * thickness}
			}
		}
		layout = append(layout, row)
	}

	directions := []direction{}
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

func simulateRobot(input string, lineEnding string, thickness int) [][]tile {
	layout, directions, robotPosition := parseInput(input, lineEnding, thickness)
	
	directionOffsets := map[direction][2]int{
		UP: {-1, 0},
		RIGHT: {0, 1},
		DOWN: {1, 0},
		LEFT: {0, -1},
	}

	out:
	for _, direction := range directions {
		objectsToMove := []map[[2]int]bool{}
		objectsToMove = append(objectsToMove, map[[2]int]bool{})
		objectsToMove[len(objectsToMove) - 1][robotPosition] = true

		for ; len(objectsToMove[len(objectsToMove) - 1]) > 0; {
			objectsToMoveRow := map[[2]int]bool{}
			
			for object := range objectsToMove[len(objectsToMove) - 1] {
				positionToCheck := [2]int{
					object[0] + directionOffsets[direction][0],
					object[1] + directionOffsets[direction][1],
				}

				switch layout[positionToCheck[0]][positionToCheck[1]] {
				case WALL:
					continue out
				case LEFTBOX:
					objectsToMoveRow[positionToCheck] = true
					if direction == UP || direction == DOWN {
						objectsToMoveRow[[2]int{positionToCheck[0], positionToCheck[1] + 1}] = true
					}
				case RIGHTBOX:
					objectsToMoveRow[positionToCheck] = true
					if direction == UP || direction == DOWN {
						objectsToMoveRow[[2]int{positionToCheck[0], positionToCheck[1] - 1}] = true
					}
				case BOX:
					objectsToMoveRow[positionToCheck] = true
				}
			}

			objectsToMove = append(objectsToMove, objectsToMoveRow)
		}

		robotPosition = [2]int{robotPosition[0] + directionOffsets[direction][0], robotPosition[1] + directionOffsets[direction][1]}

		for i := len(objectsToMove) - 1; i > 0; i-- {
			for object := range objectsToMove[i] {
				layout[object[0] + directionOffsets[direction][0]][object[1] + directionOffsets[direction][1]] = layout[object[0]][object[1]]
				layout[object[0]][object[1]] = EMPTY
			}
		}
	}

	return layout
}

func sumOfBoxes(layout [][]tile, markerTile tile) int {
	sum := 0

	for y, row := range layout {
		for x, tile := range row {
			if tile == markerTile {
				sum += (100 * y) + x
			}
		}
	}

	return sum
}

func part1(input string, lineEnding string) int {
	layout := simulateRobot(input, lineEnding, 1)

	return sumOfBoxes(layout, BOX)
}

func part2(input string, lineEnding string) int {
	layout := simulateRobot(input, lineEnding, 2)

	return sumOfBoxes(layout, LEFTBOX)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)

	result := part1(input, "\r\n")
	fmt.Println("sum of all boxes' GPS coordinates is:", result)

	result = part2(input, "\r\n")
	fmt.Println("sum of all boxes' GPS coordinates is:", result)
}