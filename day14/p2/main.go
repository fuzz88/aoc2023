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

func shiftNorth(field *[]byte, cols int, rows int) {
	f := *field
	for i := 0; i < len(f); i++ {
		if f[i] == 'O' {
			k := i / cols
			p := i % cols
			k = k - 1
			// fmt.Println(k, p)
			for k != -1 && f[k*cols+p] != '#' && f[k*cols+p] != 'O' {
				t := f[(k+1)*cols+p]
				f[(k+1)*cols+p] = f[k*cols+p]
				f[k*cols+p] = t
				k = k - 1
			}
		}
	}
}

func shiftSouth(field *[]byte, cols int, rows int) {
	f := *field
	for i := len(f) - cols - 1; i >= 0; i-- {
		if f[i] == 'O' {
			k := i / cols
			p := i % cols
			k = k + 1
			// fmt.Println(k, p)
			for k != rows && f[k*cols+p] != '#' && f[k*cols+p] != 'O' {
				t := f[(k-1)*cols+p]
				f[(k-1)*cols+p] = f[k*cols+p]
				f[k*cols+p] = t
				k = k + 1
			}
		}
	}
}

func shiftEast(field *[]byte, cols int, rows int) {
	f := *field
	for i := len(f) - 2; i >= 0; i-- {
		if f[i] == 'O' {
			k := i / cols
			p := i % cols
			p = p + 1
			// fmt.Println(k, p)
			for p != cols && f[k*cols+p] != '#' && f[k*cols+p] != 'O' {
				t := f[k*cols+p]
				f[k*cols+p] = f[k*cols+p-1]
				f[k*cols+p-1] = t
				p = p + 1
			}
		}
	}
}

func shiftWest(field *[]byte, cols int, rows int) {
	f := *field
	for i := 0; i < len(f); i++ {
		if f[i] == 'O' {
			k := i / cols
			p := i % cols
			p = p - 1
			// fmt.Println(k, p)
			for p != -1 && f[k*cols+p] != '#' && f[k*cols+p] != 'O' {
				t := f[k*cols+p]
				f[k*cols+p] = f[k*cols+p+1]
				f[k*cols+p+1] = t
				p = p - 1
			}
		}
	}
}

func printField(field *[]byte, cols int) {
	f := *field
	for i := 0; i < len(f); i++ {
		fmt.Printf("%c", f[i])
		if (i+1)%cols == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func total_load(field *[]byte, cols int, rows int) int {
	f := *field
	result := 0
	for i := 0; i < len(f); i++ {
		if f[i] == 'O' {
			k := i / cols
			k = rows - k
			result = result + k
		}
	}

	return result
}

func main() {
	fmt.Println("--- Day 14: Parabolic Reflector Dish ---")

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "ERROR: no input file as cli argument")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	cols, rows, field := readInput(inputFile)

	// fmt.Println(cols, rows, field)
	// printField(&field, cols)
	// fmt.Println(countRocksNorth(&field, cols, rows))
	// shiftNorth(&field, cols, rows)
	// printField(&field, cols)
	// shiftSouth(&field, cols, rows)
	// printField(&field, cols)
	// shiftEast(&field, cols, rows)
	// printField(&field, cols)
	// shiftWest(&field, cols, rows)
	// printField(&field, cols)

	// idk why 1_000 and 1_000_000_000 is in sync inside the cycle
	for i := 0; i < 1_000; i++ {
		shiftNorth(&field, cols, rows)
		shiftWest(&field, cols, rows)
		shiftSouth(&field, cols, rows)
		shiftEast(&field, cols, rows)
	}
	fmt.Println(total_load(&field, cols, rows))

	// this solution has a lot of issues. for example it is faster not to swap values so many times,
	// but find proper place for rock and swap one time with it
	
	// we can precalculate boundaries for each position, because cube-shaped rocks dont move. so we dont need to cycle through elements to find a place,
	// but can maybe somehow to lookup precalculated boundaries for a rock's position.

	// i am pretty sure, that we dont need to persist whole field, but only track boundaries, oval-rocks and cube-rocks to calculate state,
	// but this is too smart for me now. maybe there is proper data structure, but idk.

	// in a process i hoped that states a cycled (it is a low hanging fruit to assume that and check), so it is. count of iteration (1000) i guest from the first try, btw. hehe. if it wasn't 1000 i probably will be printing total_load for sample.txt, trying figure out cycle length or the way state cycling for input.txt. are they cycling the same way? idk.

	// in the end: approx about 10 min to run 1_000_000_000 iterations, because 1_000_000 takes 0.63sec
}
