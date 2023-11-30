package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("./input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	solvePartOne(fileLines)
	// solvePartTwo(fileLines)
}

func solvePartOne(input []string) int {
	for _, line := range input {
		fmt.Println(line)
	}

	return 0
}

func solvePartTwo(input []string) int {
	for _, line := range input {
		fmt.Println(line)
	}

	return 0
}
