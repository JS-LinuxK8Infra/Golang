// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, assigns a string and a float to two different variables, use Printf to out each in a string to the terminal
func main() {

	var name string = "Justin,"
	var months float64 = 3.5

	fmt.Printf("%v there are %.1f months left in the year!", name, months)

}
