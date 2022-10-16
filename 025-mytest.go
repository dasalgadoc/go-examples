package main

import (
	"fmt"
	"strconv"
)

type car struct {
	brand     string
	reference string
	model     int
}


func main(){
	var myCar = populateCars()
	fmt.Println(myCar)

	fmt.Println("------------")

	var myDoubleCars = doublePopulateCars()
	fmt.Println(myDoubleCars)
}

func populateCars() []car{
	theCars := []car{}
	for i := 0; i <10; i++{
		newCar := car{
			brand : "A" + strconv.Itoa(i),
			reference: "XXX",
			model: 2020,
		}
		theCars = append(theCars, newCar)
	}
	return theCars
}

func doublePopulateCars() [][]car{
	theDoubleCars := [][]car{}
	for i := 0; i < 3; i++{
		theCars := populateCars()
		theDoubleCars = append(theDoubleCars, theCars)
	}
	return theDoubleCars
}