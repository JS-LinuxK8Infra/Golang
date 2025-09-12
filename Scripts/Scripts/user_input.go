// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, defines (3) string vars, (1) bool var, (1) float64 var. Then prints each var on a new line prompting the user to enter input.
func main() {

	var name string
	var color string
	var city string
	var mood bool
	var time float64
	fmt.Print("Please tell me your name.\n")
	fmt.Print("What is your favorite color?\n")
	fmt.Print("What city do you live in?\n")
	fmt.Print("Are you having a good day? Please enter true or false.\n")
	fmt.Print("What time is it where you are? If it is 10:34 please enter 10.34.\n")
	// Uses the Scanf function to output the expected data type format to the terminal, appends each var with & to accept user input, err handles any errors produced when the user enters the wrong data type or neglects to enter all requested input.
	count, err := fmt.Scanf("%s %s %s %t %f", &name, &color, &city, &mood, &time)
	// Prints the count, errors, and each variable to the terminal so the user may enter the requested input based on the data type.
	fmt.Println("count : ", count)
	fmt.Println("error: ", err)
	fmt.Println("name: ", name)
	fmt.Println("color: ", color)
	fmt.Println("city: ", city)
	fmt.Println("mood: ", mood)
	fmt.Println("time: ", time)

}
