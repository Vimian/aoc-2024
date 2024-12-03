package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) {
	var validMul = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	muls := validMul.FindAllString(input, -1)

	result := 0

	for _, mul := range muls {
		numbers := regexp.MustCompile(`\d{1,3}`).FindAllString(mul, -1)

		a, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		result += a * b
	}

	fmt.Println("the result is", result)
}

func part2(input string) {
	//input = strings.ReplaceAll(input, "\r\n", "")
	dontWants := regexp.MustCompile(`(?s)don't\(\).{0,}?do\(\)`).FindAllString(input, -1)

	for _, dontWant := range dontWants {
		input = strings.ReplaceAll(input, dontWant, "")
	}

	part1(input)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)

	part1(input)
	part2(input)
}