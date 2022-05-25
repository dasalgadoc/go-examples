package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Imprimir con salto de linea
	var helloMessage string = "Hello"
	var worldMessage string = "World"
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Imprimir con formato
	var name string = "Diego"
	var year int16 = 2000
	fmt.Printf("%s se encuentra en el año %d\n", name, year)
	fmt.Printf("%v se encuentra en el año %v\n", name, year)

	// Tratamiento de la salida
	var message string = fmt.Sprintf("%v se encuentra en el año %v\n", name, year)
	fmt.Println(reflect.TypeOf(message))

	fmt.Println()
	// Conocer el tipo de dato de una variable
	fmt.Printf("Message tipo de dato: %T\n", message)

}
