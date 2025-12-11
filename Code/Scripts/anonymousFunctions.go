package main

import "fmt"

func anonFunction() {
	anon := func(int1 int, int2 int) int {
		return int1 + int2
	}
	fmt.Printf("%T \n", anon)
	fmt.Println(anon(43, 49))
}

func anonFunction2() {
	anon2 := func(int3 int, int4 int) int {
		return int3 + int4
	}(43, 39)
	fmt.Printf("%T \n", anon2)
	fmt.Println(anon2)
}

func main() {
	anonFunction()
	anonFunction2()
}
