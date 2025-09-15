// Define the package to use
package main

// Import the fmt and strconv packages
import (
	"fmt"
	"strconv"
)

// Call the main function, assigns a string and an integer to two different variables, then uses Printf to convert the int to a str
func main() {

	var i int = 96
	var s string = strconv.Itoa(i)
	fmt.Printf("%q\n", s)

}
