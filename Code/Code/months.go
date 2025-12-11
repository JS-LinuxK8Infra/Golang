package main

import "fmt"

func strFlt64Convert() {
	name := "Justin,"
	months := 3.5
	fmt.Printf("%v there are %.1f months left in the year!\n", name, months)
}

func main() {
	strFlt64Convert()
}
