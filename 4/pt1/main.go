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

	// Define the word to search for
	word := "XMAS"
	wordLen := len(word)
	directions := [8][2]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // down-right
		{1, -1},  // down-left
		{0, -1},  // left
		{-1, 0},  // up
		{-1, 1},  // up-right
		{-1, -1}, // up-left
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Function to check if a word exists in a given direction
	checkWord := func(x, y, dx, dy int) bool {
		for i := 0; i < wordLen; i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != word[i] {
				return false
			}
		}
		return true
	}

	// Search for the word in all directions
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				if checkWord(i, j, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	// Print the result
	fmt.Printf("The word %s appears %d times\n", word, count)
	fmt.Printf("Execution time: %s\n", time.Since(start))
}
