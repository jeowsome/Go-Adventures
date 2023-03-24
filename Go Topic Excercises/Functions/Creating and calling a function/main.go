package main

import "fmt"

func getUser(s string) string {
	// write the function code block here
	return fmt.Sprintf("%s is learning how to call functions!", s)
}

func main() {
	// Do not change this two lines of code
	var userName string
	fmt.Scanf("%s", &userName)

	// call the function directly, or within a fmt.Println statement here.
	fmt.Println(getUser(userName))
}
