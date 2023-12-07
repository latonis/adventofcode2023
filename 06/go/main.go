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

func advance(race_time int, button_time int) int {
	movement_time := race_time - button_time

	movement_distance := movement_time * button_time

	return movement_distance
}

func find_start(time int, dist int) int {
	for i := 0; i <= time; i++ {
		if advance(time, i) > dist {
			return i
		}
	}
	return -1
}

func find_end(time int, dist int) int {
	for i := time; i >= 0; i-- {
		if advance(time, i) > dist {
			return i
		}
	}
	return -1
}

func solvePartOne(input []string) int {
	time_vals := parseNums(strings.Fields(strings.Split(input[0], ":")[1]))
	distance_vals := parseNums(strings.Fields(strings.Split(input[1], ":")[1]))

	total_mut := 1

	for idx := range time_vals {
		start := find_start(time_vals[idx], distance_vals[idx])
		fin := find_end(time_vals[idx], distance_vals[idx])
		total_mut *= fin + 1 - start
	}

	fmt.Println(total_mut)
	return total_mut
}

func solvePartTwo(input []string) int {
	time_vals := strings.Join(strings.Fields(strings.Split(input[0], ":")[1]), "")
	time_val, _ := strconv.Atoi(time_vals)
	distance_vals := strings.Join(strings.Fields(strings.Split(input[1], ":")[1]), "")
	distance_val, _ := strconv.Atoi(distance_vals)

	start := find_start(time_val, distance_val)
	fin := find_end(time_val, distance_val)

	fmt.Println(fin + 1 - start)
	return fin + 1 - start
}
