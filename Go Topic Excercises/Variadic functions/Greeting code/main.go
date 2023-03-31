package main

import "fmt"

func main() {
	var prefix, name1, name2 string
	fmt.Scan(&prefix, &name1, &name2)

	for _, line := range Greeting(prefix, name1, name2) {
		fmt.Println(line)
	}
}

func Greeting(prefix string, names ...string) []string {
	var greetings []string

	for _, name := range names {
		greetings = append(greetings, fmt.Sprintf("%s %s", prefix, name))
	}

	return greetings
}
