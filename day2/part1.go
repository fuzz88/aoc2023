package main

import (
	"fmt"
	"math"
	"os"
)

func checkPossible(start int, end int, raw_data *[]byte, sum chan int) {
	//sends game number to sum channel if possible,
	//or 0 when impossible

	game_num := 0
	possible := true
main_loop:
	for i := start; i < end; i++ {
		if (*raw_data)[i] == 58 {
			pow := 0.0
			for j := i - 1; (*raw_data)[j] != 32; j-- {
				game_num = game_num + int(math.Pow(10, pow))*int(((*raw_data)[j]-48))
				pow++
			}
		}

		if (*raw_data)[i] == 59 || i == end-1 {
			red := 0
			green := 0
			blue := 0

			for j := i - 1; ((*raw_data)[j] != 58) && ((*raw_data)[j] != 59); j-- {
				if (*raw_data)[j] == 98 {
					pow := 0.0
					k := j - 2
					for ; (*raw_data)[k] != 32; k-- {
						blue = blue + int(math.Pow(10, pow))*int(((*raw_data)[k]-48))
						pow++
					}
				}
				if (*raw_data)[j] == 114 {
					pow := 0.0
					k := j - 2
					for ; (*raw_data)[k] != 32; k-- {
						red = red + int(math.Pow(10, pow))*int(((*raw_data)[k]-48))
						pow++
					}
				}
				if (*raw_data)[j] == 103 {
					pow := 0.0
					k := j - 2
					for ; (*raw_data)[k] != 32; k-- {
						green = green + int(math.Pow(10, pow))*int(((*raw_data)[k]-48))
						pow++
					}
				}
			}
			possible = red <= 12 && green <= 13 && blue <= 14
			if !possible {
				break main_loop
			}
		}

	}

	if possible {
		sum <- game_num
	} else {
		sum <- 0
	}

}

func main() {
	sum := make(chan int, 100)
	raw_data, err := os.ReadFile("test2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v bytes read.\n", len(raw_data))
	start := 0
	end := 0
	game_count := 0
	for i := 0; i < len(raw_data); i++ {
		if raw_data[i] == byte(10) {
			end = i
			go checkPossible(start, end, &raw_data, sum)
			game_count++
			start = end + 1
		}
		if i == len(raw_data)-1 {
			go checkPossible(start, i+1, &raw_data, sum)
			game_count++
		}
	}
	acc := 0
	for i := 0; i < game_count; i++ {
		c := <-sum
		acc = acc + c

	}
	fmt.Printf("\n%v\n\n", acc)

}
