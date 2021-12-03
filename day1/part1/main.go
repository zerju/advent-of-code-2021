// https://adventofcode.com/2021/day/1
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
	var previousNum int64 = -1
	biggerCount := 0
	for scanner.Scan() {
		currentNum, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if currentNum > previousNum && previousNum != -1 {
			biggerCount++
		}
		previousNum = currentNum
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bigger than previous count is: %v \n", biggerCount)
}
