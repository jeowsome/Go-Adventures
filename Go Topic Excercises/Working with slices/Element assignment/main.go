package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func solve(s []string, x string, i int) []string {
	// Write the code to make a copy of the 's' slice below:
	slCopy := make([]string, len(s))
	copy(slCopy, s)

	// Assign x to the element at the 'i' index of the 'slCopy' slice:
	slCopy[i] = x

	return slCopy // Return the copy of the slice here!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	var i int
	var x string

	mySlice := make([]string, 1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mySlice[i], x = scanner.Text(), scanner.Text()

	newSlice := solve(mySlice, x, i)
	if !reflect.DeepEqual(mySlice, newSlice) {
		log.Fatal("You didn't assign 'x' to the element at the 'i' index of the slice!")
	}
	fmt.Println(newSlice)
}
