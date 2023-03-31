package main

import "fmt"

func main() {
    var number int64
    var digit1, digit2, digit3 int8
    fmt.Scan(&number, &?, ?, ?)

    fmt.Println(suitableDigits(?))
}

func suitableDigits(number int64, digits ?) []int8 {
    var included []int8
    for number > 0 {
        digit := int8(?)
        number /= 10

        for i := 0; i < len(digits); i++ {
            if digit == digits[i] {
                digits[i] = -1 // excludes the found digit
                included = append(?, ?)
                break
            }
        }
    }

    return included
}