// https://adventofcode.com/2021/day/2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var horizontal int64 = 0
	var depth int64 = 0
	for scanner.Scan() {
		currentStep := scanner.Text()
		stepCommands := strings.Fields(currentStep)
		direction := stepCommands[0]
		units, _ := strconv.ParseInt(stepCommands[1], 10, 64)
		switch direction {
		case "forward":
			horizontal += units
		case "up":
			depth -= units
		case "down":
			depth += units
		}
		if err != nil {
			log.Fatal(err)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	answer := depth * horizontal
	fmt.Printf("Final answer is: %v \n", answer)
}
