package main

import (
	"bufio"
	"fmt"
	"os"
)

func readIntoGrid() [][]rune {
	file, _ := os.Open("./input.txt")
	var grid [][]rune
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	return grid
}

func searchDirection(grid [][]rune, rowDelta, colDelta int) int {
	word := "XMAS"
	wordLength := len(word)
	wordRunes := []rune(word)
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			match := true
			for k := 0; k < wordLength; k++ {
				r := row + k*rowDelta
				c := col + k*colDelta
				if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != wordRunes[k] {
					match = false
					break
				}
			}
			if match {
				count++
			}
		}
	}
	return count
}

func part1(lines [][]rune) {
	directions := [][2]int{
		{0, 1},   // Horizontal right
		{0, -1},  // Horizontal left
		{1, 0},   // Vertical down
		{-1, 0},  // Vertical up
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{1, -1},  // Diagonal down-left
		{-1, 1},  // Diagonal up-right
	}

	total := 0
	for _, dir := range directions {
		total += searchDirection(lines, dir[0], dir[1])
	}
	fmt.Println("Part 1:", total)
}

func main() {
	grid := readIntoGrid()
	part1(grid)
}
