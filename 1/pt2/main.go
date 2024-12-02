package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// Open the input file
	file, err := os.Open("./1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var leftList, rightList []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid input format")
			return
		}

		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	// Calculate the similarity score
	similarityScore := 0
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}

	// Print the results and the time taken
	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Similarity score: %d\n", similarityScore)
	fmt.Printf("Time taken: %s\n", time.Since(start))
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
