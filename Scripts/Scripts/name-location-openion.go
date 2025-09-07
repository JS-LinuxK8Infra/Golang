package desktop
// This defines the package to use
package main

// This imports the fmt package
import "fmt"

// This calls the main function, assigns a string to two different variables, then joins a string to both variables and prints to the terminal
func main() {

	var location string = "I live in Seattle"
	var openion string = "and I love it here"

	fmt.Print("My name is Justin, ", location, ", ", openion)
}
