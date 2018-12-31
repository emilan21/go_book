// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToc(freezingF)) // "32F = 0C"
	fmt.Printf("%gF = %gC\n", boilingF, fToc(boilingF))   // "212F = 100C"
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
