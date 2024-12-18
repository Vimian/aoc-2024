package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string, lineEnding string) ([][]int) {
	parsedInput := [][]int{}

	for _, line := range strings.Split(input, lineEnding) {
		parsedLine := []int{}

		for _, char := range strings.Split(line, "") {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}

			parsedLine = append(parsedLine, value)
		}

		parsedInput = append(parsedInput, parsedLine)
	}

	return parsedInput
}

func part1() int {
	return 0
}

func part2() int {
	return 0
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	parsedInput := parseInput(input, "\r\n")

	result := part1(parsedInput)
	fmt.Println("checksum:", result)

	/*result = part2(parsedInput)
	fmt.Println("checksum:", result)*/
}