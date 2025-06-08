package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeInputChan(fileName string) <-chan string {
	inputChan := make(chan string)

	go func(fileName string, inputChan chan string) {
		defer close(inputChan)

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			text := scanner.Text()
			inputChan <- text
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}(fileName, inputChan)

	return inputChan
}

func main() {
	fmt.Println("--- Day 13: Point of Incidence ---")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./main inputFile\n")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	inputData := makeInputChan(inputFile)

	for line := range inputData {
		fmt.Println(line)
	}
}
