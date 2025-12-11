//This is an intentional error to showcase debugging of a common issue
package main

import "fmt"

func constantError() {
	decade := 
	decade2 := 10
	fmt.Println("%v: %T \n", decade, decade2)
}

func main() {
	constantError()
}
