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
	fmt.Println("towels:", towels)

	sum := 0

	for _, design := range designs {
		fmt.Println("design:", design)
		matches := map[int]map[int]bool{}
		for _, towel := range towels {
			indexs := regexp.MustCompile(towel).FindAllStringIndex(design, -1)
			for _, index := range indexs {
				if _, found := matches[index[0]]; !found {
					matches[index[0]] = map[int]bool{}
				}
				matches[index[0]][index[1]] = true
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
				/*if _, found := validEnds[end]; !found {
					continue
				}*/

				validEnds[i] += validEnds[end]
			}
		}

		fmt.Println("matches:", matches)
		fmt.Println("validEnds:", validEnds)
		sum += validEnds[0]
		fmt.Println("sum:", sum)
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

	result = part2(towels, designs) // given answer 611033244835885
	fmt.Println("total number of different ways you could make each design:", result) // right answer 950763269786650
}