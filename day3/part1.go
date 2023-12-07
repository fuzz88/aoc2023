package main

import (
	"fmt"
	"os"
)

func main() {
	raw_data, err := os.ReadFile("test1.txt")
	if err != nil {
		panic(err)
	}
	line_len := 0
	line_num := 0
	for i := 0; i < len(raw_data); i++ {
		if raw_data[i] == 10 || i == len(raw_data)-1 {
			if line_len == 0 {
				line_len = i
			}
			line_num++
			fmt.Println(line_len, line_num)
			checkLineParts(raw_data[(line_num-1)*line_len : line_num*line_len])
		}
	}
}

func checkLineParts(line_data []byte) {
	fmt.Println(line_data)
}
