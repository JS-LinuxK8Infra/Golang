// This loop is intentionally designed to never terminate on it's own. It's intention is to showcase manual intervention of when a script or code is required and will be expanded upon in the future to demonstrate an OOM error and then runtime panic.
package main

import "fmt"

func infiniteLoop() {
	int1 := 14
	for {
		int1++
	}
	fmt.Println(int1)
}

func main() {
	infiniteLoop()
}
