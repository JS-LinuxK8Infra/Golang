// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, define two string variables and one integer variable using shorthand and outputting each to the terminal on a seperate line
func main() {
	city := "Seattle"
	{
		state := "Washington"
		county := "King"
		fmt.Println(city)
		fmt.Println(state)
		fmt.Println(county)
	}
}
