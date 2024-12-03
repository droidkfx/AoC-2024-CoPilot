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

	// Regular expressions to match valid mul(X,Y), do(), and don't() instructions
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	totalSum := 0
	mulEnabled := true // Initially, mul instructions are enabled

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check for do() and don't() instructions
		if doRe.MatchString(line) {
			mulEnabled = true
		} else if dontRe.MatchString(line) {
			mulEnabled = false
		}

		// Process mul instructions regardless of their position in the line
		matches := mulRe.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if mulEnabled {
				x, err1 := strconv.Atoi(match[1])
				y, err2 := strconv.Atoi(match[2])
				if err1 == nil && err2 == nil {
					totalSum += x * y
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the results and the time taken
	fmt.Printf("Total sum of enabled multiplications: %d\n", totalSum)
	fmt.Printf("Time taken: %s\n", time.Since(start))
}
