package main

import (
	"fmt"
	"os"
	"strings"
)

func stringTo2dArray(input string) [][]string {
	var result [][]string

	for _, line := range strings.Split(input, "\r\n") {
		var row []string
		for _, char := range strings.Split(line, "") {
			row = append(row, char)
		}
		result = append(result, row)
	}

	return result
}

func part1(input [][]string, word []string) {
	directions := [][]int{{-1, 0},{-1, 1},{0, 1},{1, 1},{1, 0},{1, -1},{0, -1},{-1, -1}}

	matches := 0

	for y, row := range input {
		for x, char := range row {
			if char == word[0] {
				for _, direction := range directions {
					for i := 1; i < len(word); i++ {
						newY := y + (direction[0] * i)
						newX := x + (direction[1] * i)

						if newY < 0 || newY >= len(input) || newX < 0 || newX >= len(row) {
							break
						}

						if input[newY][newX] != word[i] {
							break
						}

						if i == len(word) - 1 {
							matches++
						}
					}
				}
			}
		}
	}

	fmt.Printf("found %d matches\n", matches)
}

func part2(input [][]string, word []string) {
	directions := [][]int{{-1, 0},{0, 1},{1, 0},{0, -1}}

	matches := 0

	for y, row := range input {
		for x, char := range row {
			if char == word[0] {
				for _, direction := range directions {
					for i := 0; i < len(word); i++ {
						newY := y + (direction[0] * i) + (-direction[1] * i)
						newX := x + (direction[1] * i) + (-direction[0] * i)
						newY2 := y + (direction[0] * i) + (-direction[1] * (len(word) - 1 - i))
						newX2 := x + (direction[1] * i) + (-direction[0] * (len(word) - 1 - i))

						if newY < 0 || newY >= len(input) || newY2 < 0 || newY2 >= len(input) || newX < 0 || newX >= len(row) || newX2 < 0 || newX2 >= len(row) {
							break
						}

						if input[newY][newX] != word[i] || input[newY2][newX2] != word[i] {
							break
						}

						if i == len(word) - 1 {
							matches++
						}
					}
				}
			}
		}
	}

	fmt.Printf("found %d matches\n", matches)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	input2d := stringTo2dArray(input)

	word := strings.Split("XMAS", "")
	part1(input2d, word)
	
	word = strings.Split("MAS", "")
	part2(input2d, word)
}