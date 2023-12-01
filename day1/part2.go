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
	results := make([]string, len(lines))
	for _, line := range lines {
		min := -1
		for min != 10000 {
			indexes := make([]int, 9)
			for idx, number := range numbers {
				indexes[idx-1] = strings.Index(line, number)
				if indexes[idx-1] == -1 {
					indexes[idx-1] = 10000
				}
			}
			min = slices.Min(indexes)
			if min != 10000 {
				number := slices.Index(indexes, min)
				line = strings.Replace(line, numbers[number+1], strconv.Itoa(number+1), 1)
				fmt.Println(min, number+1, line)
			}

		}

		results = append(results, line)
	}

	fmt.Println(results)
	s = strings.Join(results, "\n")
	sum := 0
	t1 := -1
	t2 := -1
	for _, k := range s {

		if k == 10 {
			if t2 == -1 {
				t2 = t1
			}
			if t1 == -1 && t2 == -1 {
				continue
			}
			fmt.Printf("%v%v\n", t1, t2)
			sum = sum + t1*10 + t2
			fmt.Printf("%v\n", sum)
			t1 = -1
			t2 = -1
		}
		if (47 < k) && (k < 48+10) {
			fmt.Printf("%v\n", k-48)
			if t1 == -1 {
				t1 = int(k - 48)
				continue
			}

			t2 = int(k - 48)

		}
	}
	if t2 == -1 {
		t2 = t1
	}
	fmt.Printf("%v%v\n", t1, t2)
	fmt.Printf("%v\n", sum+t1*10+t2)
}
