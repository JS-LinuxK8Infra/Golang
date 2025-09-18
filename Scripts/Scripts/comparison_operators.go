// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// Call the main function
func main() {
	//Define two string variables
	var state string = "Washington"
	var state2 string = "Michigan"
	//Define two integar variables
	var i, i2 = 9, 6
	//Prints the string variables with the equal to operator
	fmt.Println(state == state2)
	//Prints the string variables with the not equal to operator
	fmt.Println(state != state2)
	//Prints the integar variables with the less than operator
	fmt.Println(i < i2)
	//Prints the integar variables with the less than or equal to operator
	fmt.Println(i <= i2)
	//Prints the integar variables with the greater than operator
	fmt.Println(i > i2)
	//Prints the integar variables with the greater than or equal to operator
	fmt.Println(i >= i2)
}
