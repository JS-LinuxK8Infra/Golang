package main

import "fmt"

func outerInnerBlocks() {
	city := "Seattle,"
	{
		state := "Washington,"
		county := "King."
		fmt.Println(city, state, county)
	}
}

func main() {
	outerInnerBlocks()
}
