// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, assigns two variables - one int one float64, converts the int to a float while outputting to the terminal
func main() {

	var i int = 96
	var f64 float64 = float64(i)
	fmt.Printf("%.2f\n", f64)

}
