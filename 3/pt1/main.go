package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	// Open the input file
	file, err := os.Open("./3/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Regular expression to match valid mul(X,Y) instructions
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	totalSum := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x, err1 := strconv.Atoi(match[1])
			y, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				totalSum += x * y
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the results and the time taken
	fmt.Printf("Total sum of multiplications: %d\n", totalSum)
	fmt.Printf("Time taken: %s\n", time.Since(start))
}
