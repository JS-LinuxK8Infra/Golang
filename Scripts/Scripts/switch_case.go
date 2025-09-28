// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// Call the main function
func main() {
	//Defines an int var, uses a switch-case operator, executes the case statement and prints the result to the terminal
	an := 14
	switch an {
	case 43:
		fmt.Println("an is 43")
	case 14, 39:
		fmt.Println("an is either 14 or 39")
	default:
		fmt.Println("an is neither 14, 43 or 39.")
	}
	//Defines an int var then executes against case blocks finishing with printing the output to the terminal
	barl := 47
	switch barl {
	case 20:
		fmt.Println("barl is 20")
	case 18, 7:
		fmt.Println("barl is either 18 or 7")
	default:
		fmt.Println("barl is neither 47, 20, 17, or 7")
	}
	//Defines an int var then using fallthrough runs through each subsequent case block stopping at the first block not having fallthrough and prints that to the terminal
	jt := 82
	switch jt {
	case 43:
		fmt.Println("43")
	case 39:
		fmt.Println("39")
		fallthrough
	case 14:
		fmt.Println("14")
		fallthrough
	default:
		fmt.Println("default")
	}
	//
	j, t := 43, 39
	switch {
	case j+t == 82:
		fmt.Println("equal to 82")
	case j+t >= 14:
		fmt.Println("greater than or equal to 82")
	case j+t <= 14:
		fmt.Println("less than or equal to 82")
	default:
		fmt.Println("greater than 82")
	}

}
