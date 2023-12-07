package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

var cards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func tieHand(cards1 []string, cards2 []string, joker bool) int {
	if joker {
		Remove(cards, Index(cards, "J"))
		tmpCards := []string{"J"}

		cards = append(tmpCards, cards...)

	}
	for idx := range cards1 {
		if cards1[idx] != cards2[idx] {
			if Index(cards, cards1[idx]) > Index(cards, cards2[idx]) {
				return 1
			}
			return 2
		}
	}
	return 1
}

func compareCard(cards1 []string, cards2 []string, joker bool) int {
	f := determineHand

	if joker {
		f = determineHandJoker
	}

	score1 := f(cards1)
	score2 := f(cards2)
	if score1 == score2 {
		return tieHand(cards1, cards2, joker)
	} else if score1 < score2 {
		return 1
	} else {
		return 2
	}
}

// 5ofak = 6, 4ofak = 5, full house = 4, 3ofak = 3, twop = 2, one pair = 1, high card = 0
func determineHand(cards []string) int {
	countCards := make(map[string]int)

	for _, card := range cards {
		_, ok := countCards[card]
		if ok {
			countCards[card] = countCards[card] + 1
		} else {
			countCards[card] = 1
		}
	}

	v := make([]string, 0, len(countCards))

	for _, value := range countCards {
		v = append(v, strconv.Itoa(value))
	}

	howMany := len(countCards)

	if howMany == 1 {
		return 6
	} else if howMany == 2 {
		if Index(v, "4") != -1 {
			return 5
		} else {
			return 4
		}
	} else if howMany == 3 {
		if Index(v, "3") != -1 {
			return 3
		} else {
			return 2
		}
	} else if howMany == 4 {
		return 1
	}
	return 0
}

// 5ofak = 6, 4ofak = 5, full house = 4, 3ofak = 3, twop = 2, one pair = 1, high card = 0
func determineHandJoker(cards []string) int {
	jokerMode := false

	if Index(cards, "J") != -1 {
		jokerMode = true
	}

	countCards := make(map[string]int)

	for _, card := range cards {
		_, ok := countCards[card]
		if ok {
			countCards[card] = countCards[card] + 1
		} else {
			countCards[card] = 1
		}
	}
	v := make([]string, 0, len(countCards))

	for _, value := range countCards {
		v = append(v, strconv.Itoa(value))
	}

	howMany := len(countCards)
	if howMany == 1 {
		return 6
	} else if howMany == 2 {
		if Index(v, "4") != -1 {
			if jokerMode {
				return 6
			}
			return 5
		} else {
			if jokerMode {
				return 6
			}
			return 4
		}
	} else if howMany == 3 {
		if Index(v, "3") != -1 {
			if jokerMode {
				return 5
			}
			return 3
		} else {
			if jokerMode {
				if countCards["J"] == 2 {
					return 5
				}
				return 4
			}
			return 2
		}
	} else if howMany == 4 {
		if jokerMode {
			return 3
		}
		return 1
	}

	if jokerMode {
		return 1
	}

	return 0
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

func cardsStringToSlice(input string) []string {
	cards := make([]string, 0)

	for _, card := range input {
		cards = append(cards, string(card))
	}

	return cards
}

type Hand struct {
	value int
	bid   int
	hand  []string
}

func solvePartOne(input []string) int {
	cardsList := make([]Hand, 0)

	for _, line := range input {
		var cards_key = Split(line)[0]

		cards := cardsStringToSlice(cards_key)

		bid, _ := strconv.Atoi(Split(line)[1])
		hand := new(Hand)
		hand.value = determineHand(cards)
		hand.bid = bid
		hand.hand = cards
		cardsList = append(cardsList, *hand)
	}

	sort.Slice(cardsList, func(i, j int) bool {
		if cardsList[i].value != cardsList[j].value {
			return cardsList[i].value < cardsList[j].value
		} else {
			return compareCard(cardsList[i].hand, cardsList[j].hand, false) == 2
		}
	})

	total := 0

	for idx := range cardsList {
		total += cardsList[idx].bid * (idx + 1)
	}

	fmt.Println(total)
	return total
}

func solvePartTwo(input []string) int {
	cardsList := make([]Hand, 0)

	for _, line := range input {
		var cards_key = Split(line)[0]

		cards := cardsStringToSlice(cards_key)

		bid, _ := strconv.Atoi(Split(line)[1])
		hand := new(Hand)
		hand.value = determineHandJoker(cards)
		hand.bid = bid
		hand.hand = cards
		cardsList = append(cardsList, *hand)
	}

	sort.Slice(cardsList, func(i, j int) bool {
		if cardsList[i].value != cardsList[j].value {
			return cardsList[i].value < cardsList[j].value
		} else {
			return compareCard(cardsList[i].hand, cardsList[j].hand, true) == 2
		}
	})

	total := 0

	for idx := range cardsList {
		total += cardsList[idx].bid * (idx + 1)
	}
	fmt.Println(total)
	return total
}
