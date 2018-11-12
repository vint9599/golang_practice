package main

import (
	"fmt"
)

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initilization
var y = 43

// DECLARE there is a VARIABLE with the IDENTITIER "z"
// and that the VARIABLE with the IDENTIFER　"z" is of TYPE int
// ASSIGN the ZERO VALUE of TYPE int to "z"
// and nil for pointers, functions, interfaces, slices, channels, and maps.

var z int

func main() {
	//DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	var y = 43
	fmt.Println(y)

	foo()

	fmt.Println(z)

}

func foo() {
	fmt.Println("again:", y)
}
