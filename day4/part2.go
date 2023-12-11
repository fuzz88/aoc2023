package main

import (
	"fmt"
	"math"
	"os"
)

type Ticket struct {
	count   int
	matches int
}

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

	results := make([]Ticket, 0)

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
			res := 0.0
			for _, t1 := range ticket_nums {
				for _, t2 := range winner_nums {
					if t1 == t2 {
						res++
					}
				}
			}
			ticket := Ticket{count: 1, matches: int(res)}
			results = append(results, ticket)
			for raw_data[i] != 10 && i != 0 {
				i--
			}
			i++
			ticket_nums = make([]int, 0)
			winner_nums = make([]int, 0)
		}

	}
	cards_all := 0
	L = 0

	for i := len(results) - 1; i >= 0; i-- {
		cards_all = cards_all + results[i].count
		copies_count := results[i].matches
		if i-copies_count >= 0 {
			L = i - copies_count
		} else {
			L = 0
		}
		for j := i - 1; j >= L; j-- {
			results[j].count = results[j].count + results[i].count
		}
	}
	fmt.Println(cards_all)
}
