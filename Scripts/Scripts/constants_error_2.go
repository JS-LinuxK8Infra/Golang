// This defines the package being called
package main

// This imports the format package
import "fmt"

// Defines two constants, a third constant with a division equation, then prints the result, which causes an intentional error as the int variable is not defined

func main() {
	const decade
	decade = 10
	fmt.Println("%v: %T \n", decade, decade)
}
