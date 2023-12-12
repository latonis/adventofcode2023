package main

import (
	"bufio"
	"fmt"
	"os"
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

type Universe struct {
	grid [][]Point
}

type Point struct {
	galaxy bool
	x int
	y int
}


func (universe *Universe) insert(index int, value []Point) {
	if len(universe.grid) == index { // nil or empty slice or after last element
		universe.grid = append(universe.grid, value)
		return
	}
	universe.grid = append(universe.grid[:index+1], universe.grid[index:]...) // index < len(a)
	universe.grid[index] = value
}

func insertPoint(a []Point, index int, value Point) []Point {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func (universe *Universe) Expand(count int) {

	for y := 0; y < len(universe.grid); y++ {
		emptyLine := true
		for x := 0; x < len(universe.grid[y]); x++ {
			if universe.grid[y][x].galaxy {
				emptyLine = false
			}
		}
		if emptyLine {
			for i := 0; i < count; i++ {
				universe.insert(y, universe.grid[y])
				y += 1
			}
		}
	}

	x := 0
	maxLen := len(universe.grid[0])

	for x < maxLen {
		emptyLine := true
		for y := 0; y < len(universe.grid); y++ {
			if universe.grid[y][x].galaxy {
				emptyLine = false
			}
		}

		if emptyLine {
			for i := 0; i < count; i++ {
				for y := 0; y < len(universe.grid); y++ {
					universe.grid[y] = insertPoint(universe.grid[y], x, universe.grid[y][x])
				}
				x += 1
				maxLen += 1
			}
		}
		x += 1
	}
}

func Sum(a []int) int {
	sum := 0
	for _, val := range a {
		sum += val
	}
	return sum
}

func (universe *Universe) PrintGrid() {
	for _, line := range universe.grid {
		for _, point := range line {
			if point.galaxy {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func calcPath(cX int, cY int, sX int, sY int) int {
	if cX < sX {
		return ((sX - cX) + (cY - sY))
	}
	return ((cX - sX) + (cY - sY))
}

func (universe Universe) traverse(xIndex int, yIndex int) int {
	paths := make([]int, 0)

	for y := yIndex; y < len(universe.grid); y++ {
		if y == yIndex {
			for x := xIndex + 1; x < len(universe.grid[y]); x++ {
				if universe.grid[y][x].galaxy {
					paths = append(paths, calcPath(x, y, xIndex, yIndex))
				}
			}
		} else {
			for x := 0; x < len(universe.grid[y]); x++ {
				if universe.grid[y][x].galaxy {
					paths = append(paths, calcPath(x, y, xIndex, yIndex))
				}
			}
		}
	}

	return Sum(paths)
}

func (universe Universe) FindPaths() int {
	total := 0
	for y, line := range universe.grid {
		for x, Point := range line {
			if Point.galaxy {
				total += universe.traverse(x, y)
			}
		}
	}
	return total
}

func solvePartOne(input []string) int {
	grid := new(Universe)

	for _, line := range input {
		gridX := make([]Point, 0)
		for _, char := range line {
			current := Point{galaxy: false}
			if char == '#' {
				current.galaxy = true
			}
			gridX = append(gridX, current)
		}
		grid.grid = append(grid.grid, gridX)
	}

	grid.Expand(1)
	total := grid.FindPaths()
	fmt.Println(total)
	return total
}

func solvePartTwo(input []string) int {
	grid := new(Universe)

	for y, line := range input {
		gridX := make([]Point, 0)
		for x, char := range line {
			current := Point{x: x, y: y, galaxy: false}
			if char == '#' {

				current.galaxy = true
			}
			gridX = append(gridX, current)
		}
		grid.grid = append(grid.grid, gridX)
	}

	grid.Expand(999999)
	total := grid.FindPaths()
	fmt.Println(total)
	return total
}
