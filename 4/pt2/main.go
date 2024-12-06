package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	// Read the input file
	file, err := os.Open("./4/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Define the patterns to search for
	patterns := [][][2]int{
		{{0, 0}, {1, -1}, {1, 1}, {2, -2}, {2, 2}},     // M.A.S
		{{0, 0}, {1, 1}, {1, -1}, {2, 2}, {2, -2}},     // M.S.A
		{{0, 0}, {-1, -1}, {-1, 1}, {-2, -2}, {-2, 2}}, // S.A.M
		{{0, 0}, {-1, 1}, {-1, -1}, {-2, 2}, {-2, -2}}, // S.M.A
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Function to check if a pattern exists in a given direction
	checkPattern := func(x, y int, pattern [][2]int) bool {
		for _, p := range pattern {
			nx, ny := x+p[0], y+p[1]
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != 'M' && grid[nx][ny] != 'A' && grid[nx][ny] != 'S' {
				return false
			}
		}
		return true
	}

	// Search for the patterns in the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, pattern := range patterns {
				if checkPattern(i, j, pattern) {
					count++
				}
			}
		}
	}

	// Print the result
	fmt.Printf("The pattern X-MAS appears %d times\n", count)
	fmt.Printf("Execution time: %s\n", time.Since(start))
}
