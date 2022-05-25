package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// String to Int
	intString := "-42"
	stringInt, err := strconv.Atoi(intString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("De %T a %T valor = %d\n", intString, stringInt, stringInt)

	// Int to String
	normalInt := -42
	intString2 := strconv.Itoa(normalInt)
	fmt.Printf("De %T a %T valor = %s\n", normalInt, intString2, intString2)

	// String to Bool
	booleanString := "true"
	stringBool, err := strconv.ParseBool(booleanString)
	fmt.Printf("De %T a %T valor = %v\n", booleanString, stringBool, stringBool)

	// String to Float
	floatString := "3.141516"
	stringFlaot, err := strconv.ParseFloat(floatString, 64)
	fmt.Printf("De %T a %T valor = %v\n", floatString, stringFlaot, stringFlaot)

	// String to int, force to fail
	fakeIntString := "Veintidos"
	fakeStringInt, err := strconv.ParseInt(fakeIntString, 10, 64)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("De %T a %T valor = %v", fakeIntString, fakeStringInt, fakeStringInt)
	}

}
