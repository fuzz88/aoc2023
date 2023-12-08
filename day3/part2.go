package main

import (
	"fmt"
	"math"
	"os"
	"sync"
)

var wg sync.WaitGroup

var gearsMutex sync.Mutex

type Gear struct {
	i   int
	num int
}

var gears []*Gear

func main() {
	raw_data, err := os.ReadFile("test2.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(raw_data)

	line_start := 0
	line_count := 0

	for i := 0; i < len(raw_data); i++ {
		if raw_data[i] == 10 {
			go checkLineParts(&raw_data, line_start, i)
			wg.Add(1)
			line_start = i + 1
			line_count++
		}
		if i == len(raw_data)-1 {
			go checkLineParts(&raw_data, line_start, i+1)
			wg.Add(1)
			line_count++
		}
	}
	sum := 0
	// not exactly 2, but still works
	for i := 0; i < len(gears); i++ {
		for j := 0; j < len(gears); j++ {
			if gears[i] != nil && gears[j] != nil && gears[i].i == gears[j].i && i != j {
				num := gears[i].num * gears[j].num
				sum = sum + num
				gears[j] = nil
			}
		}

	}
	fmt.Println(sum)
}

func checkLineParts(raw_data *[]byte, line_start int, line_end int) {
	defer wg.Done()
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
			if num > 0 {
				if gear := isConnectedNum(raw_data, j+1, i-1, line_start, line_end, num); gear != nil {
					// fmt.Println(gear)
					// fmt.Println(num)
					gearsMutex.Lock()
					gears = append(gears, gear)
					gearsMutex.Unlock()
				}

			}

		}
	}
	// fmt.Println(gears)
}

func isConnectedNum(raw_data *[]byte, start_idx int, end_idx int, line_start int, line_end int, num int) *Gear {
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
			if (*raw_data)[i] == 42 {
				return &Gear{i, num}
			}
		}
	}
	for i := line_start + start_idx; i <= line_start+end_idx; i++ {
		// fmt.Printf("%v %v\n", (*raw_data)[i], i)
		if (*raw_data)[i] == 42 {
			return &Gear{i, num}
		}
	}
	next_start := line_start + line_len + 1
	if next_start < len((*raw_data)) {
		for i := next_start + start_idx; i <= next_start+end_idx; i++ {
			// fmt.Printf("%v %v\n", (*raw_data)[i], i)
			if (*raw_data)[i] == 42 {
				return &Gear{i, num}
			}
		}

	}
	return nil
}
