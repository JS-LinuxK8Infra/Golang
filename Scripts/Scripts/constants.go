package main

import (
	"fmt"
)

func constantDivisor() {
	decade := 10
	leapyear := 4
	totalLeapYear := decade / leapyear

	fmt.Println("There are", totalLeapYear, "leap years in a decade.")
}

func main() {
	constantDivisor()
}
