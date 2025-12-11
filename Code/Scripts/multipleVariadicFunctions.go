package main

import "fmt"

func multNumbers(numbers ...int) int {
	mult := 20
	for _, result := range numbers {
		mult *= result
	}
	return mult
}

func printStates(name string, states ...string) {
	fmt.Println("Hi, my name is", name, "and I have lived in the following states: ")
	for _, allStates := range states {
		fmt.Printf("%s ", allStates)
	}
}

func main() {
	fmt.Println(multNumbers())
	fmt.Println(multNumbers(43))
	fmt.Println(multNumbers(39))
	fmt.Println(multNumbers(20))
	fmt.Println(multNumbers(18))
	fmt.Println(multNumbers(7))
	fmt.Println(multNumbers(2))
	printStates("Justin", "Michigan, ", "California, ", "Tennessee, ", "Montana, ", "Alabama, ", "Washington.\n")
}
