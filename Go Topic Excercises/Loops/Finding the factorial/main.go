package main

import "fmt"

func main() {
	var m, sum int
	fmt.Scan(&m)

	if m == 0 {
		sum = 1
	} else {
		sum = m
		for ; m != 1; m-- {
			sum *= m - 1
		}
	}
	fmt.Println(sum)
}
