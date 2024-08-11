package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
)

type Hand struct {
	cards	string;
	bid		int;
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
		result = append(result, Hand{cards: values[0], bid: bid}) 
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
		fmt.Println(hands)
	}
}
