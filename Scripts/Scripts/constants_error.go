// This defines the package being called
package main

// This imports the format package
import "fmt"

// Defines two constants, a third constant with a division equation, when printing to the terminal this will intentionally cause an error due to mismatched datatypes

func main() {
	const decade = 10
	decade = 96
	fmt.Println("%v: %T \n", decade, decade)
}
