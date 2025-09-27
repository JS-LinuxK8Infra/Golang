// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// Call the main function
func main() {
	//Define an int variable, a second int variable without assigning a value to it, then uses the assign operator
	var j int = 14
	var t int
	t = j
	//Prints the unassigned int variable, and the output of the assigned operator
	fmt.Println(t)
	//Define two int variables, then uses the add and assign operator to assign the sum of both to the first variable
	var b, a int = 20, 18
	b += a
	//Prints the first int variable with the addition of the second int variable to the terminal
	fmt.Println(b)
	//Define two int variables, then uses the subtract and assign operator to subtract the first variable from the second
	var m, r int = 20, 7
	m -= r
	//Prints the result of the previous operation to the terminal
	fmt.Println(m)
	//Define two int variables, then uses the multiply and assign operator to multiply the second variable with the first
	var l, n int = 2, 14
	l *= n
	//Prints the output of the previous operation to the terminal
	fmt.Println(l)
	//Assigns an additional int variable then uses the divide and assign operator to divide the first variable by the a previously defined variable
	var x int = 149
	x /= b
	//Prints the result of the previous operation to the terminal
	fmt.Println(x)
	//Uses the divide and assign modulus operator to divide the x variable by the b variable
	x %= b
	//Prints the result of the previous operation to the terminal
	fmt.Println(x)

}
