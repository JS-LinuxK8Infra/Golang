// Define the package to use
package main

// Import the fmt package
import (
	"fmt"
	"reflect"
)

// Call the main function, do not assign a value to any of the data type variables defined, print to terminal which returns the default zero value for each
func main() {

	var greeting string = "Good Evening"
	var isSummer bool = false
	var year int = 2025
	var cost float64 = 99.99

	fmt.Printf("variable greeting=%v is of type '%v' \n", greeting, reflect.TypeOf(greeting))
	fmt.Printf("variable isSummer=%v is of type %v \n", isSummer, reflect.TypeOf(isSummer))
	fmt.Printf("variable year=%v is of type %v \n", year, reflect.TypeOf(year))
	fmt.Printf("variable cost=%v is of type %v \n", cost, reflect.TypeOf(cost))

}
