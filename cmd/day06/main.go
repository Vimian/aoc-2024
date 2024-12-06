package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) ([][]bool, [2]int) {
	obstructions := [][]bool{}
	position := [2]int{}

	for y, line := range strings.Split(input, "\r\n") {
		obstructionsInLine := []bool{}

		for x, char := range strings.Split(line, "") {
			if char == "." {
				obstructionsInLine = append(obstructionsInLine, false)
			} else if char == "#" {
				obstructionsInLine = append(obstructionsInLine, true)
			} else {
				obstructionsInLine = append(obstructionsInLine, false)
				position = [2]int{y, x}
			}
		}

		obstructions = append(obstructions, obstructionsInLine)
	}

	return obstructions, position
}

func part1(obstructions [][]bool, position [2]int) {
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	direction := 0

	distinctPositions := map[[2]int]bool{}
	
	for true {
		if obstructions[position[0] + directions[direction][0]][position[1] + directions[direction][1]] {
			direction = (direction + 1) % 4
		}

		position = [2]int{position[0] + directions[direction][0], position[1] + directions[direction][1]}
		distinctPositions[position] = true

		if position[0] == 0 || position[0] == len(obstructions) - 1 || position[1] == 0 || position[1] == len(obstructions[0]) - 1 {
			break
		}
	}

	fmt.Printf("visted %d distinct positions\n", len(distinctPositions))
}

func part2(obstructions [][]bool, startPosition [2]int) {
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	startDirection := 0

	posibleLocations := map[[2]int]bool{}

	currentPosition := [2]int{startPosition[0], startPosition[1]}
	currentDirection := startDirection
	
	for true {
		if obstructions[currentPosition[0] + directions[currentDirection][0]][currentPosition[1] + directions[currentDirection][1]] {
			currentDirection = (currentDirection + 1) % 4
			continue
		}

		tempObstruction := [2]int{currentPosition[0] + directions[currentDirection][0], currentPosition[1] + directions[currentDirection][1]}

		position := [2]int{startPosition[0], startPosition[1]}
		direction := startDirection
		tempDistinctPositions := map[[2]int]int{position: direction}

		for true {
			step := [2]int{position[0] + directions[direction][0], position[1] + directions[direction][1]}
			if obstructions[step[0]][step[1]] || step == tempObstruction {
				direction = (direction + 1) % 4
				continue
			}

			position = step

			_, found := tempDistinctPositions[position]
			if found && tempDistinctPositions[position] == direction {
				posibleLocations[tempObstruction] = true
				break
			}

			tempDistinctPositions[position] = direction
	
			if position[0] == 0 || position[0] == len(obstructions) - 1 || position[1] == 0 || position[1] == len(obstructions[0]) - 1 {
				break
			}
		}

		currentPosition = tempObstruction

		if currentPosition[0] == 0 || currentPosition[0] == len(obstructions) - 1 || currentPosition[1] == 0 || currentPosition[1] == len(obstructions[0]) - 1 {
			break
		}
	}

	fmt.Printf("%d posible positions\n", len(posibleLocations))
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	obstructions, position := parseInput(input)

	part1(obstructions, position)

	part2(obstructions, position)
}