package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// same as in day01
func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// modified version of day01
func readLineByLine() [][]int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var rows [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var row []int
		for _, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalf("Error converting to int: %v", err)
			}
			row = append(row, val)
		}
		rows = append(rows, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return rows
}

func isSafe(row []int) bool {
	increasing := true
	decreasing := true

	for i := 1; i < len(row); i++ {
		difference := diff(row[i-1], row[i])
		if difference < 1 || difference > 3 {
			return false
		}
		if row[i] > row[i-1] {
			decreasing = false
		} else if row[i] < row[i-1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

func part1(rows [][]int) int {
	safeCount := 0
	for _, row := range rows {
		if isSafe(row) {
			safeCount++
		}
	}
	return safeCount
}

func isSafeWithRemoval(row []int) bool {
	if isSafe(row) {
		return true
	}
	for i := 0; i < len(row); i++ {
		newRow := append([]int{}, row[:i]...)
		newRow = append(newRow, row[i+1:]...)
		if isSafe(newRow) {
			return true
		}
	}
	return false
}

func part2(rows [][]int) int {
	safeCount := 0
	for _, row := range rows {
		if isSafeWithRemoval(row) {
			safeCount++
		}
	}
	return safeCount
}

func main() {
	rows := readLineByLine()
	fmt.Printf("Part One: %d\n", part1(rows))
	fmt.Printf("Part Two: %d\n", part2(rows))
}
