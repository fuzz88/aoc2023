package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("\nAOC_2023 Day7 Solution\n\n")

	args := os.Args[1:] // skip program filename
	for _, filePath := range args {
		fmt.Println("input file:", filePath)
	}
}
