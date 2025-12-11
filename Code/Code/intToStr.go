package main

import (
	"fmt"
	"strconv"
)

func strCombine() {

	int1 := 96
	str1 := strconv.Itoa(int1)
	fmt.Printf("%q\n", str1)
}

func main() {
	strCombine()
}
