package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseInput(input string, lineEnding string) ([]string, []string) {
	parts := strings.Split(input, lineEnding + lineEnding) 

	towels := strings.Split(parts[0], ", ")

	designs := strings.Split(parts[1], lineEnding)

	return towels, designs
}

func part1(towels []string, designs []string) (int, []string) {
	regexExpression := "^(?:" + strings.Join(towels, "|") + ")+$"

	regex := regexp.MustCompile(regexExpression)

	validDesigns := []string{}
	for _, design := range designs {
		if len(regex.FindAllString(design, -1)) == 0 {
			continue
		}

		validDesigns = append(validDesigns, design)
	}
	return len(validDesigns), validDesigns
}

func part2(towels []string, designs []string) int {
	sum := 0

	for _, design := range designs {
		matches := map[int]map[int]bool{}
		for _, towel := range towels {
			for i := 0; i < len(design) - len(towel) + 1; i++ {
				if design[i:i+len(towel)] != towel {
					continue
				}

				if _, found := matches[i]; !found {
					matches[i] = map[int]bool{}
				}
				matches[i][i+len(towel)] = true
			}
		}

		validEnds := map[int]int{len(design): 1}
		for i := len(design) - 1; i >= 0; i-- {
			if _, found := validEnds[i]; !found {
				validEnds[i] = 0
			}

			if _, found := matches[i]; !found {
				continue
			}

			for end := range matches[i] {
				if _, found := validEnds[end]; !found {
					continue
				}

				validEnds[i] += validEnds[end]
			}
		}

		sum += validEnds[0]
	}

	return sum
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	towels, designs := parseInput(input, "\r\n")

	result, _ := part1(towels, designs)
	fmt.Println("designs possible:", result)

	result = part2(towels, designs)
	fmt.Println("total number of different ways you could make each design:", result)
}