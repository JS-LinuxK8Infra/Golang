// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, assigns an int and a float to two different variables, then uses Printf to convert the float to an int and output to the terminal
func main() {

	var f64 float64 = 96.96
	var i int = int(f64)
	fmt.Printf("%v\n", i)

}
