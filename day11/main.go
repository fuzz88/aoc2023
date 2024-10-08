package main

import (
	"bufio"
	"fmt"
	"os"
)

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

type Image [][]rune
type Point struct {
	x int
	y int
}
type Galaxy struct {
	Point
}

func readImageFromFile(filePath string) Image {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result Image
	for scanner.Scan() {
		result = append(result, []rune(scanner.Text()))
	}
	return result
}

func solve(image Image, expansion_coeff int) int {
	var galaxies []Galaxy
	for row, line := range image {
		for col, item := range line {
			if item == '#' {
				galaxy := Galaxy{Point: Point{x: row, y: col}}
				galaxies = append(galaxies, galaxy)
			}
		}
	}
	fmt.Printf("galaxies count: %v\n", len(galaxies))
	var horizontal_gaps map[int]bool = make(map[int]bool)
	var vertical_gaps map[int]bool = make(map[int]bool)
	for _, galaxy := range galaxies {
		vertical_gaps[galaxy.x] = true
		horizontal_gaps[galaxy.y] = true
	}
	total_distance := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			A := galaxies[i]
			B := galaxies[j]
			distance := absInt(A.x-B.x) + absInt(A.y-B.y)
			for col_idx := min(A.x, B.x); col_idx < max(A.x, B.x); col_idx++ {
				if !vertical_gaps[col_idx] {
					distance = distance + expansion_coeff
				}
			}
			for row_idx := min(A.y, B.y); row_idx < max(A.y, B.y); row_idx++ {
				if !horizontal_gaps[row_idx] {
					distance = distance + expansion_coeff
				}
			}
			total_distance = total_distance + distance
		}
	}
	return total_distance
}

func main() {
	fmt.Printf("\nAOC-2023 Day11 Solution\n\n")

	args := os.Args[1:]

	for _, filePath := range args {
		fmt.Println(filePath)
		image := readImageFromFile(filePath)
		fmt.Println("lines count: ", len(image))
		fmt.Println("Part1 :", solve(image, 1))
		fmt.Println("Part2 :", solve(image, 999999))
	}
}
