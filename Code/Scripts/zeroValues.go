package main

import "fmt"

func zeroValues() {

	var b bool
	fmt.Printf("%t\n", b)

	var i int
	fmt.Printf("%d\n", i)

	var s string
	fmt.Printf("%s\n", s)

	var f float64
	fmt.Printf("%f\n", f)
}

func main() {
	zeroValues()
}
