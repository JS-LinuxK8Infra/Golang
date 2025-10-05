package main

import (
	"fmt"
)

func stringAddition() {
	city, state := "Seattle, ", "Washington"
	fmt.Println(city + state)
}

func float64Arithmetic() {
	float1, float2 := 18.18, 1.1
	fmt.Println(float1 - float2)
}

func intArithmetic() {
	int1, int2 := 7, 2
	fmt.Println(int1 * int2)
	fmt.Println(int1 / int2)
	fmt.Println(int1 % int2)
}

func intIncrementDecrement() {
	num1 := 18
	num1++
	num2 := 20
	num2--
	fmt.Println(num1)
	fmt.Println(num2)
}

func main() {
	stringAddition()
	float64Arithmetic()
	intArithmetic()
	intIncrementDecrement()
}
