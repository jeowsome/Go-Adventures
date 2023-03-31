package main

import "fmt"

func main() {
	//elements := []string{"Helium", "Boron", "Neon", "Calcium", "Iodine"}
	//
	//elem1 := elements[0][:2] // abbreviation of Helium
	//elem2 := elements[1][:1] // abbreviation of Boron
	//elem3 := elements[2][:2] // abbreviation of Neon
	//elem4 := elements[3][:2] // abbreviation of Calcium
	//elem5 := elements[4][:1] // abbreviation of Iodine
	//
	//// DO NOT delete the line below; it prints the substrings!
	//fmt.Println(elem1, elem2, elem3, elem4, elem5)

	words := []string{"saltwater", "backstage", "bedrock", "sandcastle", "snowman"}

	w1 := words[0][4:len(words[0])]
	w2 := words[1][4:len(words[1])]
	w3 := words[2][3:len(words[2])]
	w4 := words[3][4:len(words[3])]
	w5 := words[4][4:len(words[4])]

	fmt.Println(w1, w2, w3, w4, w5) // DO NOT delete this line, it prints the substrings!
}
