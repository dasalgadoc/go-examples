package main

import (
	"fmt"
)

type computer struct {
	ram   int
	disk  int
	brand string
}

func (pc computer) ping() {
	fmt.Println(pc.brand)
}

func (pc *computer) duplicateRAM() {
	pc.ram *= 2
}

func (pc computer) getBrand() string {
	return pc.brand
}

func (pc *computer) setRAM(newRam int) {
	pc.ram = newRam
}

func main() {
	normalVariable := 100
	pointer := &normalVariable

	fmt.Println("Valor variable: ", normalVariable)
	fmt.Println("Valor puntero: ", pointer)
	fmt.Println("Valor al que apunta el punter: ", *pointer)

	*pointer = 1000
	fmt.Println("Valor variable modificada a través del puntero: ", normalVariable)

	// Definición Objeto

	myPC := computer{ram: 16, disk: 512, brand: "Asus"}
	fmt.Println(myPC)

	// Uso de Función
	myPC.ping()

	// Función con apuntador
	myPC.duplicateRAM()
	fmt.Println(myPC)

	// Función con parámetro
	myPC.setRAM(8)
	fmt.Println(myPC)

	// Función con retorno
	fmt.Println(myPC.getBrand())

}
