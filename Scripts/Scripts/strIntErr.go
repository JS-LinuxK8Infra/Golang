package main

import (
	"fmt"
	"strconv"
)

// Intentionally causes an error using both a string and the Atoi conversion function
func strAtoiErr() {
	nameDate := "96jtbarl"
	i, err := strconv.Atoi(nameDate)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)
}

func main() {
	strAtoiErr()
}
