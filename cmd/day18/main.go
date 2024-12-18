package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type location struct {
	x int
	y int
}

func parseInput(input string, lineEnding string) ([]location) {
	parsedInput := []location{}

	for _, line := range strings.Split(input, lineEnding) {
		

		numbers := strings.Split(line, ",")
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		parsedInput = append(parsedInput, location{x: x, y: y})
	}

	return parsedInput
}

func part1(locations []location, size int, corruptedLocations int, startLocation location, endLocation location) int {
	grid := make([][]int, size + 1)
	for i := range grid {
		grid[i] = make([]int, size + 1)
	}

	for i := 0; i < corruptedLocations; i++ {
		grid[locations[i].y][locations[i].x] = 1
	}

	takenMoves := map[location]int{}
	validMoves := map[location]int{startLocation: 0}
	for i := 0; len(validMoves) > 0; i++ {
		var moveToExplore location
		for move := range validMoves {
			if moveToExplore == (location{}) || validMoves[move] < validMoves[moveToExplore] {
				moveToExplore = move
			}
		}

		if moveToExplore == endLocation {
			return validMoves[moveToExplore]
		}

		for _, direction := range []location{{x: 0, y: 1}, {x: 0, y: -1}, {x: 1, y: 0}, {x: -1, y: 0}} {
			newLocation := location{x: moveToExplore.x + direction.x, y: moveToExplore.y + direction.y}

			if newLocation.x < 0 ||
				newLocation.x > size ||
				newLocation.y < 0 ||
				newLocation.y > size ||
				grid[newLocation.y][newLocation.x] == 1 { continue }

			if _, found := takenMoves[newLocation]; found { continue }

			if _, found := validMoves[newLocation]; found { continue }

			validMoves[newLocation] = validMoves[moveToExplore] + 1
		}

		takenMoves[moveToExplore] = validMoves[moveToExplore]
		delete(validMoves, moveToExplore)
	}

	return -1
}

func part2(locations []location, size int, corruptedLocations int, startLocation location, endLocation location) location {
	for i := corruptedLocations; i < size * size; i++ {
		result := part1(locations, size, i, startLocation, endLocation)
		if result != -1 {
			continue
		}

		return locations[i - 1]
	}

	return location{}
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	parsedInput := parseInput(input, "\r\n")

	result := part1(parsedInput, 70, 1024, location{x: 0, y: 0}, location{x: 70, y: 70})
	fmt.Println("minimum number of steps needed to reach the exit:", result)

	coordinate := part2(parsedInput, 70, 1024, location{x: 0, y: 0}, location{x: 70, y: 70})
	fmt.Println("coordinates of the first byte that will prevent the exit from being reachable:", coordinate.x, ",", coordinate.y)
}