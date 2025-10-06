package main

import "fmt"

// This calls the main function, and then prints "The weather is nice today" to the terminal
func author() {
	time := "It is Saturday, September 6th when I wrote this."
	fmt.Println(time)
}

func main() {
	author()
}
