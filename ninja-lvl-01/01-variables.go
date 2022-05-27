package main

/*
*	This ninja exercise contents:
*
*	1: Using short declaration, assign the following values to variables with names (x,y,z)
*		42, "James Bond", true
*	2: Print the stored values using
*		- Println function single declaration
*		- Println function multiple declaration
 */

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
