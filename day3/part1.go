package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex
	v  int
}

func (c *SafeCounter) Inc(n int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v = c.v + n
	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v
}

func main() {
	raw_data, err := os.ReadFile("test2.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(raw_data)
	sum := SafeCounter{v: 0}
	line_start := 0
	line_count := 0
	for i := 0; i < len(raw_data); i++ {
		if raw_data[i] == 10 {
			checkLineParts(&raw_data, line_start, i, &sum)
			line_start = i + 1
			line_count++
		}
		if i == len(raw_data)-1 {
			checkLineParts(&raw_data, line_start, i+1, &sum)
			line_count++
		}
	}
	fmt.Println(sum.Value())

}

func checkLineParts(raw_data *[]byte, line_start int, line_end int, sum *SafeCounter) {
	line_data := (*raw_data)[line_start:line_end]
	// fmt.Println(line_data)
	for i := 0; i < len(line_data); i++ {
		if (line_data[i] < 48 || line_data[i] > 57) || i == len(line_data)-1 {
			pow := 0.0
			num := 0
			j := 0
			if i == len(line_data)-1 && line_data[i] >= 48 && line_data[i] <= 57 {
				j = i
			} else {
				j = i - 1
			}

			// fmt.Printf("%v %v", line_data[i], i)
			for ; j > -1 && line_data[j] >= 48 && line_data[j] <= 57; j-- {
				num = num + int(math.Pow(10, pow))*int((line_data[j]-48))
				pow++
			}
			if num > 0 && isConnectedNum(raw_data, j+1, i-1, line_start, line_end) {
				fmt.Println(num)
				sum.Inc(num)
			}

		}
	}
}

func isConnectedNum(raw_data *[]byte, start_idx int, end_idx int, line_start int, line_end int) bool {
	line_len := line_end - line_start
	if start_idx > 0 {
		start_idx--
	}
	if end_idx < line_len {
		end_idx++
	}
	// fmt.Println(start_idx, end_idx, line_len)
	prev_start := line_start - line_len - 1
	if prev_start >= 0 {
		for i := prev_start + start_idx; i <= prev_start+end_idx; i++ {
			// fmt.Printf("%v %v\n", (*raw_data)[i], i)
			if ((*raw_data)[i] < 48 || (*raw_data)[i] > 57) && (*raw_data)[i] != 46 {
				return true
			}
		}
	}
	for i := line_start + start_idx; i <= line_start+end_idx; i++ {
		// fmt.Printf("%v %v\n", (*raw_data)[i], i)
		if ((*raw_data)[i] < 48 || (*raw_data)[i] > 57) && (*raw_data)[i] != 46 {
			return true
		}
	}
	next_start := line_start + line_len + 1
	if next_start < len((*raw_data)) {
		for i := next_start + start_idx; i <= next_start+end_idx; i++ {
			// fmt.Printf("%v %v\n", (*raw_data)[i], i)
			if ((*raw_data)[i] < 48 || (*raw_data)[i] > 57) && (*raw_data)[i] != 46 {
				return true
			}
		}

	}
	return false
}
