package main

import "fmt"

func sevenFactorial(int1 int) int {
	if int1 == 1 {
		return 1
	}
	return int1 * sevenFactorial(int1-1)
}

func main() {
	int1 := 7
	result := sevenFactorial(int1)
	fmt.Println("The factorial of", int1, "is: ", result)
}
