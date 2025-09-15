// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// This the main function, defines the string to be output to the terminal, and prints it

func main() {
	const decade = 10
	const leapyear = 4
	const total_leap_year = decade / leapyear

	fmt.Println("There are", total_leap_year, "leap years in a decade.")
}
