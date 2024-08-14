package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("\nAOC-2023 Day8 Solution\n\n")
	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
	}
}
