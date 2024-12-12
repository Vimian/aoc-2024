package main

import (
	"fmt"
	"os"
	"strings"
)

type region struct {
	plant string
	plots map[[2]int]bool
	boundary [2][2]int
	area [][]int
	edges int
	sides int
}

func explore(region *region, parsedInput *[][]string, next [2]int, plant string) {
	if region.plots[next] {
		return
	}

	if next[0] < 0 || next[0] >= len(*parsedInput) || next[1] < 0 || next[1] >= len((*parsedInput)[0]) {
		return
	}

	if (*parsedInput)[next[0]][next[1]] != plant {
		return
	}

	(*parsedInput)[next[0]][next[1]] = ""
	(*region).plots[next] = true

	if next[0] < (*region).boundary[0][0] {
		(*region).boundary[0][0] = next[0]
	} else if next[0] > (*region).boundary[1][0] {
		(*region).boundary[1][0] = next[0]
	}

	if next[1] < (*region).boundary[0][1] {
		(*region).boundary[0][1] = next[1]
	} else if next[1] > (*region).boundary[1][1] {
		(*region).boundary[1][1] = next[1]
	}

	explore(region, parsedInput, [2]int{next[0]-1, next[1]}, plant)
	explore(region, parsedInput, [2]int{next[0]+1, next[1]}, plant)
	explore(region, parsedInput, [2]int{next[0], next[1]-1}, plant)
	explore(region, parsedInput, [2]int{next[0], next[1]+1}, plant)
}

func parseInput(input string, lineEnding string) ([]region) {
	parsedInput := [][]string{}

	for _, lines := range strings.Split(input, lineEnding) {
		parsedLine := []string{}

		for _, char := range strings.Split(lines, "") {
			parsedLine = append(parsedLine, char)
		}

		parsedInput = append(parsedInput, parsedLine)
	}

	garden := []region{}

	for y, row := range parsedInput {
		for x, char := range row {
			if char == "" {
				continue
			}

			newRegion := region{
				plant: char,
				plots: map[[2]int]bool{},
				boundary: [2][2]int{{y, x}, {y, x}},
				area: [][]int{},
				edges: 0,
				sides: 0,
			}

			newRegion.plots[[2]int{y, x}] = true

			explore(&newRegion, &parsedInput, [2]int{y-1, x}, char)
			explore(&newRegion, &parsedInput, [2]int{y+1, x}, char)
			explore(&newRegion, &parsedInput, [2]int{y, x-1}, char)
			explore(&newRegion, &parsedInput, [2]int{y, x+1}, char)

			parsedInput[y][x] = ""

			garden = append(garden, newRegion)
		}
	}

	for i, region := range garden {
		for y := region.boundary[0][0]; y <= region.boundary[1][0]; y++ {
			areaRow := make([]int, region.boundary[1][1]-region.boundary[0][1]+1)
			garden[i].area = append(garden[i].area, areaRow)
		}
		
		for plot := range region.plots {
			garden[i].area[plot[0]-region.boundary[0][0]][plot[1]-region.boundary[0][1]] = 1
		}
	}

	return garden
}

func part1(garden *[]region) int {
	for i, region := range *garden {
		for y, row := range region.area {
			for x, cell := range row {
				if cell == 0 {
					continue
				}

				if y == 0 || region.area[y-1][x] == 0 {
					(*garden)[i].edges++
				}
				if y == len(region.area)-1 || region.area[y+1][x] == 0 {
					(*garden)[i].edges++
				}
				if x == 0 || region.area[y][x-1] == 0 {
					(*garden)[i].edges++
				}
				if x == len(row)-1 || region.area[y][x+1] == 0 {
					(*garden)[i].edges++
				}
			}
		}
	}

	total := 0

	for _, region := range *garden {
		total += len(region.plots) * region.edges
	}

	return total
}

func part2(garden *[]region) int {
	for i, region := range *garden {
		for y, row := range region.area {
			followingSideTop := false
			followingSideBottom := false
			for x, cell := range row {
				if cell == 0 {
					if followingSideTop {
						(*garden)[i].sides++
					}
					if followingSideBottom {
						(*garden)[i].sides++
					}
					followingSideTop = false
					followingSideBottom = false
					continue
				}

				if y == 0 || region.area[y-1][x] == 0 {
					followingSideTop = true
				} else {
					if followingSideTop {
						(*garden)[i].sides++
					}
					followingSideTop = false
				}

				if y == len(region.area)-1 || region.area[y+1][x] == 0 {
					followingSideBottom = true
				} else {
					if followingSideBottom {
						(*garden)[i].sides++
					}
					followingSideBottom = false
				}
			}
			if followingSideTop {
				(*garden)[i].sides++
			}
			if followingSideBottom {
				(*garden)[i].sides++
			}
		}

		for x := 0; x < len(region.area[0]); x++ {
			followingSideLeft := false
			followingSideRight := false
			for y := 0; y < len(region.area); y++ {
				if region.area[y][x] == 0 {
					if followingSideLeft {
						(*garden)[i].sides++
					}
					if followingSideRight {
						(*garden)[i].sides++
					}
					followingSideLeft = false
					followingSideRight = false
					continue
				}

				if x == 0 || region.area[y][x-1] == 0 {
					followingSideLeft = true
				} else {
					if followingSideLeft {
						(*garden)[i].sides++
					}
					followingSideLeft = false
				}
				
				if x == len(region.area[y])-1 || region.area[y][x+1] == 0 {
					followingSideRight = true
				} else {
					if followingSideRight {
						(*garden)[i].sides++
					}
					followingSideRight = false
				}
			}
			if followingSideLeft {
				(*garden)[i].sides++
			}
			if followingSideRight {
				(*garden)[i].sides++
			}
		}
	}

	total := 0

	for _, region := range *garden {
		total += len(region.plots) * region.sides
	}

	return total
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	parsedInput := parseInput(input, "\r\n")

	result := part1(&parsedInput)
	fmt.Println("checksum is:", result)

	result = part2(&parsedInput)
	fmt.Println("checksum is:", result)
}