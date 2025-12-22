package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 8)
	ch <- 43
	ch <- 39
	ch <- 14
	var1, ok := <-ch
	fmt.Println(var1, ok)
	close(ch)
	var1, ok = <-ch
	fmt.Println(var1, ok)
	var1, ok = <-ch
	fmt.Println(var1, ok)
	var1, ok = <-ch
	fmt.Println(var1, ok)
}
