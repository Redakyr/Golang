// Importing package main, so we can create the Main function, which is the entry point for our program.
package main

// The "fmt" lib, will allow us to print output data.
import "fmt"

// To run the program, you'll need to build it and create an executable.
func main() {
	// Creating variables stalker and town.
	var stalker string
	var town string

	// Using "fmt" library and Print function, for sending output to stdout (standard output).
	fmt.Println("What should we call you, stalker?")
	// Using "fmt" library and Scan function. Scanning for inputs on stdin and "&" operator so we can obtain the memory address for out "stalker" variable.
	fmt.Scan(&stalker)

	fmt.Println("What town are you from?")
	fmt.Scan(&town)

	fmt.Printf("Hmm... Welcome to the zone, %s from... %s. Off with you now!", stalker, town)
}
