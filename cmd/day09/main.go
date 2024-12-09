package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string, lineEnding string) ([]*int) {
	disk := []*int{}

	for i, char := range strings.Split(input, "") {
		number, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}

		if i % 2 == 0 {
			part := make([]*int, number)
			value := int(i/2)
			for i := range part {
				part[i] = &value
			}
			disk = append(disk, part...)
		} else {
			disk = append(disk, make([]*int, number)...)
		}
	}

	return disk
}

func calulateChecksum(disk []*int) int {
	checksum := 0

	for i, value := range disk {
		if value != nil {
			checksum += i * *value
		}
	}

	return checksum
}

func part1(inputDisk []*int) int {
	disk := make([]*int, len(inputDisk))
	copy(disk, inputDisk)

	rightPosition := len(disk) - 1

	for i := 0; i < rightPosition; {
		if disk[rightPosition] == nil {
			rightPosition--
			continue
		}
		
		if disk[i] == nil {
			disk[i] = disk[rightPosition]
			disk[rightPosition] = nil
			rightPosition--
		}

		i++
	}

	return calulateChecksum(disk)
}

func part2(inputDisk []*int) int {
	disk := make([]*int, len(inputDisk))
	copy(disk, inputDisk)

	index := *disk[len(disk) - 1]

	for i := len(disk) - 1; i >= 0; {
		if disk[i] == nil {
			i--
			continue
		}
		
		length := 1
		for true {
			if i-length < 0 {
				break
			}

			if disk[i-length] == disk[i] {
				length++
				continue
			}

			break
		}
		if index < *disk[i] {
			i -= length
			continue
		}
		
		index--

		freeSpace := 0
		freeSpaceEndIndex := 0
		for j := 0; j < i; j++ {
			if disk[j] != nil {
				freeSpace = 0
				continue
			}

			freeSpace++

			if freeSpace == length {
				freeSpaceEndIndex = j
				break
			}
		}

		if freeSpace < length {
			i -= length
			continue
		}

		for j := 0; j < length; j++ {
			disk[freeSpaceEndIndex-length+1+j] = disk[i-j]
			disk[i-j] = nil
		}
	}
	
	return calulateChecksum(disk)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	disk := parseInput(input, "\r\n")

	result := part1(disk)
	fmt.Println("checksum is:", result)

	result = part2(disk)
	fmt.Println("checksum is:", result)
}