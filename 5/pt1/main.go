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
	for _, update := range updates {
		if isCorrectOrder(update, rules) {
			middlePage := update[len(update)/2]
			sumOfMiddlePages += middlePage
		}
	}

	fmt.Println("Sum of middle pages:", sumOfMiddlePages)
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
