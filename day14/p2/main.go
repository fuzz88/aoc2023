package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rock struct {
	x int
	y int
}
type Cube struct {
	x int
	y int
}

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

func parseInput(field *[]byte, cols int, rows int) ([]Rock, []Cube) {
	rocks := make([]Rock, 0)
	cubes := make([]Cube, 0)

	for i := 0; i < len(*field); i++ {
		line_no := i / cols
		line_pos := i % cols
		switch (*field)[i] {
		case 'O':
			rocks = append(rocks, Rock{line_pos, line_no})
		case '#':
			cubes = append(cubes, Cube{line_pos, line_no})
		}
	}

	return rocks, cubes
}

func main() {
	fmt.Println("--- Day 14: Parabolic Reflector Dish ---")

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "ERROR: no input file as cli argument")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	cols, rows, field := readInput(inputFile)
	rocks, cubes := parseInput(&field, cols, rows)

	fmt.Println("Rocks: ", rocks)
	fmt.Println("Cubes: ", cubes)

}
