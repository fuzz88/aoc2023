package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) []byte {
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

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			data = append(data, line[i])
		}
	}

	return data
}

func main() {
	fmt.Println("--- Day 14: Parabolic Reflector Dish ---")

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "ERROR: no input file as cli argument")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	inputData := readInput(inputFile)

	fmt.Println(inputFile)
	fmt.Println(inputData)

}
