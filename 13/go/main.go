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
	// solvePartTwo(fileLines)
}

type Point struct {
	x int
	y int
	c string
}

type Pattern struct {
	grid [][]Point
}

func (pattern *Pattern) PrintGrid() {
	for _, line := range pattern.grid {
		for _, entry := range line {
			fmt.Print(entry.c)
		}
		fmt.Println()
	}
}

func CompareSlices(s1 []Point, s2 []Point) bool {
	if len(s1) != len(s2) {
		fmt.Println("Slices not the same length!")
		return false
	}

	for x := range s1 {
		if s1[x].c != s2[x].c {
			return false
		}
	}

	return true
}

func (pattern *Pattern) CheckVertical() int {
	maxLen := len(pattern.grid[0])
	for x := 0; x < maxLen-1; x++ {
		bad := false
		for x2 := 0; x2 < maxLen; x2++ {
			start := x - x2
			stop := x + x2 + 1
			if start >= 0 && start < stop && stop < maxLen {
				s1 := make([]Point, 0)
				s2 := make([]Point, 0)

				for y := 0; y < len(pattern.grid); y++ {
					s1 = append(s1, pattern.grid[y][start])
					s2 = append(s2, pattern.grid[y][stop])
				}

				if !CompareSlices(s1, s2) {
					bad = true
				}
			}
		}
		if !bad {
			return x + 1
		}
	}
	return -1
}

func (pattern *Pattern) CheckHorizontal() int {
	maxLen := len(pattern.grid)
	for y := 0; y < maxLen-1; y++ {
		bad := false
		for y2 := 0; y2 < maxLen; y2++ {
			start := y - y2
			stop := y + y2 + 1
			if start >= 0 && start < stop && stop < maxLen {
				if !CompareSlices(pattern.grid[start], pattern.grid[stop]) {
					bad = true
				}
			}
		}
		if !bad {
			return y + 1
		}
	}
	return -1
}

func (pattern *Pattern) FindMirror() int {
	total := 0
	pattern.PrintGrid()
	vert := pattern.CheckVertical()
	hori := pattern.CheckHorizontal()
	if vert != -1 {
		total += vert
	} else if hori != -1 {
		total += hori * 100
	}

	return total
}

func solvePartOne(input []string) int {
	total := 0
	pattern := new(Pattern)
	for y, line := range input {
		if len(line) == 0 {
			total += pattern.FindMirror()
			pattern.grid = make([][]Point, 0)
			continue
		}
		gridX := make([]Point, 0)
		for x, char := range line {
			gridX = append(gridX, Point{x, y, string(char)})
		}
		pattern.grid = append(pattern.grid, gridX)

		if y == len(input)-1 {
			total += pattern.FindMirror()
		}
	}

	fmt.Println(total)
	return 0
}

func solvePartTwo(input []string) int {
	for _, line := range input {
		fmt.Println(line)
	}

	return 0
}
