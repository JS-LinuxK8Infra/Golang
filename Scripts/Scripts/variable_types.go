// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, assigns a single variable to four different data types, using Printf, returns both the variable and the data type it is to the terminal
func main() {

	var greeting string = "Good Evening"
	var isSummer bool = false
	var year int = 2025
	var cost float64 = 99.99

	fmt.Printf("variable greeting = '%v' is of type %T \n", greeting, greeting)
	fmt.Printf("variable isSummer = '%v' is of type %T \n", isSummer, isSummer)
	fmt.Printf("variable year = %v is of type %T \n", year, year)
	fmt.Printf("variable cost = %v is of type %T \n", cost, cost)

}
