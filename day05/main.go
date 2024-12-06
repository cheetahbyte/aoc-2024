package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() (map[int][]int, [][]int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	rules := make(map[int][]int)
	var updates [][]int
	isUpdateSection := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isUpdateSection = true
			continue
		}

		if !isUpdateSection {
			// Parse rules
			parts := strings.Split(line, "|")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			rules[x] = append(rules[x], y)
		} else {
			// Parse updates
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))
			for i, p := range parts {
				update[i], _ = strconv.Atoi(p)
			}
			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}
func isValidUpdate(update []int, rules map[int][]int) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for x, dependencies := range rules {
		if posX, exists := position[x]; exists {
			for _, y := range dependencies {
				if posY, existsY := position[y]; existsY && posY < posX {
					return false
				}
			}
		}
	}
	return true
}

func reorderUpdate(update []int, rules map[int][]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, page := range update {
		inDegree[page] = 0
	}

	for x, dependencies := range rules {
		if _, exists := inDegree[x]; exists {
			for _, y := range dependencies {
				if _, exists := inDegree[y]; exists {
					graph[x] = append(graph[x], y)
					inDegree[y]++
				}
			}
		}
	}

	var sorted []int
	var queue []int
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

// Find the middle page of a correctly ordered update
func middlePage(update []int) int {
	return update[len(update)/2]
}

func part1(rules map[int][]int, updates [][]int) {
	sum := 0

	for _, update := range updates {
		if isValidUpdate(update, rules) {
			sum += middlePage(update)
		}
	}
	fmt.Println("Part 1:", sum)
}

func part2(rules map[int][]int, updates [][]int) {
	sum := 0

	for _, update := range updates {
		if !isValidUpdate(update, rules) {
			corrected := reorderUpdate(update, rules)
			sum += middlePage(corrected)
		}
	}
	fmt.Println("Part 2:", sum)
}

func main() {
	rules, updates, _ := parseInput()
	part1(rules, updates)
	part2(rules, updates)
}
