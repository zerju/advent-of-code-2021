// https://adventofcode.com/2021/day/3
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

	var code []int = nil
	for scanner.Scan() {
		currentLine := scanner.Text()
		if code == nil {
			code = make([]int, len(currentLine))
		}
		if err != nil {
			log.Fatal(err)
		}
		for i, el := range currentLine {
			if el == '0' {
				code[i]--
			} else {
				code[i]++
			}
		}
	}
	binary := ""
	invertedBinary := ""
	for _, el := range code {
		if el > 0 {
			binary += "1"
			invertedBinary += "0"
		} else {
			binary += "0"
			invertedBinary += "1"
		}
	}
	decimal1, _ := strconv.ParseInt(binary, 2, 64)
	decimal2, _ := strconv.ParseInt(invertedBinary, 2, 64)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Answer is: %v \n", decimal1*decimal2)
}
