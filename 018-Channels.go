package main

import (
	"fmt"
)

// Canal de Entrada y Salida
//func say(text string, c chan string) {

// Canal solo de salida
//func say(text string, c <- chan string) {

// Canal solo de entrada
func say(text string, c chan<- string) {
	c <- text
}

func main() {
	// Canal / Tipo de dato / Datos simultáneos (GoRutine)
	channel := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", channel)

	// Sacar el contenido del canal (El texto insertado en la GoRoutine)
	fmt.Println(<-channel)
}
