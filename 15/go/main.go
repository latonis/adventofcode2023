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

func HASH(input string) int {
	iv := 0
	for _, c := range input {
		iv += int(c)
		iv *= 17
		iv = iv % 256
	}
	return iv
}

func solvePartOne(input []string) int {
	total := 0
	for _, line := range input {
		sequence := strings.Split(line, ",")
		for _, s := range sequence {
			val := HASH(s)
			total += val
		}
	}

	fmt.Println(total)
	return total
}

type label struct {
	id  string
	val int
}

func InitMap(hashmap map[int][]label) {
	MAPSIZE := 256
	for i := 0; i < MAPSIZE; i++ {
		hashmap[i] = make([]label, 0)
	}
}

func UpdateLabel(newLabel label, hashmap map[int][]label) {
	found := false
	for idx, oldLabel := range hashmap[HASH(newLabel.id)] {
		if oldLabel.id == newLabel.id {
			found = true
			hashmap[HASH(newLabel.id)][idx] = newLabel
		}
	}

	if !found {
		hashmap[HASH(newLabel.id)] = append(hashmap[HASH(newLabel.id)], newLabel)
	}
}

func RemoveLabel(removeLabel label, hashmap map[int][]label) {
	for idx, oldLabel := range hashmap[HASH(removeLabel.id)] {
		if oldLabel.id == removeLabel.id {
			hashmap[HASH(oldLabel.id)] = append(hashmap[HASH(oldLabel.id)][:idx], hashmap[HASH(oldLabel.id)][idx+1:]...)
		}
	}
}

func PrintMap(hashmap map[int][]label) {
	for k, v := range hashmap {
		if len(v) > 0 {
			fmt.Println(k, v)
		}
	}
}

func StringToVec(input string) []string {

	a := make([]string, 0)

	for _, s := range input {
		a = append(a, string(s))
	}

	return a
}

func solvePartTwo(input []string) int {
	hashmap := make(map[int][]label)
	total := 0
	InitMap(hashmap)

	for _, line := range input {
		sequence := strings.Split(line, ",")
		for _, s := range sequence {
			s_vec := StringToVec(s)
			op_idx := Index(s_vec, "=")

			if op_idx == -1 {
				op_idx = Index(s_vec, "-")
			}

			op := string(s[op_idx])
			label_id := strings.Split(s, op)[0]
			val, _ := strconv.Atoi(strings.Split(s, op)[1])

			if op == "=" {
				UpdateLabel(label{id: label_id, val: val}, hashmap)
			} else if op == "-" {
				RemoveLabel(label{id: label_id, val: val}, hashmap)
			} else {
				fmt.Println("Operation not supported: ", op)
			}
		}
	}

	for k, v := range hashmap {
		for idx, label := range v {
			total += (1 + k) * (idx + 1) * label.val
		}
	}

	fmt.Println(total)
	return total
}
