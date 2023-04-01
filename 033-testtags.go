package main

import (
	"fmt"
)

// función auxiliar para generar todas las combinaciones O(n^m)
func concatenar(elementos [][]string, n int, conjunto []string, resultado *[][]string) {
	// base case
	if n == len(elementos) {
		*resultado = append(*resultado, conjunto)
		return
	}
	// caso recursivo
	for _, e := range elementos[n] {
		c := make([]string, len(conjunto))
		copy(c, conjunto)
		concatenar(elementos, n+1, append(c, e), resultado)
	}
}

func main() {
	elementos := [][]string{{"DAILY", "NIGHTLY"}, {"CPF"}, {"CUSTOMIZABLE"}, {"TRANSFER"}}

	var resultado [][]string
	concatenar(elementos, 0, []string{}, &resultado)
	fmt.Println(resultado)
}
