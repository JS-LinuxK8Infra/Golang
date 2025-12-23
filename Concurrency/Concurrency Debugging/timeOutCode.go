//Defines the main package
package main 
//imports the fmt and time packages
import (
	"fmt"  
	"time" 
)

//Creates the main func
func main() { 
	//Creates an unbuffered channel
	chan1 := make(chan int) 
	//Creates a go routine and takes input from the channel
	go transmitValue(chan1) 
	//Creates a select statement
	select { 
	//Creates a case statement with a message that is sent via channel created
	case channelMessage := <-chan1: 
		//Prints the message from the case statement
		fmt.Println(channelMessage) 
	//Sets a two second timeout via time.Second
	case <-time.After(2 * time.Second): 
	//Prints a string to the terminal after the timeout expires
		fmt.Println("Code timeout concluded.") 
	}
}
//Create a function that calls the channel defined in the main function and defines the data type as int
func transmitValue(chan1 chan int) { 
//Adds a second timeout that supersedes the one in the select statement; functions as a timeout for the channel transmit
	time.Sleep(4 * time.Second) 
	chan1 <- 14                 
}
