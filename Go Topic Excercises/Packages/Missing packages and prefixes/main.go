package main

import (
	"fmt"
	. "fmt"
	"math"
	. "math"
)

const ToRadians = math.Pi / 180

func main() {
	// just add the missing prefixes to the functions below.
	var angle float64
	fmt.Scanf("%f", &angle)

	angle *= ToRadians // do not modify this line, it converts the angle to radians

	Println(Sin(angle) - math.Cos(angle))
}
