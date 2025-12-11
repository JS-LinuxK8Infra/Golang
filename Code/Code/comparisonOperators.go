package main

import (
	"fmt"
)

func stateEqualto() {
	state := "Washington"
	state2 := "Michigan"
	fmt.Println(state == state2)
}

func stateNotequalto() {
	state := "Washington"
	state2 := "Michigan"
	fmt.Println(state != state2)
}

func lessThanGreaterThan() {
	int1, int2 := 9, 6
	fmt.Println(int1 < int2)
	fmt.Println(int1 <= int2)
	fmt.Println(int1 > int2)
	fmt.Println(int1 >= int2)
}

func main() {
	stateEqualto()
	stateNotequalto()
	lessThanGreaterThan()
}
