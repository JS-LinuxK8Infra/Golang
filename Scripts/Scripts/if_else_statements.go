// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// Call the main function
func main() {
	//Defines a string var and prints the results to the terminal based off succesful matching of an if statement
	t := "SSW"
	if t == "SSW" {
		fmt.Println(t)
	}
	//Defines a string var, executes against an if-else statement and prints a string to the terminal based on if the statement is true or not
	city := "Seattle"
	if city == "Boston" {
		fmt.Println("You live in Seattle")
	} else {
		fmt.Println("You do not live in Seattle.")
	}
	//Defines a string var using shorthand, runs through if, else if and else statements, then prints a string to the terminal indicating the true statement for the three
	state := "Washington"
	if state == "Alabama" {
		fmt.Println("You live in Alabama")
	} else if state == "Montana" {
		fmt.Println("You do not live in Alabama")
	} else {
		fmt.Println("You do not live in any of these states.")
	}
}
