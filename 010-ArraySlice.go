package main

import (
	"fmt"
)

func main() {

	// Array
	var array [5]int
	fmt.Println("Array no inicializado")

	fmt.Println(array)

	array[1] = 20
	fmt.Println(array)
	fmt.Println("Longitud:", len(array))
	fmt.Println("Capacidad:", cap(array))

	initializedArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array inicializado")
	fmt.Println(initializedArray)
	fmt.Println("Longitud:", len(initializedArray))
	fmt.Println("Capacidad:", cap(initializedArray))

	// Slices

	var notInitialSlice []int

	fmt.Println("Slice no inicializado")
	fmt.Println(notInitialSlice)
	fmt.Println("Longitud:", len(notInitialSlice))
	fmt.Println("Capacidad:", cap(notInitialSlice))

	initialSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice inicializado")
	fmt.Println(initialSlice)
	fmt.Println("Longitud:", len(initialSlice))
	fmt.Println("Capacidad:", cap(initialSlice))

	// Añadir elementos - Slice
	fmt.Println("Añadir a Slice")
	notInitialSlice = append(notInitialSlice, 100)
	fmt.Println(notInitialSlice)

	notInitialSlice = append(notInitialSlice, 200, 300)
	fmt.Println(notInitialSlice)

	secondSlice := []int{400, 500, 600}
	notInitialSlice = append(notInitialSlice, secondSlice...)
	fmt.Println(notInitialSlice)

	// Slicing
	fmt.Println("Primeros tres: ", notInitialSlice[:3])
	fmt.Println("Últimos tres: ", notInitialSlice[3:])
	fmt.Println("Centro: ", notInitialSlice[2:4])
	fmt.Println("Longitud:", len(notInitialSlice))
	fmt.Println("Capacidad:", cap(notInitialSlice))

	// Borrado
	notInitialSlice = append(notInitialSlice[:len(notInitialSlice)-1])
	fmt.Println(notInitialSlice)
}
