package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func getNextValueForSeq(seq []int) int {
	var sub_seq []int
	dx := seq[0] - seq[1]
	if dx == 0 {
		return seq[0]
	}
	for i := 0; i < len(seq)-1; i++ {
		sub_seq = append(sub_seq, seq[i]-seq[i+1])
	}
	fmt.Println(sub_seq)
	return seq[0] + getNextValueForSeq(sub_seq)

}

func solve(seqs [][]int) int {
	result := 0
	for _, seq := range seqs {
		slices.Reverse(seq)
		fmt.Println(seq)
		value := getNextValueForSeq(seq)
		fmt.Println(value)
		result = result + value
	}
	return result
}

func main() {
	fmt.Printf("\nAOC-2023 Day9 Solution\n\n")
	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
		seqs := readSeqsFromFile(filePath)
		fmt.Println(solve(seqs))
	}
}
