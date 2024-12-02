package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// Open the input file
	file, err := os.Open("./2/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	safeCount := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)
		if isSafe(levels) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the results and the time taken
	fmt.Printf("Number of safe reports: %d\n", safeCount)
	fmt.Printf("Time taken: %s\n", time.Since(start))
}

// Helper function to check if a report is safe
func isSafe(levels []string) bool {
	if len(levels) < 2 {
		return false
	}

	nums := make([]int, len(levels))
	for i, level := range levels {
		num, err := strconv.Atoi(level)
		if err != nil {
			return false
		}
		nums[i] = num
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < 1 || diff > 3 {
			return false
		}
		if nums[i] > nums[i-1] {
			decreasing = false
		} else if nums[i] < nums[i-1] {
			increasing = false
		} else {
			return false
		}
	}

	return increasing || decreasing
}
