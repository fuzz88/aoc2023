package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func readInput(filename string) (int, int, []byte) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fileSize := fileInfo.Size()

	data := make([]byte, 0, fileSize)
	var scanner = bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line...)
		lineCount++
	}

	return int(fileSize/int64(lineCount) - 1), lineCount, data
}

func makeColsChan(field *[]byte, cols int, rows int) chan []byte {
	output := make(chan []byte)

	go func() {
		result := make([][]byte, cols)

		for col := 0; col < cols; col++ {
			result[col] = make([]byte, 0, rows)
		}

		for i := 0; i < len(*field); i++ {
			col := i % cols
			result[col] = append(result[col], (*field)[i])
		}

		for col := 0; col < cols; col++ {
			output <- result[col]
		}
		close(output)
	}()

	return output
}

func countRocksNorth(field *[]byte, cols int, rows int) int {
	results := make(chan int)

	var wg sync.WaitGroup

	for col := range makeColsChan(field, cols, rows) {

		wg.Add(1)

		go func() {
			defer wg.Done()
			scores := 0
			score := rows
			for idx, item := range col {
				if item == 'O' {
					scores += score
					score -= 1
				}
				if item == '#' {
					score = rows - (idx + 1)
				}
			}
			// fmt.Println(col, scores)
			results <- scores
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for r := range results {
		sum += r
	}
	return sum
}

func main() {
	fmt.Println("--- Day 14: Parabolic Reflector Dish ---")

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "ERROR: no input file as cli argument")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	cols, rows, field := readInput(inputFile)

	fmt.Println(countRocksNorth(&field, cols, rows))
}
