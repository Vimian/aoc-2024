package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type register map[string]int

type program []int

func parseInput(input string, lineEnding string) (register, program) {
	parts := strings.Split(input, lineEnding + lineEnding)

	register := register{}

	for _, line := range strings.Split(parts[0], lineEnding) {
		temp := strings.Split(line, ": ")
		key := strings.Split(temp[0], " ")[1]
		value, err := strconv.Atoi(temp[1])
		if err != nil {
			panic(err)
		}

		register[key] = value
	}
	
	program := program{}
	for _, line := range strings.Split(strings.Split(parts[1], " ")[1], ",") {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		program = append(program, value)		
	}

	return register, program
}

func getComboOperand(register *register, operand int) int {
	value := 0
	switch operand {
	case 0, 1, 2, 3:
		value = operand
	case 4:
		value = (*register)["A"]
	case 5:
		value = (*register)["B"]
	case 6:
		value = (*register)["C"]
	default:
		fmt.Println("unknown operand", operand)
	}

	return value
}

func adv(register *register, registerName string, numerator int, operand int) {
	value := getComboOperand(register, operand)
	if value == 0 {
		return
	}

	(*register)[registerName] = numerator / int(math.Pow(2, float64(value)))
}

func bxor(register *register, registerName string, value int) {
	(*register)[registerName] = (*register)[registerName] ^ value
}

func part1(register register, program program) (string, register) {
	output := ""

	for i := 0; i < len(program); {
		switch program[i] {
		case 0:
			adv(&register, "A", register["A"], program[i+1])
		case 1:
			bxor(&register, "B", program[i+1])
		case 2:
			register["B"] = getComboOperand(&register, program[i+1]) % 8
		case 3:
			if register["A"] != 0 {
				i = program[i+1]
				continue
			}
		case 4:
			bxor(&register, "B", register["C"])
		case 5:
			value := getComboOperand(&register, program[i+1]) % 8
			output += strconv.Itoa(value)
		case 6:
			adv(&register, "B", register["A"], program[i+1])
		case 7:
			adv(&register, "C", register["A"], program[i+1])
		}
		i += 2
	}
	return strings.Join(strings.Split(output, ""), ","), register
}

func part2(inputRegister register, program program) (string, int) {
	initialA := 0
	correctBits := -1
	correctBit := 0
	counter := 1

	out:
	for j := 1; true; j++ {
		inputRegister["A"] = initialA
		register := register{}
		for key, value := range inputRegister {
			register[key] = value
		}

		output := 0
		temp := false

		for i := 0; i < len(program); {
			switch program[i] {
			case 0:
				adv(&register, "A", register["A"], program[i+1])
			case 1:
				bxor(&register, "B", program[i+1])
			case 2:
				register["B"] = getComboOperand(&register, program[i+1]) % 8
			case 3:
				if register["A"] != 0 {
					i = program[i+1]
					continue
				}
			case 4:
				bxor(&register, "B", register["C"])
			case 5:
				value := getComboOperand(&register, program[i+1]) % 8
				if value != program[output] {
					break
				}
				output++
				
				if output == len(program) {
					break out
				}

				if output == counter {
					correctBits = initialA
					n := correctBits
					correctBit = 0
					for n > 0 {
						correctBit++
						n >>= 1
					}
					fmt.Println(counter, "out of", len(program), "initialA is", initialA, "correctBit is", correctBit)
					temp = true
					counter++
				}
			case 6:
				adv(&register, "B", register["A"], program[i+1])
			case 7:
				adv(&register, "C", register["A"], program[i+1])
			}
			i += 2
		}

		if temp {
			j = 1
		}

		initialA = j
		for n := 0 ; n < correctBit; n++ {
			initialA <<= 1
		}

		if correctBits == 0 {
			initialA <<= 1
		}

		initialA += correctBits
	}

	fmt.Println("initial A is", initialA)

	var expected string = ""
	for _, instruction := range program {
		expected += strconv.Itoa(instruction)
	}
	expected = strings.Join(strings.Split(expected, ""), ",")
		
	return expected, initialA
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	register, program := parseInput(input, "\r\n")

	result, _ := part1(register, program)
	fmt.Println("output is:", result)

	_, initialA := part2(register, program)
	fmt.Println("lowest positive initial value for register A is:", initialA)
}