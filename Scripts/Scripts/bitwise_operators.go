// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// Call the main function
func main() {
	//Define two int variables, a third int variable which using the bitwise AND operator is equal to the combination of the first two variables
	var j, t int = 43, 39
	tn := j & t
	//Define two int variables, a third int variable which using the bitwise AND operator is equal to the combination of the first two variables
	var b, m int = 20, 20
	mi := b | m
	//Define two int variables, a third int variable which using the bitwise AND operator is equal to the combination of the first two variables
	var a, r int = 18, 7
	mt := a ^ r
	//Define an int variable, and a second int variable which using the bitwise leftshift operator shifts all bits two to the left
	var l int = 2
	al := l << 2
	//Define an int variable, and a second int variable which using the bitwise rightshift operator shifts all bits two to the right
	var o int = 2
	wa := o >> 2
	//Prints the unassigned int variables, and the output of the assigned operator
	fmt.Println(tn)
	fmt.Println(mi)
	fmt.Println(mt)
	fmt.Println(al)
	fmt.Println(wa)

}
