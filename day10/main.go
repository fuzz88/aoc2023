package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Surface [][]rune

func readSurfaceFromFile(filePath string) Surface {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	var results Surface

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return results
}

func findStart(surface Surface) (int, int) {
	for i := 0; i < len(surface); i++ {
		for j := 0; j < len(surface[i]); j++ {
			if surface[i][j] == []rune("S")[0] {
				return i, j
			}
		}
	}
	panic("`S` tile is not found.")
}

func solve(surface Surface) int {
	start_row, start_col := findStart(surface)
	width := len(surface[0])
	height := len(surface)
	var distances []int = make([]int, width*height)
	checkBounds := func(row int, col int) bool {
		if row < 0 || col < 0 || row > height-1 || col > width-1 {
			return false
		} else {
			return true
		}
	}
	var walkNextStepAndMarkDistance func(row int, col int, counter int)
	walkNextStepAndMarkDistance = func(row int, col int, counter int) {
		if checkBounds(row, col) {
			current_tile := surface[row][col]
			surface[row][col] = []rune("V")[0]
			if current_tile != []rune(".")[0] {
				distances[width*row+col] = counter
				counter++
				if checkBounds(row-1, col) && (current_tile == []rune("|")[0] || current_tile == []rune("J")[0] || current_tile == []rune("L")[0] || current_tile == []rune("S")[0]) {
					next_tile := surface[row-1][col]
					if next_tile == []rune("F")[0] || next_tile == []rune("|")[0] || next_tile == []rune("7")[0] {
						walkNextStepAndMarkDistance(row-1, col, counter)
					}
				}
				if checkBounds(row+1, col) && (current_tile == []rune("|")[0] || current_tile == []rune("7")[0] || current_tile == []rune("F")[0] || current_tile == []rune("S")[0]) {
					next_tile := surface[row+1][col]
					if next_tile == []rune("J")[0] || next_tile == []rune("L")[0] || next_tile == []rune("|")[0] {
						walkNextStepAndMarkDistance(row+1, col, counter)
					}
				}
				if checkBounds(row, col-1) && (current_tile == []rune("-")[0] || current_tile == []rune("J")[0] || current_tile == []rune("7")[0] || current_tile == []rune("S")[0]) {
					next_tile := surface[row][col-1]
					if next_tile == []rune("L")[0] || next_tile == []rune("-")[0] || next_tile == []rune("F")[0] {
						walkNextStepAndMarkDistance(row, col-1, counter)
					}
				}
				if checkBounds(row, col+1) && (current_tile == []rune("-")[0] || current_tile == []rune("F")[0] || current_tile == []rune("L")[0] || current_tile == []rune("S")[0]) {
					next_tile := surface[row][col+1]
					if next_tile == []rune("J")[0] || next_tile == []rune("-")[0] || next_tile == []rune("7")[0] {
						walkNextStepAndMarkDistance(row, col+1, counter)
					}
				}
			}
		}
	}
	walkNextStepAndMarkDistance(start_row, start_col, 0)
	return ((slices.Max(distances) + 1) / 2)
}

func main() {
	fmt.Printf("\nAOC-2023 Day10 Solution\n\n")

	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
		surface := readSurfaceFromFile(filePath)
		fmt.Println("Part1: ", solve(surface))
	}
}
