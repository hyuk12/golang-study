package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the variable 43
// declare & variable = 43

// var 를  쓸 때는 가장 중요한 것은 범위를 잘 조절해서 써야한다.
var y = 42

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {

	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 45
	fmt.Println(x)

	fmt.Println(y)
	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
