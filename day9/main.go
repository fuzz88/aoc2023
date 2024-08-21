package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readSeqsFromFile(filePath string) [][]int {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var result [][]int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var numbers []int
		for _, item := range fields {
			num, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
		}
		result = append(result, numbers)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}

func main() {
	fmt.Printf("\nAOC-2023 Day9 Solution\n\n")
	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
		fmt.Println(len(readSeqsFromFile(filePath)))
	}
}
