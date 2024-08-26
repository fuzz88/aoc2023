package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Surface [][]rune

func readSurfaceFromFile(filePath string) Surface {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var results Surface
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		results = append(results, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return results
}

func findStart(surface Surface) (int, int) {
	for i, row := range surface {
		for j, tile := range row {
			if tile == 'S' {
				return i, j
			}
		}
	}
	panic("`S` tile not found.")
}

func solve(surface Surface) int {
	start_row, start_col := findStart(surface)
	height := len(surface)
	width := len(surface[0])
	var distances []int = make([]int, width*height)

	checkBounds := func(row int, col int) bool {
		return row >= 0 && col >= 0 && row < height && col < width
	}

	var walkNextStepAndMarkDistance func(row int, col int, counter int)
	walkNextStepAndMarkDistance = func(row int, col int, counter int) {
		if !checkBounds(row, col) || surface[row][col] == 'V' {
			return
		}

		currentTile := surface[row][col]
		surface[row][col] = 'V'

		if currentTile == '.' {
			return
		}

		distances[width*row+col] = counter
		counter++

		directions := []struct {
			dr, dc    int
			validFrom string
			validTo   string
		}{
			{-1, 0, "|JLS", "F|7"},
			{1, 0, "|7FS", "JL|"},
			{0, -1, "-J7S", "L-F"},
			{0, 1, "-FLS", "J-7"},
		}

		for _, d := range directions {
			newRow, newCol := row+d.dr, col+d.dc
			if checkBounds(newRow, newCol) && strings.ContainsRune(d.validFrom, rune(currentTile)) {
				nextTile := surface[newRow][newCol]
				if strings.ContainsRune(d.validTo, rune(nextTile)) {
					walkNextStepAndMarkDistance(newRow, newCol, counter)
				}
			}
		}
	}

	walkNextStepAndMarkDistance(start_row, start_col, 0)
	return ((slices.Max(distances) + 1) / 2)
}

func solve2(surface Surface, original_surface Surface) int {
	width := len(surface[0])
	height := len(surface)

	checkBounds := func(row int, col int) bool {
		return row >= 0 && col >= 0 && row < height && col < width
	}

	inLoop := func(row int, col int) bool {
		directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		meet := 0
		for _, dir := range directions {
			for i, j := row, col; checkBounds(i, j); i, j = i+dir[0], j+dir[1] {
				if surface[i][j] == 'V' {
					meet++
					break
				}
			}
		}
		return meet == 4
	}

	checkParity := func(row int, col int) bool {
		parityTop, parityBottom := false, false

		for dy := col + 1; dy < width; dy++ {
			if surface[row][dy] == 'V' {
				switch original_surface[row][dy] {
				case '|':
					parityBottom = !parityBottom
					parityTop = !parityTop
				case 'F', '7':
					parityBottom = !parityBottom
				case 'J', 'L':
					parityTop = !parityTop
				}
			}
		}

		return parityTop && parityBottom
	}
	var walkNextStepAndMarkOutside func(row int, col int)
	walkNextStepAndMarkOutside = func(row int, col int) {
		directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if checkBounds(newRow, newCol) && surface[newRow][newCol] == '.' {
				surface[newRow][newCol] = 'O'
				walkNextStepAndMarkOutside(newRow, newCol)
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
	for row := range surface {
		for col := range surface[row] {
			if surface[row][col] != 'V' {
				surface[row][col] = '.'
			}
		}
	}
}

func restoreStartTile(grid Surface, x, y int) rune {
	north := x > 0 && (grid[x-1][y] == '|' || grid[x-1][y] == '7' || grid[x-1][y] == 'F')
	south := x < len(grid)-1 && (grid[x+1][y] == '|' || grid[x+1][y] == 'L' || grid[x+1][y] == 'J')
	west := y > 0 && (grid[x][y-1] == '-' || grid[x][y-1] == 'L' || grid[x][y-1] == 'F')
	east := y < len(grid[x])-1 && (grid[x][y+1] == '-' || grid[x][y+1] == '7' || grid[x][y+1] == 'J')

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
