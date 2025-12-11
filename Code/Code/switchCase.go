package main

import (
	"fmt"
)

func switchCase() {
	int1 := 14
	switch int1 {
	case 43:
		fmt.Println("int1 is 43")
	case 14, 39:
		fmt.Println("int1 is either 14 or 39")
	default:
		fmt.Println("int1 is neither 14, 43 or 39.")
	}
}

func caseBlock() {
	int1 := 47
	switch int1 {
	case 20:
		fmt.Println("int1 is 20")
	case 18, 7:
		fmt.Println("int1 is either 18 or 7")
	default:
		fmt.Println("int1 is neither 47, 20, 17, or 7")
	}
}

func fallThrough() {
	int1 := 82
	switch int1 {
	case 43:
		fmt.Println("43")
	case 39:
		fmt.Println("39")
		fallthrough
	case 14:
		fmt.Println("14")
		fallthrough
	default:
		fmt.Println("default")
	}
}

func caseBlockFallThrough() {
	int1, int2 := 43, 39
	switch {
	case int1+int2 == 82:
		fmt.Println("equal to 82")
	case int1+int2 >= 14:
		fmt.Println("greater than or equal to 82")
	case int1+int2 <= 14:
		fmt.Println("less than or equal to 82")
	default:
		fmt.Println("greater than 82")
	}
}

func main() {
	switchCase()
	caseBlock()
	fallThrough()
	caseBlockFallThrough()
}
