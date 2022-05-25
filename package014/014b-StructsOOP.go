package package014

import "fmt"

type privateCar struct {
	brand     string
	reference string
	model     int
}

// PublicCar class
type PublicCar struct {
	Brand     string
	Reference string
	Model     string
	sku       string
}

// PublicFunction
func PublicFunction() {
	fmt.Println("Función Pública")
}
