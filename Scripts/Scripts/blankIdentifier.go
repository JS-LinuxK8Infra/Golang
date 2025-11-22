package main

import "fmt"

func blankIdentifierInt() (int, int, int) {
	return 43, 39
}

func main() {
	exclude, _ := blankIdentifierInt()
	fmt.Println(exclude)
}
