package main

/*
*	This ninja exercise contents:
*
*	1 - Using the code developed in 02-variables.go assign the values:
*		- 42
*		- James Bond
*		- true
*	2 - Create a variable with formmated string and print it
 */

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%d \t %s \t %v", x, y, z)
	fmt.Println(s)
}
