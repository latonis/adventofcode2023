package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fmt.Println(Index(strs, "pear"))
func Index(vs []layer, t string) int {
	for i, v := range vs {
		if v.id == t {
			return i
		}
	}
	return -1
}

// fmt.Println(Include(strs, "grape"))
// func Include(vs []string, t string) bool {
// 	return Index(vs, t) >= 0
// }

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

func StringToSlice(input string) []string {
	cards := make([]string, 0)

	for _, card := range input {
		cards = append(cards, string(card))
	}

	return cards
}

type layer struct {
	id string
	L  string
	R  string
}

func solvePartOne(input []string) int {
	var instructions []string
	var layers []layer
	for idx, line := range input {
		if idx == 0 {
			instructions = StringToSlice(line)
		} else if idx == 1 {

		} else {
			id := Split(line)[0]
			left := Split(line)[2][1:4]
			right := Split(line)[3][:3]
			layers = append(layers, layer{id, left, right})
		}
	}
	// found := false
	idx := 0
	curLayer := layers[Index(layers, "AAA")]
	count := 0

	for {

		if curLayer.id == "ZZZ" {
			fmt.Println(count)
			return count
		}
		ins := instructions[idx]

		if ins == "L" {
			curLayer = layers[Index(layers, curLayer.L)]
		} else {
			curLayer = layers[Index(layers, curLayer.R)]
		}

		count += 1
		if idx == len(instructions)-1 {
			idx = 0
		} else {
			idx += 1
		}
	}
}

// yoinked from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// yoinked from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func solvePartTwo(input []string) int {
	var instructions []string
	var startingLayers []layer
	var layers []layer

	for idx, line := range input {
		if idx == 0 {
			instructions = StringToSlice(line)
		} else if idx == 1 {

		} else {
			id := Split(line)[0]
			id_splice := StringToSlice(id)
			left := Split(line)[2][1:4]
			right := Split(line)[3][:3]
			if id_splice[len(id_splice)-1] == "A" {
				startingLayers = append(startingLayers, layer{id, left, right})
			}
			layers = append(layers, layer{id, left, right})
		}
	}

	c1 := make(chan int)
	for _, curLayer := range startingLayers {
		go func(c1 chan int, curLayer layer, instructions []string) int {
			idx := 0
			count := 0
			for {
				id_splice := StringToSlice(curLayer.id)
				if id_splice[len(id_splice)-1] == "Z" {
					c1 <- count
					return count
				}
				ins := instructions[idx]

				if ins == "L" {
					curLayer = layers[Index(layers, curLayer.L)]
				} else {
					curLayer = layers[Index(layers, curLayer.R)]
				}

				count += 1
				if idx == len(instructions)-1 {
					idx = 0
				} else {
					idx += 1
				}
			}
		}(c1, curLayer, instructions)
	}

	var counts = make([]int, 0)

	for i := 0; i < len(startingLayers); i++ {
		val := <-c1
		counts = append(counts, val)
	}

	fmt.Println(LCM(1, counts[0], counts[1:]...))
	return 0
}
