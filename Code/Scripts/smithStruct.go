package main

import "fmt"

type Smith struct {
	Name  string
	Age   int
	Role  string
	Color string
}

type JT struct {
	B int
}

type JTB struct {
	B int
}

type JTBA struct {
	B int
}

func main() {
	r := JTBA{B: 14}
	r1 := JTBA{B: 11}
	r2 := JTBA{B: 14}
	if r != r1 {
		fmt.Println("r and r1 have different integers")
	}
	if r == r2 {
		fmt.Println("r has the same integer as r2")
	}
	var t Smith
	t.Name = "Tom"
	t.Age = 43
	t.Role = "Dad"
	t.Color = "red"
	j := Smith{"Jane", 39, "Mom", "yellow"}
	fmt.Printf("%#v \n%#v \n", t, j)
	kids := []Smith{
		{Name: "Chris", Age: 20, Role: "Son", Color: "blue"},
		{Name: "Bob", Age: 18, Role: "Son", Color: "green"},
		{Name: "Suzie", Age: 7, Role: "daughter", Color: "violet"},
	}
	for _, kids := range kids {
		fmt.Printf("%#v \n", kids)
	}

}
