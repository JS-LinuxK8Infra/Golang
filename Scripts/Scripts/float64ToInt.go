package main

import "fmt"

// Call the main function, assigns a string and a float to two different variables, then uses Printf to combine each in a string to the terminal
func floatToInt() {
	f64 := 96.96
	int1 := int(f64)
	fmt.Printf("%v\n", int1)
}

func main() {
	floatToInt()
}
