package main

import "fmt"

func intFloat64Combine() {
	int1 := 96
	f64 := float64(int1)
	fmt.Printf("%.2f\n", f64)
}

func main() {
	intFloat64Combine()
}
