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

func solve2(surface Surface) int {
	height := len(surface)
	width := len(surface[0])
	checkBounds := func(row int, col int) bool {
		if row < 0 || col < 0 || row > height-1 || col > width-1 {
			return false
		} else {
			return true
		}
	}
	inLoop := func(row int, col int) bool {
		meet := 0
		for i := row; checkBounds(i, col); i++ {
			if surface[i][col] == []rune("V")[0] {
				meet++
				break
			}
		}
		for i := row; checkBounds(i, col); i-- {
			if surface[i][col] == []rune("V")[0] {
				meet++
				break
			}
		}
		for i := col; checkBounds(row, i); i++ {
			if surface[row][i] == []rune("V")[0] {
				meet++
				break
			}
		}
		for i := col; checkBounds(row, i); i-- {
			if surface[row][i] == []rune("V")[0] {
				meet++
				break
			}
		}
		if meet < 4 {
			return false
		} else {
			return true
		}
	}
	var walkNextStepAndMarkOutside func(row int, col int)

	walkNextStepAndMarkOutside = func(row int, col int) {
		if checkBounds(row - 1, col) {
			if surface[row - 1][col] == []rune(".")[0] {
				surface[row - 1][col] = []rune("O")[0]
				walkNextStepAndMarkOutside(row - 1, col)
			}
		}
		if checkBounds(row + 1, col) {
			if surface[row + 1][col] == []rune(".")[0] {
				surface[row + 1][col] = []rune("O")[0]
				walkNextStepAndMarkOutside(row - 1, col)
			}
		}
		if checkBounds(row, col - 1) {
			if surface[row][col - 1] == []rune(".")[0] {
				surface[row][col - 1] = []rune("O")[0]
				walkNextStepAndMarkOutside(row, col - 1)
			}
		}
		if checkBounds(row, col + 1) {
			if surface[row][col + 1] == []rune(".")[0] {
				surface[row][col + 1] = []rune("O")[0]
				walkNextStepAndMarkOutside(row, col + 1)
			}
		}
	}
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if !inLoop(row, col) && surface[row][col] != []rune("O")[0] {
				surface[row][col] = []rune("O")[0]
				walkNextStepAndMarkOutside(row, col)
			}
		}
	}
	count := 0
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
		if surface[row][col] == []rune(".")[0] {
			count++	
			}
		}
	}

	return count
}

func cleanSurface(surface Surface) {
	height := len(surface)
	width := len(surface[0])
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if surface[row][col] != []rune("V")[0] {
				surface[row][col] = []rune(".")[0]
			}
		}
	}
}

func main() {
	fmt.Printf("\nAOC-2023 Day10 Solution\n\n")

	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
		surface := readSurfaceFromFile(filePath)
		for i:=0; i < len(surface); i++ {
			fmt.Println(string(surface[i]))
		}
		fmt.Println("Part1: ", solve(surface))
		cleanSurface(surface)
		for i:=0; i < len(surface); i++ {
			fmt.Println(string(surface[i]))
		}
		fmt.Println("Part2: ", solve2(surface))
	}
}
