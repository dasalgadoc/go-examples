package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) bool {
	text = strings.ToUpper(text)
	text = strings.ReplaceAll(text, " ", "")

	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		return true
	}

	return false
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	fmt.Println(slice)

	// Range para recorrer Slices
	fmt.Println("Range para recorrer")

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i, _ := range slice {
		fmt.Println(i)
	}

	// Palíndromo
	fmt.Println("Palíndromos")

	palindromo := "anitalavalatina"
	fmt.Printf("La palabra %s es palindromo? : %v\n", palindromo, isPalindromo(palindromo))

	caseSensitivePalindromo := "aniTalAvAlATINA"
	fmt.Printf("La palabra %s es palindromo? : %v\n", caseSensitivePalindromo, isPalindromo(caseSensitivePalindromo))

	spacePalindromo := "Anita Lava La Tina"
	fmt.Printf("La palabra %s es palindromo? : %v\n", spacePalindromo, isPalindromo(spacePalindromo))
}
