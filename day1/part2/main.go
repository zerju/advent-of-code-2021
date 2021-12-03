// https://adventofcode.com/2021/day/1#part2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []int64
	for scanner.Scan() {
		line, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, line)
	}

	var previousSum int64 = -1
	biggerCount := 0
	var currentSum int64 = 0
	numOfSums := 0
	for i := 0; i < len(lines)*3; i++ {
		linesIndex := numOfSums + (i % 3)

		if len(lines) < linesIndex+1 || len(lines) < linesIndex+2 {
			break
		}
		currentNum := lines[linesIndex]
		if err != nil {
			log.Fatal(err)
		}

		currentSum += currentNum

		if (i+1)%3 == 0 {
			if currentSum > previousSum {
				biggerCount++
			}
			previousSum = currentSum
			currentSum = 0
			numOfSums++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bigger than previous count is: %v \n", biggerCount)
}
