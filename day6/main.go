package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

type Race struct {
	time     int
	distance int
}

func read_data_from_file(filePath string) ([]Race, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		nums := strings.Fields(text)
		for _, num := range nums {
			s, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			data = append(data, s)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var results []Race
	for i := 0; i < len(data)/2; i++ {
		results = append(results, Race{time: data[i], distance: data[len(data)/2+i]})
	}
	return results, nil
}

func solvePart1(races []Race) int {
	return 0
}

func main() {
	fmt.Printf("\nAOC_2023 Day6 Solution\n")

	args := os.Args[1:]
	for _, filePath := range args {
		data, err := read_data_from_file(filePath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("input: %v\n", data)
		fmt.Printf("Part 1: %d\n\n", solvePart1(data))
	}
}
