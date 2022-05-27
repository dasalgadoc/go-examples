package main

import (
	"fmt"
)

func main() {

	// Constantes

	const pi float64 = 3.14
	const au = 1.14

	fmt.Println("Valor de PI: ", pi)
	fmt.Println("Valor de AU: ", au)

	fmt.Printf("\n")

	// Variables

	var area int
	var altura int = 14
	base := 12

	fmt.Println("Base: ", base, " Altura: ", altura, " Área: ", area)

	fmt.Printf("\n")

	// Zero Values
	var entero int
	var flotante float64
	var cadena string
	var booleano bool

	fmt.Println("Entero por defecto: ", entero)
	fmt.Println("Flotante por defecto: ", flotante)
	fmt.Println("Cadena por defecto: ", cadena)
	fmt.Println("Booleano por defecto: ", booleano)

	type money int

	var price money
	price = 4000
	fmt.Printf("%T \t %v \n", price, price)
}
