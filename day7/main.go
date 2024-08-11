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

func getHandTypeAsNum(cards string) int {
	var cards_counter = make(map[rune]int)
	for _, card := range cards {
		cards_counter[card]++
	}
	var cards_structure []int
	cards_structure = maps.Values(cards_counter)
	sort.Slice(cards_structure, func(i, j int) bool {
		return cards_structure[i] < cards_structure[j]
	})

	if slices.Equal(cards_structure, []int{5}) {
		return 7
	}
	if slices.Equal(cards_structure, []int{1, 4}) {
		return 6
	}
	if slices.Equal(cards_structure, []int{2, 3}) {
		return 5
	}
	if slices.Equal(cards_structure, []int{1, 1, 3}) {
		return 4
	}
	if slices.Equal(cards_structure, []int{1, 2, 2}) {
		return 3
	}
	if slices.Equal(cards_structure, []int{1, 1, 1, 2}) {
		return 2
	}
	if slices.Equal(cards_structure, []int{1, 1, 1, 1, 1}) {
		return 1
	}

	return 0
}

func (a ByCards) Len() int      { return len(a) }
func (a ByCards) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCards) Less(i, j int) bool {
	cardOrder := "23456789TJQKA"
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

func readHandsFromFile(filePath string) ([]Hand, error) {
	fmt.Println("input file:", filePath)
	file, err := os.Open(filePath)
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
		result = append(result, Hand{
			cards:     values[0],
			bid:       bid,
			hand_type: getHandTypeAsNum(values[0]),
		})
	}
	return result, nil
}

func main() {
	fmt.Printf("\nAOC_2023 Day7 Solution\n\n")

	args := os.Args[1:] // skip program filename
	for _, filePath := range args {
		hands, err := readHandsFromFile(filePath)
		if err != nil {
			panic(err)
		}
		sort.Sort(ByCards(hands))
		total_win := 0
		for i, v := range hands {
			t := i + 1
			total_win = total_win +  t * v.bid
		}
		fmt.Printf("Part 1: %v\n", total_win)
	}
}
