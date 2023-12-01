package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_SolvePartOne(t *testing.T) {
	readFile, err := os.Open("./test-input")

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

	result := solvePartOne(fileLines)
	expected := 0

	if result != expected {
		t.Errorf("Incorrect result! given: %d, expected: %d.", result, expected)
	}
}

func Test_SolvePartTwo(t *testing.T) {
	readFile, err := os.Open("./test-input")

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

	result := solvePartTwo(fileLines)
	expected := 0

	if result != expected {
		t.Errorf("Incorrect result! given: %d, expected: %d.", result, expected)
	}

}
