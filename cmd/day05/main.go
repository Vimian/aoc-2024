package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitOrderAndUpdates(input string) (map[int]map[int]bool, [][]int) {
	orders := map[int]map[int]bool{}
	updates := [][]int{}

	
	lookForOrders := true
	for _, line := range strings.Split(input, "\r\n") {
		if line == "" {
			lookForOrders = false
			continue
		}

		if lookForOrders {
			numbers := strings.Split(line, "|")
			firstNumber, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			secondNumber, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}

			if _, found := orders[firstNumber]; !found {
				orders[firstNumber] = map[int]bool{}
			}
			orders[firstNumber][secondNumber] = true

		} else {
			updates = append(updates, []int{})
			for _, rawNumber := range strings.Split(line, ",") {
				number, err := strconv.Atoi(rawNumber)
				if err != nil {
					panic(err)
				}

				updates[len(updates)-1] = append(updates[len(updates)-1], number)
			}
		}
	}

	return orders, updates
}

func part1(orders map[int]map[int]bool, updates [][]int) [][]int {
	incorrect := [][]int{}
	result := 0

	out:
	for _, update := range updates {
		for i := 0; i < len(update); i++ {
			if _, found := orders[update[i]]; !found {
				continue
			}

			for j := 0; j < i; j++ {
				if _, found := orders[update[i]][update[j]]; found {
					incorrect = append(incorrect, update)
					continue out
				}
			}
		}

		result += update[len(update) / 2]
	}

	fmt.Println("the result is", result)
	return incorrect
}

func part2(orders map[int]map[int]bool, updates [][]int) {
	out:
	for i := 0; i < len(updates); {
		update := updates[i]
		for j := 0; j < len(update); j++ {
			if _, found := orders[update[j]]; !found {
				continue
			}

			for k := 0; k < j; k++ {
				if _, found := orders[update[j]][update[k]]; found {
					updates[i][j], updates[i][j - 1] = updates[i][j - 1], updates[i][j]
					continue out
				}
			}
		}
		i++
	}

	result := 0
	for _, update := range updates {
		result += update[len(update) / 2]
	}
	fmt.Println("the result is", result)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	orders, updates := splitOrderAndUpdates(input)

	incorrect := part1(orders, updates)

	part2(orders, incorrect)
}