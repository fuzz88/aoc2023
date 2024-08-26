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

func solvePart1(image Image) int {
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
			galaxy_A := galaxies[i]
			galaxy_B := galaxies[j]
			distance := absInt(galaxy_A.x-galaxy_B.x) + absInt(galaxy_A.y-galaxy_B.y)
			for t := min(galaxy_A.x, galaxy_B.x); t < max(galaxy_A.x, galaxy_B.x); t++ {
				if !vertical_gaps[t] {
					distance++
				}
			}
			for t := min(galaxy_A.y, galaxy_B.y); t < max(galaxy_A.y, galaxy_B.y); t++ {
				if !horizontal_gaps[t] {
					distance++
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
		fmt.Println("Part1 :", solvePart1(image))
	}
}
