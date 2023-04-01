package main

import (
	"fmt"
)

func main() {

	elementos := [][]string{{"DAILY", "NIGHTLY"}, {"CPF", "CNPJ"}, {"CUSTOMIZABLE"}, {"TRANSFER"}}

	var resultado [][]string
	var index []int
	for i := 0; i < len(elementos); i++ {
		index = append(index, 0)
	}

	for {
		// agregando una combinación
		var conjunto []string
		for i, e := range elementos {
			conjunto = append(conjunto, e[index[i]])
		}
		resultado = append(resultado, conjunto)

		// aumentando el indice
		for i := len(index) - 1; i >= 0; i-- {
			index[i]++
			if index[i] >= len(elementos[i]) {
				index[i] = 0
				if i == 0 {
					break
				}
			} else {
				break
			}
		}

		//break condition
		stop := true
		for i := range index {
			if index[i] != 0 {
				stop = false
				break
			}
		}
		if stop {
			break
		}
	}
	fmt.Println(resultado)

}
