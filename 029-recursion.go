package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hola " + strconv.Itoa(i))
	}

	printHello(0, 10)
}

func printHello(number int, iteration int) {
	fmt.Println("Hola F " + strconv.Itoa(number))
	if number < iteration {
		printHello(number+1, iteration)
	}
}
