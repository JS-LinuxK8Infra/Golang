// This defines the package being called
package main

// This imports the format package
import (
	"fmt"
)

// Call the main function
func main() {
	//Defines an array and assigns values to all 3 strings
	var vegtibles [3]string = [3]string{"sweet potatoes,", "broccolini,", "asparagus"}
	fmt.Println(vegtibles)
	//Defines an array and assigns values to all 4 int
	years := [4]int{10, 11, 23, 25}
	fmt.Println(years)
	//Defines an array and assigns values to all 4 strings
	names := [...]string{"Frodo,", "Sam,", "Merry,", "Pippin"}
	fmt.Println(names)
	//Defines an array, assigns values to all 3 strings, prints the name of the array, calls the element from index 2 (plums), also (array indexing) replaces the element at index 0 with another string and prints the updated array
	var fruits [3]string = [3]string{"mangos,", "cantelopes,", "plums"}
	fruits[0] = "blueberries,"
	fmt.Println(len(fruits))
	fmt.Println(fruits[2])
	fmt.Println(fruits)
	//Assign ints to an 8 element array, then runs a for loop through the ints provided until reaching the end
	var jt [7]int = [7]int{10, 11, 14, 18, 20, 22, 25}

	for t := 0; t < len(jt); t++ {
		fmt.Println(jt[t])
	}
	//Second array, forloop through the array and provides the element with the corosponding index location
	for index, element := range jt {
		fmt.Println(element, "/", index)
	}
	//Third array, muiltidimensional - 2D, specifies the number of indexes as well as the number of elements per index, then calls the desired element from the specified index
	s := [2][4]string{{"j", "t"}, {"b", "a", "r", "l"}}
	fmt.Println(s[1][2])

}
