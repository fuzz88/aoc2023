package main

import (
	"fmt"
	"math"
	"os"
	"slices"
)

func parse_line(start int, end int, raw_data *[]byte) []int {
	/* converts numbers from byte string representation to slice of integers */
	result := make([]int, 0)
	num := 0
	pow := 0.0
	// fmt.Println(start, end)
	for i := end; i >= start; i-- {
		if (*raw_data)[i] >= 48 && (*raw_data)[i] <= 57 {
			num = num + int(math.Pow(10, pow))*int((*raw_data)[i]-48)
			pow++
		}
		if (*raw_data)[i] == 32 || (i == start && (*raw_data)[i] >= 48 && (*raw_data)[i] <= 57) {
			result = append(result, num)
			num = 0
			pow = 0.0
		}
	}
	slices.Reverse(result)
	return result

}

func convert_seeds(seeds *[]int, conv *[]int) {
	fmt.Println("conv")
}

func main() {
	raw_data, err := os.ReadFile("test1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(raw_data)

	var seeds []int

	first_line := true
	prev_start := 0

	for i := 0; i < len(raw_data); i++ {
		if (raw_data[i] == 10 && !first_line && raw_data[i-1] != 10) || i == len(raw_data)-1 {
			// maps, empty line skipped
			nums := parse_line(prev_start, i, &raw_data)
			fmt.Println(nums)
			convert_seeds(&seeds, &nums)
			prev_start = i + 1

		}
		if raw_data[i] == 10 && (raw_data[i-1] < 48 || raw_data[i-1] > 57) {
			// it is map description line, skip to next conversion map
			i = i + 2
			for ; raw_data[i] != 10; i++ {
			}
			prev_start = i + 1
		}
		if raw_data[i] == 10 && first_line {
			// seeds at first line
			seeds = parse_line(0, i, &raw_data)
			fmt.Println(seeds)
			first_line = false
			prev_start = i + 1
		}

	}

}
