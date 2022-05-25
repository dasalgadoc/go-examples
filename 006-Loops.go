package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	var counter int = 10

	for counter < 20 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println()

	for i := counter; i > 0; i-- {
		fmt.Println(i)
	}

	for {
		fmt.Println("Me ejecutaré por siempre por favor DETENME!")
	}

}
