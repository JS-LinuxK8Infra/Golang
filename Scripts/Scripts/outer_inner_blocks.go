// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, define three string variables, one in the outer block and two in the inner block using shorthand and outputting all to the terminal on seperate lines
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
