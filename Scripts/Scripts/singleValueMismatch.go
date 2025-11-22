package main

import "fmt"

func subNumbers(int1 int, int2 int) bool {
	sub := int1 - int2
	return sub
}

func main() {
	intDelta := subNumbers(43, 39)
	fmt.Println(intDelta)
}
