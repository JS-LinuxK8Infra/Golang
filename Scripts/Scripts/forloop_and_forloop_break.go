// Define the package to use
package main

// Import the fmt package
import "fmt"

// Forloop that assigns an int var then loops through until it has exceeded 125. Additionally, it will increment by 5 each time
func main() {

	s := 25
	for s <= 125 {
		fmt.Println(s * 1)
		s += 5
	}
	//Infinite loop that increments the jt var increasing by 1 each time and will not terminate until manually ended
	//jt := 14
	//for jt <= 25 {
	//	fmt.Println(s * 1)
	//	s += 1
	//}
	//Forloop with break statement forcing the script to terminate the as soon as it hits 20
	for b := 1; b <= 20; b++ {
		if b == 20 {
			break
		}
		fmt.Println(b)
	}
	//Forloop with continue statement which continues immediatly when encountering 3 and 6, terminating when it encounters 9
	for a := 1; a <= 18; a++ {
		if a == 3 || a == 6 || a == 9 || a == 12 || a == 15 || a == 18 {
			continue
		}
		fmt.Println(a)
	}

}
