package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	raw_data, err := os.ReadFile("test2.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(raw_data)

	L := len(raw_data)
	// fmt.Println(L)

	num := 0
	pow := 0.0

	ticket_nums := make([]int, 0)
	winner_nums := make([]int, 0)

	sw := true

	result := 0

	for i := L - 1; i > -1; i-- {
		// fmt.Printf("%v %v %s\n", raw_data[i], i, string(raw_data[i]))
		for raw_data[i] >= 48 && raw_data[i] <= 57 {
			num = num + int(math.Pow(10, pow))*int(raw_data[i]-48)
			pow++
			i--
		}

		if raw_data[i] == 32 {
			if sw {
				ticket_nums = append(ticket_nums, num)
			} else {
				winner_nums = append(winner_nums, num)
			}
			pow = 0.0
			num = 0
			for raw_data[i] == 32 {
				i--
			}
			i++
		}
		if raw_data[i] == 124 {
			sw = !sw
			i--
		}

		if raw_data[i] == 58 {
			// fmt.Println(ticket_nums)
			// fmt.Println(winner_nums)
			res := -1.0
			for _, t1 := range ticket_nums {
				for _, t2 := range winner_nums {
					if t1 == t2 {
						res++
					}
				}
			}
			if res != -1 {
				result = result + int(math.Pow(2, res))
			}
			for raw_data[i] != 10 && i != 0 {
				i--
			}
			i++
			ticket_nums = make([]int, 0)
			winner_nums = make([]int, 0)
		}

	}
	fmt.Println(result)
}
