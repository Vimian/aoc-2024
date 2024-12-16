package main

import (
	"fmt"
	"os"
	"strings"
)

type tile int
const (
	EMPTY tile = iota
	WALL
	END
)

type direction int
const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

type position struct {
	coordinates [2]int
	direction direction
}

func parseInput(input string, lineEnding string) ([][]tile, position) {
	layout := [][]tile{}
	startPosition := position{}

	for _, line := range strings.Split(input, lineEnding) {
		row := []tile{}

		for _, symbol := range strings.Split(line, "") {
			switch symbol {
			case ".":
				row = append(row, EMPTY)
			case "#":
				row = append(row, WALL)
			case "S":
				startPosition = position{[2]int{len(layout), len(row)}, RIGHT}
				row = append(row, EMPTY)
			case "E":
				row = append(row, END)
			}
			
		}

		layout = append(layout, row)
	}

	return layout, startPosition
}

type state struct {
	previous map[position]bool
	cost int
}

var directionOffsets map[direction][2]int = map[direction][2]int{
	UP: {-1, 0},
	RIGHT: {0, 1},
	DOWN: {1, 0},
	LEFT: {0, -1},
}

func calculateValidMoves(layout [][]tile, costs *map[position]state, validMoves *map[position]bool, moveToExplore position) {
	y := moveToExplore.coordinates[0]
	x := moveToExplore.coordinates[1]
	if layout[y][x] == END {
		return
	}

	forwardY := y + directionOffsets[moveToExplore.direction][0]
	forwardX := x + directionOffsets[moveToExplore.direction][1]

	if forwardY >= 0 && forwardY < len(layout) && forwardX >= 0 && forwardX < len(layout[0]) && layout[forwardY][forwardX] != WALL {
		forwardPosition := position{
			coordinates: [2]int{forwardY, forwardX},
			direction: moveToExplore.direction}

		if _, found := (*costs)[forwardPosition]; !found {
			(*costs)[forwardPosition] = state{
				previous: map[position]bool{moveToExplore: true},
				cost: (*costs)[moveToExplore].cost + 1}
			(*validMoves)[forwardPosition] = true
		} else if (*costs)[moveToExplore].cost + 1 == (*costs)[forwardPosition].cost {
			(*costs)[forwardPosition].previous[moveToExplore] = true
		}
	}

	rightPosition := position{
		coordinates: [2]int{y, x},
		direction: (moveToExplore.direction + 1) % 4}

	if _, found := (*costs)[rightPosition]; !found {
		(*costs)[rightPosition] = state{
			previous: map[position]bool{moveToExplore: true},
			cost: (*costs)[moveToExplore].cost + 1000}
		(*validMoves)[rightPosition] = true
	}

	leftPosition := position{
		coordinates: [2]int{y, x},
		direction: (moveToExplore.direction + 3) % 4}
	
	if _, found := (*costs)[leftPosition]; !found {
		(*costs)[leftPosition] = state{
			previous: map[position]bool{moveToExplore: true},
			cost: (*costs)[moveToExplore].cost + 1000}		 
		(*validMoves)[leftPosition] = true
	}

	delete((*validMoves), moveToExplore)

	nextToExplore := position{}
	for move, _ := range (*validMoves) {
		if (nextToExplore == position{} || (*costs)[move].cost < (*costs)[nextToExplore].cost) {
			nextToExplore = move
		}
	}
	
	calculateValidMoves(layout, costs, validMoves, nextToExplore)
}

func calculateCostsToEnd(layout [][]tile, start position) (map[position]state, state) {
	costs := map[position]state{}
	costs[start] = state{previous: map[position]bool{}, cost: 0}
	validMoves := map[position]bool{start: true}

	calculateValidMoves(layout, &costs, &validMoves, start)

	endCoordinates := [2]int{}
	for y, row := range layout {
		for x, tile := range row {
			if tile == END {
				endCoordinates = [2]int{y, x}
			}
		}
	}

	end := state{}

	if _, found := costs[position{endCoordinates, UP}]; found {
		end = costs[position{endCoordinates, UP}]
	} else if _, found := costs[position{endCoordinates, RIGHT}]; found {
		end = costs[position{endCoordinates, RIGHT}]
	} else if _, found := costs[position{endCoordinates, DOWN}]; found {
		end = costs[position{endCoordinates, DOWN}]
	} else if _, found := costs[position{endCoordinates, LEFT}]; found {
		end = costs[position{endCoordinates, LEFT}]
	}

	return costs, end
}

func part1(end state) int {
	return end.cost
}

func part2(costs map[position]state, end state) int {
	previousMoves := getPreviousMoves(&costs, end)
	
	return len(previousMoves) + 1
}

func getPreviousMoves(costs *map[position]state, state state) map[[2]int]bool {
	previousMoves := map[[2]int]bool{}

	for move := range state.previous {
		previousMoves[move.coordinates] = true
		tempMoves := getPreviousMoves(costs, (*costs)[move])
		for key := range tempMoves {
			previousMoves[key] = true
		}
	}

	return previousMoves
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	layout, start := parseInput(input, "\r\n")

	costs, end := calculateCostsToEnd(layout, start)

	result := part1(end)
	fmt.Println("lowest score is:", result)

	result = part2(costs, end)
	fmt.Println("checksum is:", result)
}