package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var partOneRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var partTwoRegex = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func part1(lines []string) {
	// Compile the regex once

	sum := 0

	for _, line := range lines {
		// Find all matches in the line
		matches := partOneRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// Convert the matched numbers to integers
			first, _ := strconv.Atoi(match[1])
			second, _ := strconv.Atoi(match[2])
			sum += first * second
		}
	}

	fmt.Println("Part 1:", sum)
}

func part2(lines []string) {
	// Compile the regex once
	mulEnabled := true
	sum := 0

	for _, line := range lines {
		// Find all matches in the line
		matches := partTwoRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch match[0] {
			case "do()":
				mulEnabled = true
			case "don't()":
				mulEnabled = false
			default: // Handle "mul(x, y)" cases
				if mulEnabled {
					first, _ := strconv.Atoi(match[1])
					second, _ := strconv.Atoi(match[2])
					sum += first * second
				}
			}
		}
	}

	fmt.Println("Part 2:", sum)
}

func main() {
	lines := readLines("input.txt")
	part1(lines)
	part2(lines)
}
