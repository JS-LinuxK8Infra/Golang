// Define the package to use
package main

// Import the fmt and strconv packages
import (
	"fmt"
	"strconv"
)

// Call the main function, assigns a string and an int to variables, then using the Atoi function, prints the converted data type to the terminal, along with the var output, and any errors, this script will intentionally produce an error as the string is a combination of both letters and numerators so cannot be cleanly converted
func main() {

	var s string = "96jtbarl"
	i, err := strconv.Atoi(s)

	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)

}
