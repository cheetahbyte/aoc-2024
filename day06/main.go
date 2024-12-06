package main

import (
	"bufio"
	"fmt"
	"os"
)

func readIntoGrid() [][]rune {
	var grid [][]rune
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	return grid
}

var directions = []struct {
	dx, dy int
}{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func part1(grid [][]rune) {
	// Find the guard's starting position and direction
	var x, y, dir int
	for i := range grid {
		for j, cell := range grid[i] {
			if cell == '^' {
				x, y, dir = i, j, 0 // Facing up
				break
			}
		}
	}

	// Set to track visited positions
	visited := make(map[[2]int]bool)
	visited[[2]int{x, y}] = true

	rows, cols := len(grid), len(grid[0])

	// Simulate guard's movement
	for {
		// Calculate the next position
		nx, ny := x+directions[dir].dx, y+directions[dir].dy

		// Check if next position is out of bounds
		if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
			break
		}

		// Check if there's an obstacle
		if grid[nx][ny] == '#' {
			// Turn right (90 degrees)
			dir = (dir + 1) % 4
		} else {
			// Move to the next position
			x, y = nx, ny
			visited[[2]int{x, y}] = true
		}
	}

	// Output the number of distinct positions visited
	fmt.Println("Part 1:", len(visited))
}
func simulate(grid [][]rune, startX, startY, startDir int) bool {
	// Simulate guard's movement and detect loops
	x, y, dir := startX, startY, startDir
	visited := make(map[[3]int]bool) // Track visited positions with direction

	rows, cols := len(grid), len(grid[0])

	for {
		// Check for loops
		state := [3]int{x, y, dir}
		if visited[state] {
			return true
		}
		visited[state] = true

		// Calculate the next position
		nx, ny := x+directions[dir].dx, y+directions[dir].dy

		// Check if next position is out of bounds
		if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
			return false
		}

		// Check if there's an obstacle
		if grid[nx][ny] == '#' {
			// Turn right (90 degrees)
			dir = (dir + 1) % 4
		} else {
			// Move to the next position
			x, y = nx, ny
		}
	}
}

func part2(grid [][]rune) {
	// Find the guard's starting position and direction
	var startX, startY, startDir int
	for i := range grid {
		for j, cell := range grid[i] {
			if cell == '^' {
				startX, startY, startDir = i, j, 0 // Facing up
				break
			}
		}
	}

	rows, cols := len(grid), len(grid[0])
	loopCount := 0

	// Test placing an obstruction at each empty space
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '.' && !(i == startX && j == startY) {
				// Temporarily place an obstruction
				grid[i][j] = '#'
				if simulate(grid, startX, startY, startDir) {
					loopCount++
				}
				// Remove the obstruction
				grid[i][j] = '.'
			}
		}
	}

	fmt.Println("Part 2:", loopCount)
}

func main() {
	grid := readIntoGrid()
	part1(grid)
	part2(grid)
}
