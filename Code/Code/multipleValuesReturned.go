package main

import "fmt"

func ops(int1 int, int2 int) (int, int) {
	mult := int1 * int2
	div := int1 / int2
	return mult, div
}

func main() {
	mult, div := ops(43, 39)
	fmt.Println(mult, div)
}
