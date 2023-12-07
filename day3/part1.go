package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	raw_data, err := os.ReadFile("test1.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(raw_data)
	sum := make(chan int, 1000)
	for i := 0; i < len(raw_data); i++ {
		line_start := 0
		if raw_data[i] == 10 || i == len(raw_data)-1 {
			checkLineParts(raw_data[line_start:i], sum)
			line_start = i
		}
	}
}

func checkLineParts(line_data []byte, sum chan int) {
	fmt.Println(string(line_data))
	for i := 0; i < len(line_data); i++ {
		if (line_data[i] < 48 || line_data[i] > 57) || i == len(line_data)-1 {
			pow := 0.0
			num := 0
			j := i - 1
			// fmt.Printf("%v %v", line_data[i], i)
			for ; j > -1 && line_data[j] >= 48 && line_data[j] <= 57; j-- {
				num = num + int(math.Pow(10, pow))*int((line_data[j]-48))
				pow++
			}
			if num > 0 && isConnectedNum(j+1, i-1) {
				fmt.Println(num)
				sum <- num
			}

		}
	}
}

func isConnectedNum(start_idx int, end_idx int) bool {
	fmt.Println(start_idx, end_idx)
	return true
}
