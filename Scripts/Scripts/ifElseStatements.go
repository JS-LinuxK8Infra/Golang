package main

import (
	"fmt"
)

func abrieve() {
	abriev := "SSW"
	if abriev == "SSW" {
		fmt.Println(abriev)
	}
}

func city() {
	city := "Seattle"
	if city == "Boston" {
		fmt.Println("You live in Seattle")
	} else {
		fmt.Println("You do not live in Seattle.")
	}
}

func state() {
	state := "Washington"
	if state == "Alabama" {
		fmt.Println("You live in Alabama")
	} else if state == "Montana" {
		fmt.Println("You do not live in Alabama")
	} else {
		fmt.Println("You do not live in any of these states.")
	}
}

func main() {
	abrieve()
	city()
	state()
}
