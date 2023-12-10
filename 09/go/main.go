package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fmt.Println(Index(strs, "pear"))
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// fmt.Println(Include(strs, "grape"))
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//	fmt.Println(Any(strs, func(v string) bool {
//	    return strings.HasPrefix(v, "p")
//	}))
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//	fmt.Println(All(strs, func(v string) bool {
//	    return strings.HasPrefix(v, "p")
//	}))
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//	fmt.Println(Filter(strs, func(v string) bool {
//	    return strings.Contains(v, "e")
//	}))
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// fmt.Println(Map(strs, strings.ToUpper))
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Split(s string) []string {
	return strings.Fields(s)
}

func First(T []any) any {
	if len(T) > 0 {
		return T[0]
	}
	return nil
}

func Last(T []any) any {
	if len(T) > 0 {
		return T[len(T)-1]
	}
	return nil
}

func parseNums(input []string) []int {
	a := make([]int, 0)

	for item := range input {
		parsed, _ := strconv.Atoi(input[item])
		a = append(a, parsed)
	}
	return a
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

func AllZeroes(input []int) bool {
	if len(input) == 0 {
		return false
	}

	for _, value := range input {
		if value != 0 {
			return false
		}
	}
	return true
}

func solvePartOne(input []string) int {
	total := 0
	for _, line := range input {
		differences := make([][]int, 0)
		numbers := parseNums(Split(line))
		difference := make([]int, 0)
		differences = append(differences, numbers)

		flag := false
		for !flag {
			for idx := range numbers {
				if idx == 0 {
					continue
				}
				difference = append(difference, numbers[idx]-numbers[idx-1])
			}
			numbers = difference
			differences = append(differences, difference)
			flag = AllZeroes(difference)
			difference = make([]int, 0)
		}

		differences[len(differences)-1] = append(differences[len(differences)-1], 0)

		for i := len(differences) - 1; i > 0; i-- {
			currLayer := differences[i]
			nextLayer := differences[i-1]
			differences[i-1] = append(differences[i-1], currLayer[len(currLayer)-1]+nextLayer[len(nextLayer)-1])
		}
		total += differences[0][len(differences[0])-1]
	}

	fmt.Println(total)
	return total
}

func solvePartTwo(input []string) int {
	total := 0
	for _, line := range input {
		differences := make([][]int, 0)
		numbers := parseNums(Split(line))
		difference := make([]int, 0)
		differences = append(differences, numbers)

		flag := false
		for !flag {
			for idx := range numbers {
				if idx == 0 {
					continue
				}
				difference = append(difference, numbers[idx]-numbers[idx-1])
			}
			numbers = difference
			differences = append(differences, difference)
			flag = AllZeroes(difference)
			difference = make([]int, 0)
		}

		differences[len(differences)-1] = append(differences[len(differences)-1], 0)

		for i := len(differences) - 1; i > 0; i-- {
			currLayer := differences[i]
			nextLayer := differences[i-1]
			differences[i-1] = append([]int{nextLayer[0] - currLayer[0]}, differences[i-1]...)
		}
		total += differences[0][0]
	}

	fmt.Println(total)
	return total
}
