package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	ch <- 43
	ch <- 39
	ch <- 14
	close(ch)

	for t := range ch {
		fmt.Println(t)
	}
}
