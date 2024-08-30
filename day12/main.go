package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	scheme []rune
	groups []int
}

func convert(list []string) []int {
	var result []int
	for _, item := range list {
		converted, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		result = append(result, converted)
	}
	return result
}

func parseLineAsRecord(line string) Record {
	fields := strings.Fields(line)
	return Record{
		scheme: []rune(fields[0]),
		groups: convert(strings.Split(fields[1], ",")),
	}
}
func readRecordsFromFile(filePath string) []Record {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var results []Record

	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, parseLineAsRecord(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return results
}

func main() {
	fmt.Printf("\nAOC-2023 Day12 Solution\n\n")

	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
		records := readRecordsFromFile(filePath)
		fmt.Println(len(records))
	}
}
