package main

import (
	"fmt"
)

func main() {
	var impar int = 33
	var par int = 20

	if impar%2 != 0 {
		fmt.Println("Número Impar")
	}

	if impar%2 == 0 {
		fmt.Println("Número Par")
	} else {
		fmt.Println("Número Impar")
	}

	if impar%2 == 0 {
		fmt.Println("Número Par")
	} else if par%2 == 0 {
		fmt.Println("Número Par")
	} else {
		fmt.Println("Comparación inoficiosa")
	}

	fmt.Println()

	var multiplica int = impar * par
	var modulo bool = multiplica%2 == 0

	switch modulo {
	case true:
		fmt.Println("La multiplicación es par")
	default:
		fmt.Println("La multiplicación es impar")
	}

	switch mod := multiplica % 2; mod {
	case 0:
		fmt.Println("La multiplicación es par")
	default:
		fmt.Println("La multiplicación es impar")
	}

	switch {
	case modulo && par > impar:
		fmt.Println("La multiplicación es par, par mayor")
	case modulo && par < impar:
		fmt.Println("La multiplicación es par, impar mayor")
	case modulo && impar == par:
		fmt.Println("La multiplicación es par, números iguales")
	default:
		fmt.Println("La multiplicación es impar")
	}

}
