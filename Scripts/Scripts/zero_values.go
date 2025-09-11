// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, do not assign a value to any of the data type variables defined, print to terminal which returns the default zero value for each
func main() {

	var b bool
	fmt.Printf("%t\n", b)

	var i int
	fmt.Printf("%d\n", i)

	var s string
	fmt.Printf("%s\n", s)

	var f float64
	fmt.Printf("%f\n", f)
}
