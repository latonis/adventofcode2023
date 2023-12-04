package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func parseNums(input []string) []int {
	a := make([]int, 0)

	for item := range input {
		parsed, _ := strconv.Atoi(input[item])
		a = append(a, parsed)
	}
	return a
}

func solvePartOne(input []string) int {
	total := 0
	for _, line := range input {
		card_total := 0
		game_setp := strings.Split(line, ": ")
		nums_split := strings.Split(game_setp[1], " | ")
		winning_nums := parseNums(strings.Fields(nums_split[0]))
		given_nums := parseNums(strings.Fields(nums_split[1]))
		count_in := 0

		for _, given_num := range given_nums {
			if slices.Contains(winning_nums, given_num) {
				if count_in == 0 {
					card_total += 1
				} else {
					card_total *= 2
				}
				count_in += 1
			}
		}
		total += card_total
	}
	fmt.Println(total)
	return total
}

type CardCount struct {
	matches int
	count   int
}

func solvePartTwo(input []string) int {
	resultMap := make(map[int]CardCount)
	total := 0
	for _, line := range input {
		game_setp := strings.Split(line, ": ")
		game_id, _ := strconv.Atoi(strings.Fields(game_setp[0])[1])
		nums_split := strings.Split(game_setp[1], " | ")
		winning_nums := parseNums(strings.Fields(nums_split[0]))
		given_nums := parseNums(strings.Fields(nums_split[1]))
		count_in := 0

		for _, given_num := range given_nums {
			if slices.Contains(winning_nums, given_num) {
				count_in += 1
			}
		}

		resultMap[game_id] = CardCount{matches: count_in, count: resultMap[game_id].count + 1}

		for x := 1; x <= resultMap[game_id].count; x++ {
			for i := 1; i <= count_in; i++ {
				resultMap[game_id+i] = CardCount{matches: resultMap[game_id+i].matches, count: resultMap[game_id+i].count + 1}
			}
		}

		total += resultMap[game_id].count
	}
	fmt.Println(total)
	return total
}
