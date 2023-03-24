package main

import "fmt"

func main() {
	// DO NOT delete this code block. It takes as an input the `n` value to be used in the for loop
	var n int
	fmt.Scanln(&n)

	// DO NOT delete the variable declarations that will be used within the for loop!
	var currentFib, previousFib = 1, 0
	var lastDigit int

	// Write the additional required code to finish the for loop declaration below:
	for i := 0; i < n; i++ {
		lastDigit = (currentFib + previousFib) % 10
		currentFib = previousFib
		previousFib = lastDigit
	}

	fmt.Println(lastDigit) // This line outputs the calculated `lastDigit`
}
