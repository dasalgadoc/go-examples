package main

/*
*	This ninja exercise contents:
*
*	1 - Use the var statement to declare variables in the package scope.
*		Let no value in the variables and named them as follow:
+			- x as int
*			- y as string
*			- z as bool
*	2 - Print values from variables in exercise 1
*/

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
}
