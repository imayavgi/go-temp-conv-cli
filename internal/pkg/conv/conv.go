package conv

import "fmt"

//ConvertToCelsius ...
func ConvertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

//ConvertToFahrenheit ...
func ConvertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
