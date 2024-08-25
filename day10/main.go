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
			if surface[i][j] == 'S' {
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
			surface[row][col] = 'V'
			if current_tile != '.' {
				distances[width*row+col] = counter
				counter++
				if checkBounds(row-1, col) && (current_tile == '|' || current_tile == 'J' || current_tile == 'L' || current_tile == 'S') {
					next_tile := surface[row-1][col]
					if next_tile == 'F' || next_tile == '|' || next_tile == '7' {
						walkNextStepAndMarkDistance(row-1, col, counter)
					}
				}
				if checkBounds(row+1, col) && (current_tile == '|' || current_tile == '7' || current_tile == 'F' || current_tile == 'S') {
					next_tile := surface[row+1][col]
					if next_tile == 'J' || next_tile == 'L' || next_tile == '|' {
						walkNextStepAndMarkDistance(row+1, col, counter)
					}
				}
				if checkBounds(row, col-1) && (current_tile == '-' || current_tile == 'J' || current_tile == '7' || current_tile == 'S') {
					next_tile := surface[row][col-1]
					if next_tile == 'L' || next_tile == '-' || next_tile == 'F' {
						walkNextStepAndMarkDistance(row, col-1, counter)
					}
				}
				if checkBounds(row, col+1) && (current_tile == '-' || current_tile == 'F' || current_tile == 'L' || current_tile == 'S') {
					next_tile := surface[row][col+1]
					if next_tile == 'J' || next_tile == '-' || next_tile == '7' {
						walkNextStepAndMarkDistance(row, col+1, counter)
					}
				}
			}
		}
	}
	walkNextStepAndMarkDistance(start_row, start_col, 0)
	return ((slices.Max(distances) + 1) / 2)
}

func solve2(surface Surface, original_surface Surface) int {
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
			if surface[i][col] == 'V' {
				meet++
				break
			}
		}
		for i := row; checkBounds(i, col); i-- {
			if surface[i][col] == 'V' {
				meet++
				break
			}
		}
		for i := col; checkBounds(row, i); i++ {
			if surface[row][i] == 'V' {
				meet++
				break
			}
		}
		for i := col; checkBounds(row, i); i-- {
			if surface[row][i] == 'V' {
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

	checkParity := func(row int, col int) bool {
		parity_bottom := false
		parity_top := false
		for dy := col + 1; dy < width; dy++ {
			if surface[row][dy] == 'V' && (original_surface[row][dy] == '|' || original_surface[row][dy] == 'F' || original_surface[row][dy] == '7') {
				parity_bottom = !parity_bottom
			}
			if surface[row][dy] == 'V' && (original_surface[row][dy] == '|' || original_surface[row][dy] == 'J' || original_surface[row][dy] == 'L') {
				parity_top = !parity_top
			}
		}
		if parity_top && parity_bottom {
			return true
		}
		return false
	}
	var walkNextStepAndMarkOutside func(row int, col int)

	walkNextStepAndMarkOutside = func(row int, col int) {
		if checkBounds(row-1, col) {
			if surface[row-1][col] == '.' {
				surface[row-1][col] = 'O'
				walkNextStepAndMarkOutside(row-1, col)
			}
		}
		if checkBounds(row+1, col) {
			if surface[row+1][col] == '.' {
				surface[row+1][col] = 'O'
				walkNextStepAndMarkOutside(row-1, col)
			}
		}
		if checkBounds(row, col-1) {
			if surface[row][col-1] == '.' {
				surface[row][col-1] = 'O'
				walkNextStepAndMarkOutside(row, col-1)
			}
		}
		if checkBounds(row, col+1) {
			if surface[row][col+1] == '.' {
				surface[row][col+1] = 'O'
				walkNextStepAndMarkOutside(row, col+1)
			}
		}
	}
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if !inLoop(row, col) && surface[row][col] != 'O' {
				surface[row][col] = 'O'
				walkNextStepAndMarkOutside(row, col)
			}
		}
	}
	count := 0
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if surface[row][col] == '.' && checkParity(row, col) {
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
			if surface[row][col] != 'V' {
				surface[row][col] = '.'
			}
		}
	}
}

func restoreStartTile(grid Surface, x, y int) rune {

	north, south, east, west := false, false, false, false

	if x > 0 && (grid[x-1][y] == '|' || grid[x-1][y] == '7' || grid[x-1][y] == 'F') {
		north = true
	}

	if x < len(grid)-1 && (grid[x+1][y] == '|' || grid[x+1][y] == 'L' || grid[x+1][y] == 'J') {
		south = true
	}

	if y > 0 && (grid[x][y-1] == '-' || grid[x][y-1] == 'L' || grid[x][y-1] == 'F') {
		west = true
	}

	if y < len(grid[x])-1 && (grid[x][y+1] == '-' || grid[x][y+1] == '7' || grid[x][y+1] == 'J') {
		east = true
	}

	switch {
	case north && south:
		return '|'
	case east && west:
		return '-'
	case north && east:
		return 'L'
	case north && west:
		return 'J'
	case south && west:
		return '7'
	case south && east:
		return 'F'
	default:
		panic("wrong input")
	}
}

func main() {
	fmt.Printf("\nAOC-2023 Day10 Solution\n\n")

	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
		surface := readSurfaceFromFile(filePath)
		original_surface := make(Surface, len(surface))
		for i := 0; i < len(surface); i++ {
			original_surface[i] = make([]rune, len(surface[i]))
			copy(original_surface[i], surface[i])
		}
		fmt.Println("Part1: ", solve(surface))
		cleanSurface(surface)

		x, y := findStart(original_surface)
		original_surface[x][y] = restoreStartTile(original_surface, x, y)
		fmt.Println("Part2: ", solve2(surface, original_surface))

	}
}
