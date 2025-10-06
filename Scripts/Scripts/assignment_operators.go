package main

import (
	"fmt"
)

func intReplace() {
	int1 := 14
	var int2 int
	int2 = int1
	fmt.Println(int2)
}

func addAssign() {
	int1, int2 := 20, 18
	int1 += int2
	fmt.Println(int1)
}

func subAssign() {
	int1, int2 := 20, 7
	int1 -= int2
	fmt.Println(int1)
}

func multAssign() {
	int1, int2 := 2, 14
	int1 *= int2
	fmt.Println(int1)
}

func divAssign() {
	int1, int2 := 914, 149
	int1 /= int2
	fmt.Println(int1)
}

func main() {
	intReplace()
	addAssign()
	subAssign()
	multAssign()
	divAssign()

}
