package main

import "fmt"

var y string

func main() {
	fmt.Println("printing the value of y", y, "ending")

	fmt.Printf("%T", y)
	//DECLARE a variable to be of a certain type
	// and then ASSIGN a VALUE of that type to the variable
	y = "Bond, James Bond"
}
