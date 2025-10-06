package main

import (
	"fmt"
)

func strArray() {
	vegetables := [3]string{"sweet potatoes,", "eggplant,", "asparagus"}
	fmt.Println(vegetables)
}

func yearArray() {
	years := [4]int{10, 11, 23, 25}
	fmt.Println(years)
}

func nameArray() {
	names := [...]string{"Frodo,", "Sam,", "Merry,", "Pippin"}
	fmt.Println(names)
}

func indexElementSwap() {
	fruits := [3]string{"mangos,", "banana,", "pineapple"}
	fruits[0] = "blueberries,"
	fmt.Println(len(fruits))
	fmt.Println(fruits[2])
	fmt.Println(fruits)
}

func twodArray() {
	str2DArray := [2][4]string{{"j", "t"}, {"b", "a", "r", "l"}}
	fmt.Println(str2DArray[1][2])
}

func main() {
	strArray()
	yearArray()
	nameArray()
	indexElementSwap()
	twodArray()
}
