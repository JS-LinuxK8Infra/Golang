package main

import "fmt"

func dataTypeSamples() {
	greeting := "Good Evening"
	isSummer := false
	year := 2025
	cost := 99.99
	fmt.Printf("variable greeting = '%v' is of datatype %T. \n", greeting, greeting)
	fmt.Printf("variable isSummer = '%v' is of datatype %T. \n", isSummer, isSummer)
	fmt.Printf("variable year = %v is of datatype %T. \n", year, year)
	fmt.Printf("variable cost = %v is of datatype %T. \n", cost, cost)
}

func main() {
	dataTypeSamples()
}
