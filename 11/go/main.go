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

type Universe struct {
	grid     [][]Point
	galaxies []Point
}

type Point struct {
	x      int
	y      int
	galaxy bool
}

func (universe *Universe) Expand() [][]int {
	indexes := make([][]int, 0)
	indexes = append(indexes, []int{})
	indexes = append(indexes, []int{})

	for y := 0; y < len(universe.grid); y++ {
		emptyLine := true
		for x := 0; x < len(universe.grid[y]); x++ {
			if universe.grid[y][x].galaxy {
				emptyLine = false
			}
		}
		if emptyLine {
			indexes[1] = append(indexes[1], y)
		}
	}

	x := 0
	for x < len(universe.grid[0]) {
		emptyLine := true
		for y := 0; y < len(universe.grid); y++ {
			if universe.grid[y][x].galaxy {
				emptyLine = false
			}
		}

		if emptyLine {
			indexes[0] = append(indexes[0], x)
		}
		x += 1
	}
	return indexes
}

func Sum(a []int) int {
	sum := 0
	for _, val := range a {
		sum += val
	}
	return sum
}

func calcPath(cX int, cY int, sX int, sY int, indexes [][]int, factor int) int {
	path := 0
	if cX < sX {
		cX, sX = sX, cX
	}

	for _, x := range indexes[0] {
		if sX < x && x < cX {
			path += 1 * factor
		}
	}

	for _, y := range indexes[1] {
		if sY < y && y < cY {
			path += 1 * factor
		}
	}

	path += (cX - sX) + (cY - sY)

	return path
}

func (universe Universe) FindPaths(indexes [][]int, factor int) int {
	total := 0
	for idx := 0; idx < len(universe.galaxies)-1; idx++ {
		g1 := universe.galaxies[idx]
		for idx2 := idx + 1; idx2 < len(universe.galaxies); idx2++ {
			g2 := universe.galaxies[idx2]
			total += calcPath(g2.x, g2.y, g1.x, g1.y, indexes, factor-1)
		}
	}
	return total
}

func solvePartOne(input []string) int {
	grid := new(Universe)

	for y, line := range input {
		gridX := make([]Point, 0)
		for x, char := range line {
			current := Point{x: x, y: y, galaxy: false}
			if char == '#' {
				current.galaxy = true
				grid.galaxies = append(grid.galaxies, current)
			}
			gridX = append(gridX, current)
		}
		grid.grid = append(grid.grid, gridX)
	}

	indexes := grid.Expand()
	total := grid.FindPaths(indexes, 2)
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
				grid.galaxies = append(grid.galaxies, current)

			}
			gridX = append(gridX, current)
		}
		grid.grid = append(grid.grid, gridX)
	}

	indexes := grid.Expand()
	total := grid.FindPaths(indexes, 1000000)
	fmt.Println(total)
	return total
}
