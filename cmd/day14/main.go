package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	position [2]int
	velocity [2]int
}

func parseInput(input string, lineEnding string) ([]robot) {
	parsedInput := []robot{}

	for _, line := range strings.Split(input, lineEnding) {
		parsedLine := robot{}

		temp := strings.Split(line, " v=")
		
		position := strings.Split(strings.Split(temp[0], "p=")[1], ",")

		velocity := strings.Split(temp[1], ",")
		
		for i, number := range position {
			value, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			parsedLine.position[i] = value
		}

		for i, number := range velocity {
			value, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			parsedLine.velocity[i] = value
		}

		parsedInput = append(parsedInput, parsedLine)
	}

	return parsedInput
}

func simulate(robots *[]robot, dimentions [2]int, duration int) {
	for j, robot := range *robots {
		(*robots)[j].position[0] = (robot.position[0] + (robot.velocity[0] * duration)) % dimentions[0]
		if (*robots)[j].position[0] < 0 {
			(*robots)[j].position[0] += dimentions[0]
		}

		(*robots)[j].position[1] = (robot.position[1] + (robot.velocity[1] * duration)) % dimentions[1]
		if (*robots)[j].position[1] < 0 {
			(*robots)[j].position[1] += dimentions[1]
		}
	}
}

func part1(inputRobots []robot, dimentions [2]int, duration int) int {
	robots := make([]robot, len(inputRobots))
	copy(robots, inputRobots)

	simulate(&robots, dimentions, duration)

	quadrants := [2][2]int{{0, 0}, {0, 0}}
	for _, robot := range robots {
		if robot.position[0] < dimentions[0] / 2 {
			if robot.position[1] < dimentions[1] / 2 {
				quadrants[0][0]++
			} else if robot.position[1] > dimentions[1] / 2 {
				quadrants[0][1]++
			}
		} else if robot.position[0] > dimentions[0] / 2 {
			if robot.position[1] < dimentions[1] / 2 {
				quadrants[1][0]++
			} else if robot.position[1] > dimentions[1] / 2 {
				quadrants[1][1]++
			}
		}
	}

	return quadrants[0][0] * quadrants[0][1] * quadrants[1][0] * quadrants[1][1]
}

func part2(robots []robot, dimentions [2]int) int {
	i := 0
	out:
	for true {
		simulate(&robots, dimentions, 1)
		i++

		temp := map[[2]int]bool{}
		for _, robot := range robots {
			temp[robot.position] = true
		}

		inner:
		for position := range temp {
			for j := 1; j <= 10; j++ {
				if _, found := temp[[2]int{position[0], position[1] + j}]; !found {
					continue inner
				}
			}
			break out
		}
	}

	fmt.Println("After", i, "seconds:")

	grid := make([][]string, dimentions[1])
	for i := range grid {
		grid[i] = make([]string, dimentions[0])
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = " "
		}
	}

	for _, robot := range robots {
		grid[robot.position[1]][robot.position[0]] = "X"
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	return i
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	parsedInput := parseInput(input, "\r\n")

	result := part1(parsedInput, [2]int{101, 103}, 100)
	fmt.Println("safety factor is:", result)

	result = part2(parsedInput, [2]int{101, 103})
	fmt.Println("seconds to first easter egg:", result)
}