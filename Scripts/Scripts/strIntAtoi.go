package main

import (
	"fmt"
	"strconv"
)

func strConvAtoi() {
	numStr := "96"
	i, err := strconv.Atoi(numStr)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)
}

func main() {
	strConvAtoi()
}
