package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

var (
	ColorMap = map[string]int{
		"red": 12, "blue": 14, "green": 13,
	}
)

func solvePartOne(input []string) int {
	totalIds := 0
	for _, line := range input {
		splitColon := strings.Split(line, ":")
		gameEntry := splitColon[0]
		entries := splitColon[1]
		invalid := false

		for _, game := range strings.Split(entries, ";") {
			colors := strings.Split(game, ",")
			for _, color := range colors {
				splitColor := strings.Fields(color)
				colorEntry := splitColor[1]
				colorVal, _ := strconv.Atoi(splitColor[0])

				if colorVal > ColorMap[colorEntry] {
					invalid = true
				}
			}
		}
		if !invalid {
			gameNum, _ := strconv.Atoi(strings.Fields(gameEntry)[1])
			totalIds += gameNum
		}
	}
	fmt.Println(totalIds)
	return totalIds
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func solvePartTwo(input []string) int {
	totalMut := 0
	var maxColor map[string]int

	for _, line := range input {
		maxColor = make(map[string]int)

		splitColon := strings.Split(line, ":")
		entries := splitColon[1]

		for _, game := range strings.Split(entries, ";") {
			colors := strings.Split(game, ",")
			for _, color := range colors {
				splitColor := strings.Fields(color)
				colorEntry := splitColor[1]
				colorVal, _ := strconv.Atoi(splitColor[0])
				curVal := maxColor[colorEntry]
				maxColor[colorEntry] = max(curVal, colorVal)
			}
		}
		mut := 1
		for _, val := range maxColor {
			mut *= val
		}
		totalMut += mut
	}

	fmt.Println(totalMut)
	return totalMut
}
