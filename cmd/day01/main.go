package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func loadInput(input string) ([]int, []int) {
	leftList := []int{}
	rightList := []int{}

	lines := strings.Split(input, "\r\n")

	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		left, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		
		leftList = append(leftList, left)
		
		right, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		rightList = append(rightList, right)
	}

	return leftList, rightList
}

func part1(leftList []int, rightList []int) {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0

	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	fmt.Printf("distance is %d\n", totalDistance)
}

func part2(leftList []int, rightList []int) {
	rightMap := make(map[int]int)

	for _, right := range rightList {
		rightMap[right]++
	}

	totalSimilarity := 0

	for _, left := range leftList {
		similarity := left * rightMap[left]
		totalSimilarity += similarity
	}

	fmt.Printf("similarity is %d\n", totalSimilarity)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)

	leftList, rightList := loadInput(input)

	part1(leftList, rightList)
	part2(leftList, rightList)	
}