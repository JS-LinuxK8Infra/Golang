package main

import "fmt"

func main() {
	int1 := 14
	fmt.Printf("%T %v \n", int1, int1)
	pint1 := &int1
	*pint1 = 23
	fmt.Printf("%T %v \n", int1, int1)
}
