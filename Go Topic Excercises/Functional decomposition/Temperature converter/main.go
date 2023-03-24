package main

import "fmt"

// nolint: gomnd // <-- DO NOT delete this comment!
func fahrenheitToCelsius(temperature float64) {
	celsius := (temperature - 32) * 5 / 9
	fmt.Printf("%.2f C", celsius)
}

// nolint: gomnd // <-- DO NOT delete this comment!
func celsiusToFahrenheit(temperature float64) {
	fahrenheit := (temperature * 9 / 5) + 32
	fmt.Printf("%.2f F", fahrenheit)
}

// Create a function convertTemperature() that takes `temperature` and `unit` as arguments
// And then calls the correct function to convert the temperature based on the `unit`
func convertTemperature(temperature float64, unit string) {
	if unit == "F" {
		fahrenheitToCelsius(temperature)
	} else if unit == "C" {
		celsiusToFahrenheit(temperature)
	}
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var temperature float64
	var unit string
	fmt.Scanln(&temperature, &unit)

	convertTemperature(temperature, unit)
}
