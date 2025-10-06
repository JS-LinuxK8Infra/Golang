package main

import (
	"fmt"
	"reflect"
)

func dataType() {

	greeting := "Good Evening"
	isSummer := false
	year := 2025
	cost := 99.99
	fmt.Printf("greeting=%v is of datatype '%v' \n", greeting, reflect.TypeOf(greeting))
	fmt.Printf("isSummer=%v is of datatype %v \n", isSummer, reflect.TypeOf(isSummer))
	fmt.Printf("year=%v is of datatype %v \n", year, reflect.TypeOf(year))
	fmt.Printf("cost=%v is of datatype %v \n", cost, reflect.TypeOf(cost))
}

func main() {
	dataType()
}
