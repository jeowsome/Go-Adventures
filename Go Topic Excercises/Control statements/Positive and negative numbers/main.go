package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)

	switch {
	case number == 0:
		fmt.Println("Zero!")
	case number < 0:
		fmt.Println("Negative!")
	default:
		fmt.Println("Positive!")
	}
}
