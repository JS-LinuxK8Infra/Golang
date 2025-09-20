// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// Call the main function
func main() {
	//Define two string variables
	var s, w string = "Seattle, ", "Washington"
	//Define two float64 variables
	var j, t float64 = 18.18, 1.1
	//Defines two integers
	var r, l int = 7, 2
	//Defines an int and increments it by one
	var a int = 18
	a++
	//Defines an int and decrements it by one
	var b int = 20
	b--
	//Adds both string variables together and combines them
	fmt.Println(s + w)
	//Prints the second decimal place, subtacts the second float64 var from the first float64 var
	fmt.Println(j - t)
	//Multiplies the left and right operand
	fmt.Println(r * l)
	//Divides the left operand by the right operand
	fmt.Println(r / l)
	//Returns the remainder after the left operand is divided by the right operand
	fmt.Println(r % l)
	//Returns the var and increments it by 1
	fmt.Println(a)
	//Returns the var and decrements it by 1
	fmt.Println(b)

}
