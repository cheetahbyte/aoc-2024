package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
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

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func part1(left, right []int) {
	distance := 0
	for i := 0; i < len(left); i++ {
		distance += diff(left[i], right[i])
	}
	println("Part 1:", distance)
}

func part2(left, right []int) {
	similarityScore := 0
	for _, element := range left {
		appearances := 0
		for _, rightElem := range right {
			if rightElem == element {
				appearances += 1
			}
		}
		similarityScore += element * appearances
	}
	println("Part 2:", similarityScore)
}

func main() {
	lines := readLineByLine()
	var left []int
	var right []int
	for _, element := range lines {
		slice := strings.Split(element, "   ")
		first, _ := strconv.Atoi(slice[0])
		last, _ := strconv.Atoi(slice[1])
		left = append(left, first)
		right = append(right, last)
	}
	slices.Sort(left)
	slices.Sort(right)
	//
	part1(left, right)
	part2(left, right)
}
