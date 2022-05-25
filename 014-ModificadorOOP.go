package main

import (
	"fmt"
	pack "package014"
)

func main() {
	var car pack.PublicCar

	car.Brand = "Nissan"
	car.Reference = "Versa"
	fmt.Println(car)

	//pack.PublicFunction()
}
