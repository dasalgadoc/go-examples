package main

import (
	"fmt"
)

func main() {
	var manolo string
	manolo = "manolo"

	fmt.Printf("%T \t %s\n", manolo, manolo)

	type alex string

	var diego alex
	diego = "salgado"

	fmt.Printf("%T \t %s\n", diego, diego)
}
