package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string, lineEnding string) ([][]int) {
	topographicMap := [][]int{}

	for _, lines := range strings.Split(input, lineEnding) {
		topographicLine := []int{}

		for _, char := range strings.Split(lines, "") {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}

			topographicLine = append(topographicLine, value)
		}

		topographicMap = append(topographicMap, topographicLine)
	}

	return topographicMap
}

func exploreTrail(inputTopographicMap [][]int, tops map[[2]int]bool, previousX int, previousY int, x int, y int) {
	if x < 0 || x >= len(inputTopographicMap[0]) || y < 0 || y >= len(inputTopographicMap) {
		return
	}

	if inputTopographicMap[previousY][previousX] + 1 != inputTopographicMap[y][x] {
		return
	}

	if inputTopographicMap[y][x] == 9 {
		tops[[2]int{x, y}] = true
		return
	}

	exploreTrail(inputTopographicMap, tops, x, y, x - 1, y)
	exploreTrail(inputTopographicMap, tops, x, y, x + 1, y)
	exploreTrail(inputTopographicMap, tops, x, y, x, y - 1)
	exploreTrail(inputTopographicMap, tops, x, y, x, y + 1)
}

func part1(inputTopographicMap [][]int) int {
	topographicMap := make([][]int, len(inputTopographicMap))
	for y, line := range inputTopographicMap {
		topographicMap[y] = make([]int, len(line))
	}

	for y, line := range inputTopographicMap {
		for x, value := range line {
			if value == 0 {
				tops := map[[2]int]bool{}
				exploreTrail(inputTopographicMap, tops, x, y, x - 1, y)
				exploreTrail(inputTopographicMap, tops, x, y, x + 1, y)
				exploreTrail(inputTopographicMap, tops, x, y, x, y - 1)
				exploreTrail(inputTopographicMap, tops, x, y, x, y + 1)
				topographicMap[y][x] = len(tops)
			}
		}
	}

	sum := 0

	for _, line := range topographicMap {
		for _, value := range line {
			sum += value
		}
	}
	
	return sum
}


func exploreTrailAll(inputTopographicMap [][]int, topographicMap [][]int, headX int, headY int, previousX int, previousY int, x int, y int) {
	if x < 0 || x >= len(inputTopographicMap[0]) || y < 0 || y >= len(inputTopographicMap) {
		return
	}

	if inputTopographicMap[previousY][previousX] + 1 != inputTopographicMap[y][x] {
		return
	}

	if inputTopographicMap[y][x] == 9 {
		topographicMap[headY][headX]++
		return
	}

	exploreTrailAll(inputTopographicMap, topographicMap, headX, headY, x, y, x - 1, y)
	exploreTrailAll(inputTopographicMap, topographicMap, headX, headY, x, y, x + 1, y)
	exploreTrailAll(inputTopographicMap, topographicMap, headX, headY, x, y, x, y - 1)
	exploreTrailAll(inputTopographicMap, topographicMap, headX, headY, x, y, x, y + 1)
}

func part2(inputTopographicMap [][]int) int {
	topographicMap := make([][]int, len(inputTopographicMap))
	for y, line := range inputTopographicMap {
		topographicMap[y] = make([]int, len(line))
	}

	for y, line := range inputTopographicMap {
		for x, value := range line {
			if value == 0 {
				exploreTrailAll(inputTopographicMap, topographicMap, x, y, x, y, x - 1, y)
				exploreTrailAll(inputTopographicMap, topographicMap, x, y, x, y, x + 1, y)
				exploreTrailAll(inputTopographicMap, topographicMap, x, y, x, y, x, y - 1)
				exploreTrailAll(inputTopographicMap, topographicMap, x, y, x, y, x, y + 1)
			}
		}
	}

	sum := 0

	for _, line := range topographicMap {
		for _, value := range line {
			sum += value
		}
	}
	
	return sum
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	topographicMap := parseInput(input, "\r\n")

	result := part1(topographicMap)
	fmt.Println("checksum is:", result)

	result = part2(topographicMap)
	fmt.Println("checksum is:", result)
}