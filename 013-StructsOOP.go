package main

import (
	"fmt"
)

type car struct {
	brand     string
	reference string
	model     int
}

func NewCar(brand, reference string, model int) *car {
	return &car{
		brand:     brand,
		reference: reference,
		model:     model,
	}
}

func main() {
	// Método 1
	myCar := car{brand: "Nissan", reference: "March", model: 2019}
	fmt.Println(myCar)

	// Método 2
	var daniCar car
	daniCar.brand = "Renault"
	daniCar.reference = "Sandero"
	fmt.Println(daniCar)

	// Método 3
	mantiCar := new(car)
	mantiCar.brand = "Volkswagen"
	mantiCar.reference = "Bora"
	fmt.Println(mantiCar)

	// Método 4
	mayeCar := NewCar("Suzuki", "Swift", 0)
	fmt.Println(mayeCar)

}
