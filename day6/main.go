package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func read_data_from_file(filePath string) ([]Race, error) {
	file, err := os.Open(filePath)
	defer file.Close()
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

func read_single_race_from_file(filePath string) ([]Race, error) {
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
		s := ""
		for _, num := range nums[1:] {
			s = s + num
		}
		value, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var results []Race
	results = append(results, Race{time: data[0], distance: data[1]})
	return results, nil
}

func solve(races []Race) int {
	result := 1
	for i := 0; i < len(races); i++ {
		success_count := 0
		for speed := 0; speed < races[i].time; speed++ {
			time := races[i].time - speed
			distance := speed * time
			if distance > races[i].distance {
				success_count++
			}
		}
		result = result * success_count
	}
	return result
}

func main() {
	fmt.Printf("\nAOC_2023 Day6 Solution\n\n")

	args := os.Args[1:]
	for _, filePath := range args {
		data, err := read_data_from_file(filePath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("input: %v\n", data)
		fmt.Printf("Part 1: %d\n\n", solve(data))
		data, err = read_single_race_from_file(filePath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("input: %v\n", data)
		fmt.Printf("Part 2: %d\n\n", solve(data))
	}
}
