package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["Diego"] = 100
	myMap["Daniela"] = 200

	fmt.Println(myMap)

	// Recorrer Map
	for i, valor := range myMap {
		fmt.Println(i, valor)
	}

	// Encontrar valores
	value := myMap["Daniela"]
	fmt.Println(value)

	value = myMap["Mantilla"]
	fmt.Println(value)

	value, ok := myMap["Mantilla"]
	fmt.Println(ok, value)

	if !ok {
		fmt.Printf("El valor %s no existe\n", "Mantilla")
	}

	value, ok = myMap["Daniela"]
	fmt.Println(ok, value)

	if !ok {
		fmt.Printf("El valor %s no existe\n", "Daniela")
	}

}
