package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// Read the input file
	file, err := os.Open("./5/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rules []string
	var updates [][]int
	isUpdateSection := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isUpdateSection = true
			continue
		}
		if isUpdateSection {
			updates = append(updates, parseUpdate(line))
		} else {
			rules = append(rules, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Process the updates
	sumOfMiddlePages := 0
	sumOfMiddlePagesIncorrect := 0
	for _, update := range updates {
		if isCorrectOrder(update, rules) {
			middlePage := update[len(update)/2]
			sumOfMiddlePages += middlePage
		} else {
			correctedUpdate := correctOrder(update, rules)
			middlePage := correctedUpdate[len(correctedUpdate)/2]
			sumOfMiddlePagesIncorrect += middlePage
		}
	}

	fmt.Println("Sum of middle pages of correctly ordered updates:", sumOfMiddlePages)
	fmt.Println("Sum of middle pages of corrected updates:", sumOfMiddlePagesIncorrect)
	fmt.Println("Execution time:", time.Since(start))
}

func parseUpdate(line string) []int {
	parts := strings.Split(line, ",")
	update := make([]int, len(parts))
	for i, part := range parts {
		fmt.Sscanf(part, "%d", &update[i])
	}
	return update
}

func isCorrectOrder(update []int, rules []string) bool {
	pageIndex := make(map[int]int)
	for i, page := range update {
		pageIndex[page] = i
	}

	for _, rule := range rules {
		var x, y int
		fmt.Sscanf(rule, "%d|%d", &x, &y)
		xIndex, xExists := pageIndex[x]
		yIndex, yExists := pageIndex[y]
		if xExists && yExists && xIndex > yIndex {
			return false
		}
	}
	return true
}

func correctOrder(update []int, rules []string) []int {
	// Create a map to store the dependencies
	dependencies := make(map[int][]int)
	for _, rule := range rules {
		var x, y int
		fmt.Sscanf(rule, "%d|%d", &x, &y)
		dependencies[x] = append(dependencies[x], y)
	}

	// Use a topological sort to order the pages correctly
	ordered := topologicalSort(update, dependencies)
	return ordered
}

func topologicalSort(update []int, dependencies map[int][]int) []int {
	inDegree := make(map[int]int)
	for _, page := range update {
		inDegree[page] = 0
	}
	for _, deps := range dependencies {
		for _, dep := range deps {
			inDegree[dep]++
		}
	}

	queue := []int{}
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	ordered := []int{}
	for len(queue) > 0 {
		page := queue[0]
		queue = queue[1:]
		ordered = append(ordered, page)
		for _, dep := range dependencies[page] {
			inDegree[dep]--
			if inDegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	// Ensure the ordered list contains all pages from the update
	pageSet := make(map[int]bool)
	for _, page := range ordered {
		pageSet[page] = true
	}
	for _, page := range update {
		if !pageSet[page] {
			ordered = append(ordered, page)
		}
	}

	return ordered
}
