package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	s := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	numbers := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}

	lines := strings.Split(s, "\n")

	sum := 0
	for _, line := range lines {

		left := 0
		right := 0

		left_pos := len(line) + 2
		right_pos := -1

		for num, number := range numbers {
			num_pos := strings.Index(line, strconv.Itoa(num))
			number_pos := strings.Index(line, number)

			n_pos := -1

			if num_pos == -1 {
				if number_pos != -1 {
					n_pos = number_pos
				}
			} else {
				if number_pos != -1 {
					n_pos = slices.Min([]int{num_pos, number_pos})
				} else {
					n_pos = num_pos
				}
			}

			if left_pos > n_pos && n_pos != -1 {
				left = num
				left_pos = n_pos
				fmt.Println(num, n_pos, left_pos)
			}

		}
		fmt.Println()
		for num, number := range numbers {
			num_pos := strings.LastIndex(line, strconv.Itoa(num))
			number_pos := strings.LastIndex(line, number)

			n_pos := -1

			if num_pos == -1 {
				if number_pos != -1 {
					n_pos = number_pos
				}
			} else {
				if number_pos != -1 {
					n_pos = slices.Max([]int{num_pos, number_pos})
				} else {
					n_pos = num_pos
				}

			}

			if right_pos < n_pos && n_pos != -1 {
				right = num
				right_pos = n_pos
				fmt.Println(num, n_pos, right_pos)
			}

		}
		fmt.Println()
		fmt.Println(left, right)
		sum = sum + left*10 + right

	}
	fmt.Println(sum)
}
