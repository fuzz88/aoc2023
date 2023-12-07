package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("test1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
