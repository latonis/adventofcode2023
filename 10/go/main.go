package main

import (
	"bufio"
	"fmt"
	"math"
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
	readFile, err := os.Open("../test-input")

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

func PrintGrid(grid [][]Point) {
	for _, line := range grid {
		for _, point := range line {
			fmt.Print(point.c)
		}
		fmt.Println()
	}
}

type Point struct {
	x int
	y int
	c string
}

func (p *Point) distance(p2 *Point) int {
	return int(math.Sqrt(float64(p2.x-p.x)*float64(p2.x-p.x) + float64(p2.y-p.y)*float64(p2.y-p.y)))
}

func (p *Point) grid(grid [][]Point) {
	for _, line := range grid {
		for _, point := range line {
			fmt.Print(p.distance(&point))
		}
		fmt.Println()
	}
}

func solvePartOne(input []string) int {
	grid := make([][]Point, 0)
	var start Point
	for y, line := range input {
		gridX := make([]Point, 0)
		for x, char := range line {
			if char == 'S' {
				start = Point{x: x, y: y, c: string(char)}
			}
			gridX = append(gridX, Point{x: x, y: y, c: string(char)})
		}
		grid = append(grid, gridX)
	}
	fmt.Println(start)
	start.grid(grid)
	return 0
}

func solvePartTwo(input []string) int {
	for _, line := range input {
		fmt.Println(line)
	}

	return 0
}
