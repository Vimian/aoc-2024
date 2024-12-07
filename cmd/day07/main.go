package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	testValue int
	numbers []int
}

func parseInput(input string, lineEnding string) ([]equation) {
	equations := []equation{}

	for _, line := range strings.Split(input, lineEnding) {
		equation := equation{}

		parts := strings.Split(line, " ")

		var err error
		equation.testValue, err = strconv.Atoi(parts[0][0:len(parts[0]) - 1])
		if err != nil {
			panic(err)
		}

		for _, part := range parts[1:] {
			number, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			equation.numbers = append(equation.numbers, number)
		}

		equations = append(equations, equation)
	}

	return equations
}

func addition(a int, b int) int {
	return a + b
}

func multiplication(a int, b int) int {
	return a * b
}

func concatenate(a int, b int) int {
	result, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		panic(err)
	}

	return result
}

func calculate(number []int, setting int, operators []func(int, int) int) int {
	result := number[0]
	for i := 1; i < len(number); i++ {
		operator := setting % len(operators)
		setting /= len(operators)

		result = operators[operator](result, number[i])
	}

	return result
}

func sumOfValidEquations(equations []equation, operators []func(int, int) int) int {
	result := 0

	for _, equation := range equations {
		for i := 0; i < int(math.Pow(float64(len(operators)), float64(len(equation.numbers) - 1))); i++ {
			if equation.testValue == calculate(equation.numbers, i, operators) {
				result += equation.testValue
				break
			}
		}
	}

	return result
}

func part1(equations []equation) int {
	operators := []func(int, int) int{addition, multiplication}
	
	result := sumOfValidEquations(equations, operators)

	fmt.Printf("total calibration result is %d\n", result)
	return result
}

func part2(equations []equation) int {
	operators := []func(int, int) int{addition, multiplication, concatenate}

	result := sumOfValidEquations(equations, operators)

	fmt.Printf("total calibration result is %d\n", result)
	return result
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	equations := parseInput(input, "\r\n")

	part1(equations)

	part2(equations)
}