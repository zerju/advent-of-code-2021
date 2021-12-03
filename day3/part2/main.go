// https://adventofcode.com/2021/day/3#part2
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

	var lines []string
	var mostCommon []int = nil
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if mostCommon == nil {
			mostCommon = make([]int, len(line))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	oxyBinary := searchForBinaryNumber(0, lines, mostCommon, true)
	co2Binary := searchForBinaryNumber(0, lines, mostCommon, false)

	oxyDecimal, _ := strconv.ParseInt(oxyBinary, 2, 64)
	co2Decimal, _ := strconv.ParseInt(co2Binary, 2, 64)

	fmt.Printf("Answer is: %v \n", oxyDecimal*co2Decimal)
}

func searchForBinaryNumber(arrayDepthIndex int, lines []string, common []int, searchMostCommon bool) string {
	var addend int = 1
	if !searchMostCommon {
		addend *= -1
	}
	for _, line := range lines {
		if line[arrayDepthIndex] == '0' {
			common[arrayDepthIndex] -= addend
		} else {
			common[arrayDepthIndex] += addend
		}
	}
	var nextLines []string
	for _, line := range lines {
		if line[arrayDepthIndex] == '0' && common[arrayDepthIndex] < 0 {
			nextLines = append(nextLines, line)
		} else if line[arrayDepthIndex] == '1' && common[arrayDepthIndex] >= 0 {
			nextLines = append(nextLines, line)
		}
	}
	if len(nextLines) > 1 {
		return searchForBinaryNumber(arrayDepthIndex+1, nextLines, common, searchMostCommon)
	}
	return nextLines[0]
}
