package main

import "fmt"

/*
*	This ninja exercise contents:
*
*	1 - Create int data type based.
*	2 - Create a variable named x with the type declared using VAR
*	3 - In the main function:
*		- Print the value of X
*		- Print the type of X
*		- Assing 42 to X
*		- Print the value of X
 */

type variable int

func main() {
	var x variable

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
