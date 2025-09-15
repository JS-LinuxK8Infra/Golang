package desktop
// Define the package to use
package main

// Import the fmt and strconv packages
import (
	"fmt"
	"strconv"
)

// Call the main function, assigns a string and a float to two different variables, then uses Printf to combine each in a string to the terminal
func main() {

	var s string = "96"
	i, err := strconv.Atoi(s)

	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)

}
