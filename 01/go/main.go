package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type LetterInt int

const (
	one LetterInt = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
)

var (
	LetterIntMap = map[string]LetterInt{
		"one": one, "two": two, "three": three, "four": four, "five": five, "six": six, "seven": seven, "eight": eight, "nine": nine,
	}
)

func ParseString(str string) (LetterInt, bool) {
	c, ok := LetterIntMap[strings.ToLower(str)]
	return c, ok
}

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

func solvePartOne(input []string) int {
	totalVal := 0
	for _, line := range input {
		var lineDigits []string
		for _, c := range line {
			if unicode.IsDigit(c) {
				num := string(c)
				lineDigits = append(lineDigits, num)
			}
		}
		val, _ := strconv.Atoi(lineDigits[0] + lineDigits[len(lineDigits)-1])
		totalVal += val
	}

	fmt.Println(totalVal)
	return totalVal
}

func trimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func findNum(line string) string {
	for i := range line {
		for j := range line[i:] {
			val, ok := LetterIntMap[line[i:len(line)-j]]
			if ok {
				return strings.Replace(line, line[i:len(line)-j], fmt.Sprint(int(val))+string(line[len(line)-j-1]), int(-1))
			}
		}
	}
	return line
}

func solvePartTwo(input []string) int {
	var totalVal = 0
	for _, line := range input {

		for i := 1; i <= len(LetterIntMap); i++ {
			line = findNum(line)
		}

		var lineDigits []string
		for _, c := range line {
			if unicode.IsDigit(c) {
				num := string(c)
				lineDigits = append(lineDigits, num)
			}
		}
		val, err := strconv.Atoi(lineDigits[0] + lineDigits[len(lineDigits)-1])

		if err != nil {
			fmt.Println(err)
		}

		totalVal += val
	}

	fmt.Println(totalVal)

	return totalVal

}
