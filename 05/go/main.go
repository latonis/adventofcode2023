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
	defer solvePartTwo(fileLines)
}

func parseNums(input []string) []int {
	a := make([]int, 0)

	for item := range input {
		parsed, _ := strconv.Atoi(input[item])
		a = append(a, parsed)
	}
	return a
}

func Min(input []int) int {
	min := 0
	for idx, num := range input {
		if idx == 0 {
			min = num
		} else {
			if num < min {
				min = num
			}
		}
	}
	return min
}

type Almanac struct {
	maps []AlmanacMap
}

type AlmanacMap struct {
	rows []AlmanacRow
}

type AlmanacRow struct {
	destinationRange int
	sourceRange      int
	rangeLength      int
}

func (amap AlmanacMap) translate(seed int) int {
	for _, row := range amap.rows {
		if seed >= row.sourceRange && seed < row.sourceRange+row.rangeLength {
			return row.destinationRange + (seed - row.sourceRange)
		}
	}
	return seed
}

func (almanac Almanac) translate(seed int) int {
	for _, amap := range almanac.maps {
		seed = amap.translate(seed)
	}
	return seed
}

func parseMap(input []string) Almanac {
	almanac := new(Almanac)
	currMap := new(AlmanacMap)

	for idx, line := range input[1:] {
		if strings.Contains(line, ":") {
			currMap = new(AlmanacMap)
		} else if strings.Trim(line, "\n") != "" {
			currRow := new(AlmanacRow)
			inNums := parseNums(Split(line))
			currRow.destinationRange = inNums[0]
			currRow.sourceRange = inNums[1]
			currRow.rangeLength = inNums[2]
			currMap.rows = append(currMap.rows, *currRow)
			if idx == (len(input[1:]) - 1) {
				almanac.maps = append(almanac.maps, *currMap)
			}
		} else {
			almanac.maps = append(almanac.maps, *currMap)
		}
	}

	return *almanac
}

func solvePartOne(input []string) int {
	seeds := parseNums(Split(input[0])[1:])
	almanac := parseMap(input)

	min_val := int(^uint32(0))

	for _, seed := range seeds {
		min_val = min(almanac.translate(seed), min_val)
	}

	fmt.Println(min_val)
	return min_val
}

func solvePartTwo(input []string) int {
	seeds_parsed := parseNums(Split(input[0])[1:])
	c1 := make(chan int, len(seeds_parsed)/2)
	almanac := parseMap(input)

	for i := 0; i < len(seeds_parsed); i += 2 {
		go func(i int) {
			min_val := int(^uint32(0))
			for x := seeds_parsed[i]; x < seeds_parsed[i]+seeds_parsed[i+1]; x++ {
				min_val = min(almanac.translate(x), min_val)
			}
			c1 <- min_val
		}(i)
	}

	min_val := int(^uint32(0))
	for x := 0; x < len(seeds_parsed)/2; x++ {
		min_val = min(<-c1, min_val)
	}
	fmt.Println(min_val)
	return min_val
}
