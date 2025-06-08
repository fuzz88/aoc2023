package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func makeInputChan(fileName string) <-chan string {
	inputChan := make(chan string)

	go func() {
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
	}()

	return inputChan
}

type Terrain struct {
	terrain    []rune
	lineLength int
	id         int
}

func parseInput(inputChan <-chan string) <-chan *Terrain {
	terrainChan := make(chan *Terrain)

	go func() {
		defer close(terrainChan)

		var terrain []rune
		lineCount := 0
		id := 0

		for line := range inputChan {
			if len(line) == 0 {
				id++
				terrainChan <- &Terrain{
					terrain:    terrain,
					lineLength: len(terrain) / lineCount,
					id:         id,
				}
				terrain = nil
				lineCount = 0
				continue
			}

			for _, ch := range line {
				terrain = append(terrain, ch)
			}
			lineCount++
		}
		// there is no empty line after last terrain,
		// so yield it when line iteration had stopped.
		id++
		terrainChan <- &Terrain{
			terrain:    terrain,
			lineLength: len(terrain) / lineCount,
			id:         id,
		}
	}()

	return terrainChan
}

func compareRows(row1 int, row2 int, t *Terrain) int {
	diff := 0
	for shift := 0; shift < t.lineLength; shift++ {
		idx1 := row1*t.lineLength + shift
		idx2 := row2*t.lineLength + shift
		if t.terrain[idx1] != t.terrain[idx2] {
			diff++
		}
	}
	return diff
}

func checkHorizontal(t *Terrain, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	found := 0
	row_total := len(t.terrain) / t.lineLength
	total_diff := 0

	for row := 0; row < row_total-1; row++ {
		next_row := row + 1
		diff := compareRows(row, next_row, t)
		if diff <= 1 {
			total_diff += diff
			found = row + 1

			first_row := row
			second_row := next_row
			for {
				first_row--
				second_row++
				if first_row >= 0 && second_row < row_total {
					diff = compareRows(first_row, second_row, t)
					if diff > 1 {
						found = 0
						break
					}
					total_diff += diff
				} else {
					break
				}
			}
		}

		if found != 0 && total_diff == 1 {
			result <- found * 100
			found = 0
		}
		total_diff = 0
	}
}

func compareCols(col1 int, col2 int, t *Terrain) int {
	diff := 0
	row_total := len(t.terrain) / t.lineLength
	for row := 0; row < row_total; row++ {
		row_shift := row * t.lineLength
		if t.terrain[row_shift+col1] != t.terrain[row_shift+col2] {
			diff++
		}
	}
	return diff
}

func checkVertical(t *Terrain, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	found := 0
	total_diff := 0

	for col := 0; col < t.lineLength-1; col++ {
		next_col := col + 1
		diff := compareCols(col, next_col, t)
		if diff <= 1 {
			total_diff += diff
			found = col + 1

			first_col := col
			second_col := next_col
			for {
				first_col--
				second_col++
				if first_col >= 0 && second_col < t.lineLength {
					diff = compareCols(first_col, second_col, t)
					if diff > 1 {
						found = 0
						break
					}
					total_diff += diff
				} else {
					break
				}
			}
		}
		if found != 0 && total_diff == 1 {
			result <- found
			found = 0
		}
		total_diff = 0
	}
}

func checkTerrainForMirrors(terrains <-chan *Terrain) chan int {
	result := make(chan int)

	var wg sync.WaitGroup
	for terrain := range terrains {
		wg.Add(2)
		go checkHorizontal(terrain, result, &wg)
		go checkVertical(terrain, result, &wg)
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	fmt.Println("--- Day 13: Point of Incidence ---")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./main inputFile\n")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	terrainsChan := parseInput(makeInputChan(inputFile))
	results := checkTerrainForMirrors(terrainsChan)

	var answer int
	for result := range results {
		answer += result
	}

	fmt.Println(answer)
}
