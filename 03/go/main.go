package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func main() {
	readFile, err := os.Open("../input")

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
	solvePartTwo(fileLines)
}

func isSymbol(c byte) bool {
	asciiSymbolStart := 33
	asciiSymbolEnd := 64
	numStart := 48
	numEnd := 57

	return (!(c <= byte(numEnd) && c >= byte(numStart))) && c >= byte(asciiSymbolStart) && c <= byte(asciiSymbolEnd) && rune(c) != '.'
}

func getNum(line string, idx int) int {
	numStr := ""

	if idx > 0 {
		if idx > 1 {
			if line[idx-1] != '.' && !isSymbol(line[idx-1]) {
				if unicode.IsDigit(rune(line[idx-2])) {
					numStr += string(line[idx-2])
				}
			}
		}
		if unicode.IsDigit(rune(line[idx-1])) {
			numStr += string(line[idx-1])
		}
	}

	numStr += string(line[idx])

	if idx < len(line)-1 {
		if unicode.IsDigit(rune(line[idx+1])) {
			numStr += string(line[idx+1])
		}
		if idx+1 < len(line)-1 && line[idx+1] != '.' && !isSymbol(line[idx+1]) {
			if unicode.IsDigit(rune(line[idx+2])) {
				numStr += string(line[idx+2])
			}
		}
	}

	tempVal, _ := strconv.Atoi(numStr)
	return tempVal

}

func makeGrid(grid [][]string, lines []string) {
	for _, line := range lines {
		idx := 0
		for idx < len(line) {
			for unicode.IsDigit(rune(line[idx])) {

			}
		}
	}
}
func solvePartOne(input []string) int {
	total := 0
	for line_idx, line := range input {
		sameNum := false
		for char_idx, char := range line {
			char_indexes := []int{char_idx - 1, char_idx, char_idx + 1}
			line_indexes := []int{line_idx - 1, line_idx, line_idx + 1}
			if unicode.IsDigit(char) && !sameNum {
				for _, l_idx := range line_indexes {
					for _, c_idx := range char_indexes {
						if l_idx >= 0 && c_idx >= 0 && l_idx < len(input) && c_idx < len(input[l_idx]) {
							if isSymbol(input[l_idx][c_idx]) {
								sameNum = true
								total += getNum(line, char_idx)
							}
						}
					}
				}
			}
			if line[char_idx] == '.' || isSymbol(line[char_idx]) {
				sameNum = false
			}
		}
	}
	println(total)
	return total
}

func solvePartTwo(input []string) int {
	total := 0
	for line_idx, line := range input {
		for char_idx, char := range line {
			char_indexes := []int{char_idx - 1, char_idx, char_idx + 1}
			line_indexes := []int{line_idx - 1, line_idx, line_idx + 1}
			if char == '*' {
				adjNums := make([]int, 0)
				for _, l_idx := range line_indexes {
					for _, c_idx := range char_indexes {
						if l_idx >= 0 && c_idx >= 0 && c_idx < len(input[l_idx]) && l_idx < len(input) {
							if unicode.IsDigit(rune(input[l_idx][c_idx])) {
								if !slices.Contains(adjNums, getNum(input[l_idx], c_idx)) {
									adjNums = append(adjNums, getNum(input[l_idx], c_idx))
								}
							}
						}
					}
				}
				if len(adjNums) == 2 {
					total += adjNums[0] * adjNums[1]
				}
			}
		}
	}

	println(total)
	return total
}
