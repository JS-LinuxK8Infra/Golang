package main

import (
	"fmt"
)

func boolOperators() {
	int1, int2 := 14, 96
	fmt.Println((int2 < 96) && (int2 < 58))
	fmt.Println((int1 < 7) && (int1 < 2))
	fmt.Println((int2 < 96) || (int2 < 58))
	fmt.Println((int1 < 7) || (int1 > 58))
	fmt.Println(!(int1 > int2))
	fmt.Println(!(true))
	fmt.Println(!(false))
}

func main() {
	boolOperators()
}
