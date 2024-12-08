package main

import (
	"fmt"
	"os"
	"strings"
)

type antenna struct {
	x int
	y int
}
type antennas map[string][]antenna
type dimensions struct {
	width int
	height int
}

func parseInput(input string, lineEnding string) (antennas, dimensions) {
	antennas := antennas{}

	lines := strings.Split(input, lineEnding)
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "." {
				continue
			}
			
			if _, found := antennas[char]; !found {
				antennas[char] = []antenna{}
			}
			
			antennas[char] = append(antennas[char], antenna{x, y})
		}
	}

	return antennas, dimensions{width: len(lines[0]), height: len(lines)}
}

func part1(antennas antennas, dimensions dimensions) int {
	antinodes := map[[2]int]bool{}

	for _, antennaList := range antennas {
		for i, antenna := range antennaList {
			for j, otherAntenna := range antennaList {
				if i == j {
					continue
				}

				antinode := [2]int{antenna.x + antenna.x - otherAntenna.x, antenna.y + antenna.y - otherAntenna.y}

				if antinode[0] < 0 || antinode[0] >= dimensions.width || antinode[1] < 0 || antinode[1] >= dimensions.height {
					continue
				}

				antinodes[antinode] = true
			}
		}
	}

	return len(antinodes)
}

func part2(antennas antennas, dimensions dimensions) int {
	antinodes := map[[2]int]bool{}

	for _, antennaList := range antennas {
		for i, antenna := range antennaList {
			antinodes[[2]int{antenna.x, antenna.y}] = true

			for j, otherAntenna := range antennaList {
				if i == j {
					continue
				}

				difference := [2]int{antenna.x - otherAntenna.x, antenna.y - otherAntenna.y}

				for k := 1; true; k++ {
					antinode := [2]int{antenna.x + (difference[0] * k), antenna.y + (difference[1] * k)}

					if antinode[0] < 0 || antinode[0] >= dimensions.width || antinode[1] < 0 || antinode[1] >= dimensions.height {
						break
					}

					antinodes[antinode] = true
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	antennas, dimensions := parseInput(input, "\r\n")

	result := part1(antennas, dimensions)
	fmt.Println("unique antinodes on the map:", result)

	result = part2(antennas, dimensions)
	fmt.Println("unique antinodes on the map:", result)
}