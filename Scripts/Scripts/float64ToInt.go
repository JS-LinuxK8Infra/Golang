package main

import "fmt"

func floatToInt() {
	f64 := 96.96
	int1 := int(f64)
	fmt.Printf("%v\n", int1)
}

func main() {
	floatToInt()
}
