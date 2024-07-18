package main

import (
	"fmt"
	"math"
	"os"
	"slices"
)

func parse_line(start int, end int, raw_data *[]byte) []int {
	/* converts numbers from byte string representation to a slice of integers */
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

func convert_seeds(seeds []int, convs [][]int) {
	//fmt.Printf("seeds: %v\n\n", seeds)
	for i, seed := range seeds {
		for _, conv := range convs {
			//fmt.Printf("conv : %v\n\n", conv)
			if (seed >= conv[1]) && (seed <= (conv[1] + conv[2] - 1)) {
				seeds[i] = conv[0] - conv[1] + seed
				break
			}
		}
	}
}

func solveFilePart1(filePath string) {
	raw_data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// fmt.Println(raw_data)

	var seeds []int
	var nums [][]int

	first_line := true
	prev_start := 0

	for i := 0; i < len(raw_data); i++ {
		if (raw_data[i] == 10 && !first_line && raw_data[i-1] != 10) || i == len(raw_data)-1 {
			// maps, empty line skipped
			num := parse_line(prev_start, i, &raw_data)
			nums = append(nums, num)
			// fmt.Println(nums)
			prev_start = i + 1
		}
		if raw_data[i] == 10 && (raw_data[i-1] < 48 || raw_data[i-1] > 57) {
			// it is map description line, skip to next conversion map
			convert_seeds(seeds, nums)
			nums = nil
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

	convert_seeds(seeds, nums)
	fmt.Println("The answer of part1 is :", slices.Min(seeds))
	
}

type seedRange struct {
	start 	int
	length 	int
}



func solveFilePart2(filePath string) {
	
	raw_data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// fmt.Println(raw_data)

	var seeds []seedRange
	var nums [][]int

	first_line := true
	prev_start := 0

	for i := 0; i < len(raw_data); i++ {
		if (raw_data[i] == 10 && !first_line && raw_data[i-1] != 10) || i == len(raw_data)-1 {
			// maps, empty line skipped
			num := parse_line(prev_start, i, &raw_data)
			nums = append(nums, num)
			// fmt.Println(nums)
			prev_start = i + 1
		}
		if raw_data[i] == 10 && (raw_data[i-1] < 48 || raw_data[i-1] > 57) {
			// it is map description line, skip to next conversion map
			//convert_seeds(seeds, nums)
			nums = nil
			i = i + 2
			for ; raw_data[i] != 10; i++ {
			}
			prev_start = i + 1
		}
		if raw_data[i] == 10 && first_line {
			// seeds at first line
			seed_nums := parse_line(0, i, &raw_data)
			
			for i, seed := range seed_nums {
				if i%2 == 0 {
					seeds = append(seeds, seedRange{start: seed, length: seed_nums[i+1]})
				}
			}
			fmt.Println(seeds)

			first_line = false
			prev_start = i + 1
		}

	}

	//convert_seeds(seeds, nums)
	fmt.Println("The answer of part2 is :", nil)
	



}

func main() {
	for _, filePath := range os.Args[1:] {
		solveFilePart1(filePath)
		solveFilePart2(filePath)
	}
}
