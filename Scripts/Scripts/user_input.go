package desktop
// Define the package to use
package main

// Import the fmt package
import "fmt"

// Call the main function, do not assign a value to any of the data type variables defined, print to terminal which returns the default zero value for each
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

	count, err := fmt.Scanf("%s %s %s %t %f", &name, &color, &city, &mood, &time)

	fmt.Println("count : ", count)
	fmt.Println("error: ", err)
	fmt.Println("name: ", name)
	fmt.Println("color: ", color)
	fmt.Println("city: ", city)
	fmt.Println("mood: ", mood)
	fmt.Println("time: ", time)

}
