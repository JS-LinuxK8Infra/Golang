// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// Call the main function
func main() {
	//Define an int variable
	var j, t int = 14, 96
	//Prints the int variable, executes a less than calculation against two other integers and outputs the results in a boolan format using the AND logical operator
	fmt.Println((t < 96) && (t < 58))
	//Prints the int variable again, executes a less than calculation against two other integers and outputs the results in a boolan format using the AND logical operator
	fmt.Println((t < 7) && (t < 2))
	//Prints the int variable again, executes a less than calculation against two other integers and outputs the results in a boolan format using the OR logical operator
	fmt.Println((t < 96) || (t < 58))
	//Prints the int variable again, executes a less than calculation against two other integers and outputs the results in a boolan format using the OR logical operator
	fmt.Println((t < 7) || (t > 58))
	//Prints the int variable again, executes a less than calculation against two other integers and outputs the results in a boolan format using the NOT logical operator
	fmt.Println(!(j > t))
	fmt.Println(!(true))
	fmt.Println(!(false))

}
