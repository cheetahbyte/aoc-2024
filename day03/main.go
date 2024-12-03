package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readLineByLine() []string {
	file, _ := os.Open("./input.txt")
	var lines []string
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func part1(lines []string) {
	regex := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	var sum int
	for _, line := range lines {
		statements := regex.FindAllStringSubmatch(line, -1)
		for _, statement := range statements {
			first, _ := strconv.Atoi(statement[1])
			second, _ := strconv.Atoi(statement[2])
			sum += first * second
		}
	}
	fmt.Println("Part 1: ", sum)
}

func part2(lines []string) {
	// Combined regex pattern to match mul(x,y), do(), and don't()
	// The pattern captures the instruction and the numbers if present
	regex := regexp.MustCompile(`do\(\)|don't\(\)|mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	// Track the state of `mul` operations
	mulEnabled := true
	var sum int

	for _, line := range lines {
		// Find all instructions in order
		instructions := regex.FindAllStringSubmatch(line, -1)

		// Process each instruction sequentially
		for _, instr := range instructions {
			command := instr[0]
			switch {
			case command == "do()":
				mulEnabled = true
			case command == "don't()":
				mulEnabled = false
			default:
				// Must be a mul(x,y)
				if mulEnabled {
					first, _ := strconv.Atoi(instr[1])
					second, _ := strconv.Atoi(instr[2])
					sum += first * second
				}
			}
		}
	}

	fmt.Println("Part 2: ", sum)

}

func main() {
	lines := readLineByLine()
	part1(lines)
	part2(lines)

}
