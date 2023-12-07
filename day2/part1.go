package main

import (
	"fmt"
	"os"
)

func checkPossible(start int, end int, raw_data *[]byte, sum chan int) {
	//sends game number to sum channel if possible,
	//or 0 when impossible
	fmt.Printf("%v\n", string((*raw_data)[start:end]))
	sum <- 1
}

func main() {
	sum := make(chan int)
	raw_data, err := os.ReadFile("test2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v bytes read.\n", len(raw_data))
	start := 0
	end := 0
	game_count := 0
	for i := 0; i < len(raw_data); i++ {
		if raw_data[i] == byte(10) {
			end = i
			go checkPossible(start, end, &raw_data, sum)
			game_count++
			start = end + 1
		}
		if i == len(raw_data)-1 {
			go checkPossible(start, i+1, &raw_data, sum)
			game_count++
		}
	}
	acc := 0
	for i := 0; i < game_count; i++ {
		acc = acc + <-sum
	}
	fmt.Printf("%v\n", acc)
}
