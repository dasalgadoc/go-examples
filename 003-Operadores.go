package main

import (
	"fmt"
)

func main() {

	var a int = 15
	var b int = 10
	var c int
	var d int = 1

	// Suma
	c = a + b
	fmt.Println("a + b = ", c)

	// Resta
	c = a - b
	fmt.Println("a - b = ", c)

	// Mutiplicación
	c = a * b
	fmt.Println("a * b = ", c)

	// División
	c = a / b
	fmt.Println("a / b = ", c)

	// Módulo
	c = a % b
	fmt.Println("a mod b = ", c)

	// Incremental
	d = 1
	d++
	fmt.Println("d++ = ", d)

	// Decremental
	d--
	fmt.Println("d-- = ", d)

}
