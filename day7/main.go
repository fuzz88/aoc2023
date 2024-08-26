package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Hand struct {
	cards     string
	bid       int
	hand_type int
}

type ByCards []Hand
type ByCardsWithJoker []Hand

func identifyCardStructure(cards_structure []int) int {
	sort.Slice(cards_structure, func(i, j int) bool {
		return cards_structure[i] < cards_structure[j]
	})
	structureMap := map[string]int{
		"[5]":         7,
		"[1 4]":       6,
		"[2 3]":       5,
		"[1 1 3]":     4,
		"[1 2 2]":     3,
		"[1 1 1 2]":   2,
		"[1 1 1 1 1]": 1,

	}
	key := fmt.Sprint(cards_structure)
	if val, ok := structureMap[key]; ok {
		return val
	}
	return 0
}

func getHandTypeAsNum(cards string) int {
	var cards_counter = make(map[rune]int)
	for _, card := range cards {
		cards_counter[card]++
	}
	var cards_structure []int
	cards_structure = maps.Values(cards_counter)
	return identifyCardStructure(cards_structure)

}

func getHandTypeAsNumWithJoker(cards string) int {
	joker_count := 0
	var cards_counter = make(map[rune]int)
	for _, card := range cards {
		if card == 'J' {
			joker_count++
		} else {
			cards_counter[card]++
		}
	}
	var cards_structure []int
	cards_structure = maps.Values(cards_counter)
	if joker_count == 5 {
		cards_structure = append(cards_structure, joker_count)
	} else {
		max_value := slices.Max(cards_structure)
		max_index := slices.Index(cards_structure, max_value)
		cards_structure[max_index] = cards_structure[max_index] + joker_count
	}
	return identifyCardStructure(cards_structure)
}

func (a ByCards) Len() int      { return len(a) }
func (a ByCards) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCards) Less(i, j int) bool {
	cardOrder := "J23456789TQKA"
	if a[i].hand_type != a[j].hand_type {
		return a[i].hand_type < a[j].hand_type
	} else {
		for t := 0; t < 5; t++ {
			if a[i].cards[t] != a[j].cards[t] {
				return strings.IndexByte(cardOrder, a[i].cards[t]) < strings.IndexByte(cardOrder, a[j].cards[t])
			}
		}
		return true
	}
}
func (a ByCardsWithJoker) Len() int      { return len(a) }
func (a ByCardsWithJoker) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCardsWithJoker) Less(i, j int) bool {
	cardOrder := "J23456789TQKA"
	if a[i].hand_type != a[j].hand_type {
		return a[i].hand_type < a[j].hand_type
	} else {
		for t := 0; t < 5; t++ {
			if a[i].cards[t] != a[j].cards[t] {
				return strings.IndexByte(cardOrder, a[i].cards[t]) < strings.IndexByte(cardOrder, a[j].cards[t])
			}
		}
		return true
	}
}

func readHandsFromFile(filePath string, withJoker bool) ([]Hand, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var result []Hand

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		bid, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, err
		}
		var hand Hand
		if !withJoker {
			hand = Hand{
				cards:     values[0],
				bid:       bid,
				hand_type: getHandTypeAsNum(values[0]),
			}
		} else {
			hand = Hand{
				cards:     values[0],
				bid:       bid,
				hand_type: getHandTypeAsNumWithJoker(values[0]),
			}
		}
		result = append(result, hand)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func solvePart1(filePath string) int {
	withJoker := false
	hands, err := readHandsFromFile(filePath, withJoker)
	if err != nil {
		panic(err)
	}
	sort.Sort(ByCards(hands))
	total_win := 0
	for i, v := range hands {
		t := i + 1
		total_win = total_win + t*v.bid
	}
	return total_win
}
func solvePart2(filePath string) int {
	withJoker := true
	hands, err := readHandsFromFile(filePath, withJoker)
	if err != nil {
		panic(err)
	}
	sort.Sort(ByCardsWithJoker(hands))
	total_win := 0
	for i, v := range hands {
		t := i + 1
		total_win = total_win + t*v.bid
	}
	return total_win
}

func main() {
	fmt.Printf("\nAOC_2023 Day7 Solution\n\n")

	args := os.Args[1:] // skip program filename
	for _, filePath := range args {
		fmt.Println("input file:", filePath)
		fmt.Println("part1: ", solvePart1(filePath))
		fmt.Println("part2: ", solvePart2(filePath))
	}
}
