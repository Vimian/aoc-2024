package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type machine struct {
	a [2]int
	b [2]int
	prize [2]int
}

func parseCoordinates(button string, divider [2]string) [2]int {
	coordinates := [2]int{}

	var err error
	coordinates[0], err = strconv.Atoi(strings.Split(strings.Split(button, divider[0])[1], ",")[0])
	if err != nil {
		panic(err)
	}

	coordinates[1], err = strconv.Atoi(strings.Split(strings.Split(button, divider[1])[1], ",")[0])
	if err != nil {
		panic(err)
	}

	return coordinates
}

func parseInput(input string, lineEnding string) ([]machine) {
	machines := []machine{}

	for _, rawMachine := range strings.Split(input, lineEnding + lineEnding) {
		machine := machine{}

		line := strings.Split(rawMachine, lineEnding)
		machine.a = parseCoordinates(line[0], [2]string{"Y", "X"})
		machine.b = parseCoordinates(line[1], [2]string{"Y", "X"})
		machine.prize = parseCoordinates(line[2], [2]string{"Y=", "X="})
		
		machines = append(machines, machine)
	}

	return machines
}

func part1(machines []machine) int {
	sum := 0

	for _, machine := range machines {
		bMultiplier := ((machine.prize[1] * machine.a[0]) - (machine.prize[0] * machine.a[1])) / ((machine.b[1] * machine.a[0]) - (machine.b[0] * machine.a[1]))
		
		aMultiplier := (machine.prize[0] - (machine.b[0] * bMultiplier)) / machine.a[0]

		if bMultiplier * machine.b[0] + aMultiplier * machine.a[0] == machine.prize[0] && bMultiplier * machine.b[1] + aMultiplier * machine.a[1] == machine.prize[1] {
			sum += bMultiplier + (aMultiplier * 3)
		}
	}

	return sum
}

func part2(machines []machine) int {
	for i := 0; i < len(machines); i++ {
		machines[i].prize[0] += 10000000000000
		machines[i].prize[1] += 10000000000000
	}

	return part1(machines)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	parsedInput := parseInput(input, "\r\n")

	result := part1(parsedInput)
	fmt.Println("fewest tokens:", result)

	result = part2(parsedInput)
	fmt.Println("fewest tokens:", result)
}