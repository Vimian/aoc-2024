package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var input string = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	parsedRegister, parsedProgram := parseInput(input, "\n")

	var expected string = `4,6,3,5,6,3,5,2,1,0`

	var result, _ = part1(parsedRegister, parsedProgram)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart1_1(t *testing.T) {
	var input string = `Register A: 0
Register B: 0
Register C: 9

Program: 2,6`
	parsedRegister, parsedProgram := parseInput(input, "\n")
	var expected int = 1

	var _, register = part1(parsedRegister, parsedProgram)
	var result int = register["B"]
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart1_2(t *testing.T) {
	var input string = `Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4`
	parsedRegister, parsedProgram := parseInput(input, "\n")
	var expected string = `0,1,2`

	var result, _ = part1(parsedRegister, parsedProgram)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart1_3(t *testing.T) {
	var input string = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	parsedRegister, parsedProgram := parseInput(input, "\n")
	var expected string = `4,2,5,6,7,7,7,7,3,1,0`

	var result, register = part1(parsedRegister, parsedProgram)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	var expectedA int = 0
	if register["A"] != expectedA {
		t.Errorf("Expected %d, got %d", expectedA, register["A"])
	}
}

func TestPart1_4(t *testing.T) {
	var input string = `Register A: 0
Register B: 29
Register C: 0

Program: 1,7`
	parsedRegister, parsedProgram := parseInput(input, "\n")
	var expected int = 26

	var _, register = part1(parsedRegister, parsedProgram)
	var result int = register["B"]
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart1_5(t *testing.T) {
	var input string = `Register A: 0
Register B: 2024
Register C: 43690

Program: 4,0`
	parsedRegister, parsedProgram := parseInput(input, "\n")
	var expected int = 44354

	var _, register = part1(parsedRegister, parsedProgram)
	var result int = register["B"]
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
var input string = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`
	parsedRegister, parsedProgram := parseInput(input, "\n")
	var expectedInitialA int = 117440
	var expected string = ""
	for _, instruction := range parsedProgram {
		expected += strconv.Itoa(instruction)
	}
	expected = strings.Join(strings.Split(expected, ""), ",")

	var result, initialA = part2(parsedRegister, parsedProgram)
	if initialA != expectedInitialA {
		t.Errorf("Expected %d, got %d", expectedInitialA, initialA)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2_1(t *testing.T) {
var input string = `Register A: 27334280
Register B: 0
Register C: 0

Program: 2,4,1,2,7,5,0,3,1,7,4,1,5,5,3,0`
	parsedRegister, parsedProgram := parseInput(input, "\n")
	var expectedInitialA int = 190615597431823
	var expected string = ""
	for _, instruction := range parsedProgram {
		expected += strconv.Itoa(instruction)
	}
	expected = strings.Join(strings.Split(expected, ""), ",")

	var result, initialA = part2(parsedRegister, parsedProgram)
	if initialA != expectedInitialA {
		t.Errorf("Expected %d, got %d", expectedInitialA, initialA)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}