package main

import (
	"fmt"
	"strconv"
)

// Call the main function, assigns a string and a float to two different variables, then uses Printf to combine each in a string to the terminal
func strCombine() {

	int1 := 96
	str1 := strconv.Itoa(int1)
	fmt.Printf("%q\n", str1)
}

func main() {
	strCombine()
}
