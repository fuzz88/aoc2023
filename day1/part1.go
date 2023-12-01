package main

import "fmt"

func main() {
	s := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

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
