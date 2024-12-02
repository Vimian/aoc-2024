package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadInput(input string) [][]int {
	lines := strings.Split(input, "\r\n")
	
	levels := [][]int{}

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		
		levels = append(levels, []int{})
		levelsLast := len(levels) - 1

		for _, number := range numbers {
			level, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			levels[levelsLast] = append(levels[levelsLast], level)
		}
	}

	return levels
}

func isSafe(levels []int) bool {
	increases := false
	decreases := false

	for i := 1; i < len(levels); i++ {
		if levels[i] == levels[i-1] {
			return false
		} else if levels[i-1] < levels[i] {
			if levels[i] - levels[i-1] > 3 {
				return false
			}
			increases = true
		} else {
			if levels[i-1] - levels[i] > 3 {
				return false
			}
			decreases = true
		}

		if increases && decreases {
			return false
		}
	}

	return true
}

func part1(allLevels [][]int) {
	safeLevels := 0

	for _, levels := range allLevels {
		if isSafe(levels) {
			safeLevels++
		}
	}

	fmt.Printf("%d levels are safe\n", safeLevels)
}

func part2(allLevels [][]int) {
	safeLevels := 0

	for _, levels := range allLevels {
		if isSafe(levels) {
			safeLevels++
		} else {
			for i := 0; i < len(levels); i++ {
				tempLevels := make([]int, len(levels))
				copy(tempLevels, levels)
				if isSafe(append(tempLevels[:i], tempLevels[i+1:]...)) {
					safeLevels++
					break
				}
			}
		}
	}

	fmt.Printf("%d levels are safe\n", safeLevels)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)

	levels := loadInput(input)

	part1(levels)
	part2(levels)
}