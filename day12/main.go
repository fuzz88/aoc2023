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

// Recursive function to calculate the number of valid configurations
func countValidConfigurations(scheme []rune, groups []int, index int, groupIndex int) int {
	if groupIndex == len(groups) {
		// If all groups have been placed, make sure there are no more '#' to be placed
		c := 0
		gc := 0
		for i := 0; i < len(scheme); i++ {
			if scheme[i] == '#' {
				c++
			}
		}
		for _, k := range groups {
			gc = gc + k
		}
		if gc != c {
			return 0
		}
		fmt.Println(string(scheme), index)
		return 1
	}

	count := 0
	groupSize := groups[groupIndex]

	// Try to place the current group at every possible valid starting position
	for i := index; i <= len(scheme)-groupSize; i++ {
		valid := true
		if i == 0 || scheme[i-1] != '#' {
			// Check if the current segment can hold the group of `groupSize` broken springs
			for j := 0; j < groupSize; j++ {
				if scheme[i+j] == '.' {
					valid = false
					break
				}
			}
		} else {
			valid = false
		}

		if valid {
			// Temporarily place the group
			original := make([]rune, groupSize)
			copy(original, scheme[i:i+groupSize])
			for j := 0; j < groupSize; j++ {
				scheme[i+j] = '#'
			}

			// Ensure the next position after the group is operational (.)
			if i+groupSize < len(scheme) && scheme[i+groupSize] == '#' {
				valid = false
			}

			if valid {
				// Recursively place the next group
				count += countValidConfigurations(scheme, groups, i+groupSize+1, groupIndex+1)
			}

			// Restore the original state
			copy(scheme[i:i+groupSize], original)
		}
	}

	return count
}

func solve(records []Record) int {
	totalArrangements := 0
	for _, rec := range records {
		// Count the number of valid configurations for each row
		fmt.Println("-----------")
		fmt.Println(string(rec.scheme), rec.groups)
		count := countValidConfigurations(rec.scheme, rec.groups, 0, 0)
		fmt.Println(count)
		totalArrangements += count
	}
	fmt.Println()
	return totalArrangements
}

func main() {
	fmt.Printf("\nAOC-2023 Day12 Solution\n\n")

	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
		records := readRecordsFromFile(filePath)
		fmt.Println(solve(records))
	}
}
