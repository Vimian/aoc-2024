package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string, lineEnding string) ([]int) {
	stones := []int{}

	for _, rawStone := range strings.Split(input, " ") {
		stone, err := strconv.Atoi(rawStone)
		if err != nil {
			panic(err)
		}

		stones = append(stones, stone)
	}

	return stones
}

func simulateChange(inputStones map[int]int) map[int]int {
	stones := map[int]int{}

	for stone, amount := range inputStones {
		if stone == 0 {
			stones[1] += amount
			continue
		}
		
		length := len(strconv.Itoa(stone))
		if length % 2 == 0 {
			var halfLength int = int(math.Pow(10, float64(length / 2)))
			stones[stone / halfLength] += amount
			stones[stone % halfLength] += amount
		} else {
			stones[stone * 2024] += amount
		}
	}

	return stones
}

func part1(inputStones []int, timesToBlink int) int {
	stones := map[int]int{}
	for _, stone := range inputStones {
		stones[stone]++
	}
	
	for i := 0; i < timesToBlink; i++ {
		stones = simulateChange(stones)
	}

	total := 0
	for _, amount := range stones {
		total += amount
	}

	return total
}

func part2(stones []int, timesToBlink int) int {
	return part1(stones, timesToBlink)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	stones := parseInput(input, "\r\n")

	result := part1(stones, 25)
	fmt.Println("amount of stones:", result)

	result = part2(stones, 75)
	fmt.Println("amount of stones:", result)
}