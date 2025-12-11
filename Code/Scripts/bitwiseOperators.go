package main

import (
	"fmt"
)

func bitwiseAND() {
	int1, int2 := 43, 39
	int3 := int1 & int2
	fmt.Println(int3)
}

func bitwiseOR() {
	int1, int2 := 20, 20
	int3 := int1 | int2
	fmt.Println(int3)
}

func bitwiseXOR() {
	int1, int2 := 18, 7
	int3 := int1 ^ int2
	fmt.Println(int3)
}

func bitwiseLeftShift() {
	int1 := 2
	int2 := int1 << 2
	fmt.Println(int2)
}

func bitwiseRightShift() {
	int1 := 2
	int2 := int1 >> 2
	fmt.Println(int2)
}

func main() {
	bitwiseAND()
	bitwiseOR()
	bitwiseXOR()
	bitwiseLeftShift()
	bitwiseRightShift()
}
