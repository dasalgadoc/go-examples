package main

import "fmt"

/*
*	This ninja exercise contents:
*
*	1 - Use the code 04-customdatatype.go to create hotdog data type based on int
*	2 - Cast the value to int a print it
 */

type hotdog int

var y hotdog

func main() {
	y = 42
	fmt.Printf("%T\t%v\n", y, y)

	x := int(y)
	fmt.Printf("%T\t%v\n", x, x)
}